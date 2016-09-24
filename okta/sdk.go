package okta

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"sync"
	"time"

	"github.com/google/go-querystring/query"

	"reflect"
)

const (
	libraryVersion             = "1"
	userAgent                  = "oktasdk-go/" + libraryVersion
	productionURLFormat        = "https://%s.okta.com/api/v1/"
	previewProductionURLFormat = "https://%s.oktapreview.com/api/v1/"
	headerRateLimit            = "X-Rate-Limit-Limit"
	headerRateRemaining        = "X-Rate-Limit-Remaining"
	headerRateReset            = "X-Rate-Limit-Reset"
	headerOKTARequestID        = "X-Okta-Request-Id"
	headerAuthorization        = "Authorization"
	headerAuthorizationFormat  = "SSWS %v"
	mediaTypeJSON              = "application/json"
	defaultLimit               = 50
	filterEqualOperator        = "eq"
	filterStartsWithOperator   = "sw"
	filterGreaterThanOperator  = "gt"
	filterLessThanOperator     = "lt"
)

// A Client manages communication with the API.
type Client struct {
	clientMu sync.Mutex   // clientMu protects the client during calls that modify the CheckRedirect func.
	client   *http.Client // HTTP client used to communicate with the API.

	// Base URL for API requests.
	//  This will be built automatically based on inputs to NewClient
	//  If needed you can override this if needed (your URL is not *.okta.com or *.oktapreview.com)
	BaseURL *url.URL

	// User agent used when communicating with the GitHub API.
	UserAgent string

	apiKey                   string
	authorizationHeaderValue string
	PauseOnRateLimit         bool
	rateMu                   sync.Mutex
	Limit                    int8
	// mostRecent rateLimitCategory

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// Services used for talking to different parts of the  API.
	// Service for Working with Users
	Users *UsersService

	// Service for Working with Groups
	Groups *GroupsService
}

type service struct {
	client *Client
}

// NewClient returns a new  API client.  If a nil httpClient is
// provided, http.DefaultClient will be used.

func NewClient(httpClient *http.Client, orgName string, apiToken string, isProduction bool) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	var baseURL *url.URL
	if isProduction {
		baseURL, _ = url.Parse(fmt.Sprintf(productionURLFormat, orgName))
	} else {
		baseURL, _ = url.Parse(fmt.Sprintf(previewProductionURLFormat, orgName))

	}

	c := &Client{client: httpClient, BaseURL: baseURL, UserAgent: userAgent}
	c.PauseOnRateLimit = true // If rate limit found it will block until that time. If false then Error will be returned
	c.authorizationHeaderValue = fmt.Sprintf(headerAuthorizationFormat, apiToken)
	c.apiKey = apiToken
	c.Limit = defaultLimit
	c.common.client = c

	c.Users = (*UsersService)(&c.common)
	c.Groups = (*GroupsService)(&c.common)

	return c
}

// Rate represents the rate limit for the current client.
type Rate struct {
	// The number of requests per minute the client is currently limited to.
	Limit int

	// The number of remaining requests the client can make this minute
	Remaining int

	// The time at which the current rate limit will reset.
	ResetTime time.Time
}

type Response struct {
	*http.Response

	// These fields provide the page values for paginating through a set of
	// results.

	NextURL *url.URL
	// PrevURL       *url.URL
	SelfURL       *url.URL
	OKTARequestID string
	Rate
}

// newResponse creates a new Response for the provided http.Response.
func newResponse(r *http.Response) *Response {
	response := &Response{Response: r}

	response.OKTARequestID = r.Header.Get(headerOKTARequestID)

	response.populatePaginationURLS()
	response.Rate = parseRate(r)
	return response
}

// populatePageValues parses the HTTP Link response headers and populates the
// various pagination link values in the Response.

// OKTA LINK Header takes this form:
// 		Link: <https://yoursubdomain.okta.com/api/v1/users?after=00ubfjQEMYBLRUWIEDKK>; rel="next",
// 			<https://yoursubdomain.okta.com/api/v1/users?after=00ub4tTFYKXCCZJSGFKM>; rel="self"

func (r *Response) populatePaginationURLS() {

	for k, v := range r.Header {

		if k == "Link" {
			nextRegex := regexp.MustCompile(`<(.*?)>; rel="next"`)
			// prevRegex := regexp.MustCompile(`<(.*?)>; rel="prev"`)
			selfRegex := regexp.MustCompile(`<(.*?)>; rel="self"`)

			for _, linkValue := range v {
				nextLinkMatch := nextRegex.FindStringSubmatch(linkValue)
				if len(nextLinkMatch) != 0 {
					r.NextURL, _ = url.Parse(nextLinkMatch[1])
				}
				selfLinkMatch := selfRegex.FindStringSubmatch(linkValue)
				if len(selfLinkMatch) != 0 {
					r.SelfURL, _ = url.Parse(selfLinkMatch[1])
				}
				// prevLinkMatch := prevRegex.FindStringSubmatch(linkValue)
				// if len(prevLinkMatch) != 0 {
				// 	r.PrevURL, _ = url.Parse(prevLinkMatch[1])
				// }
			}
		}
	}

}

// parseRate parses the rate related headers.
func parseRate(r *http.Response) Rate {
	var rate Rate

	if limit := r.Header.Get(headerRateLimit); limit != "" {
		rate.Limit, _ = strconv.Atoi(limit)
	}
	if remaining := r.Header.Get(headerRateRemaining); remaining != "" {
		rate.Remaining, _ = strconv.Atoi(remaining)
	}
	if reset := r.Header.Get(headerRateReset); reset != "" {
		if v, _ := strconv.ParseInt(reset, 10, 64); v != 0 {
			rate.ResetTime = time.Unix(v, 0)
		}
	}
	return rate
}

// Do sends an API request and returns the API response.  The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred.  If v implements the io.Writer
// interface, the raw response body will be written to v, without attempting to
// first decode it.  If rate limit is exceeded and reset time is in the future,
// Do returns *RateLimitError immediately without making a network API call.
func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {

	// If we've hit rate limit, don't make further requests before Reset time.
	if err := c.checkRateLimitBeforeDo(req); err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		// Drain up to 512 bytes and close the body to let the Transport reuse the connection
		io.CopyN(ioutil.Discard, resp.Body, 512)
		resp.Body.Close()
	}()

	response := newResponse(resp)

	// c.rateMu.Lock()

	// c.rateMu.Unlock()

	err = CheckResponse(resp)
	if err != nil {
		// even though there was an error, we still return the response
		// in case the caller wants to inspect it further
		return response, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err == io.EOF {
				err = nil // ignore EOF errors caused by empty response body
			}
		}
	}

	return response, err
}

// checkRateLimitBeforeDo does not make any network calls, but uses existing knowledge from
// current client state in order to quickly check if *RateLimitError can be immediately returned
// from Client.Do, and if so, returns it so that Client.Do can skip making a network API call unnecessarily.
// Otherwise it returns nil, and Client.Do should proceed normally.
func (c *Client) checkRateLimitBeforeDo(req *http.Request) error {

	// TODO:
	// c.rateMu.Lock()
	// // rate := c.rateLimits[rateLimitCategory]
	// c.rateMu.Unlock()
	// if !rate.Reset.Time.IsZero() && rate.Remaining == 0 && time.Now().Before(rate.Reset.Time) {
	// 	// Create a fake response.
	// 	resp := &http.Response{
	// 		Status:     http.StatusText(http.StatusForbidden),
	// 		StatusCode: http.StatusForbidden,
	// 		Request:    req,
	// 		Header:     make(http.Header),
	// 		Body:       ioutil.NopCloser(strings.NewReader("")),
	// 	}
	// 	return &RateLimitError{
	// 		Rate:     rate,
	// 		Response: resp,
	// 		Message:  fmt.Sprintf("API rate limit of %v still exceeded until %v, not making remote request.", rate.Limit, rate.Reset.Time),
	// 	}
	// }

	return nil
}

// CheckResponse checks the API response for errors, and returns them if
// present.  A response is considered an error if it has a status code outside
// the 200 range.  API error responses are expected to have either no response
// body, or a JSON response body that maps to ErrorResponse.  Any other
// response body will be silently ignored.
//
// The error type will be *RateLimitError for rate limit exceeded errors,
// and *TwoFactorAuthError for two-factor authentication errors.

// TODO - check rate limit
// TODO - check un-authorized
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, &errorResponse.ErrorDetail)
	}
	return errorResponse
	// switch {

	// case r.StatusCode == http.StatusNotFound:

	// case r.StatusCode == http.StatusUnauthorized:

	// }
	// errorResponse := &ErrorResponse{Response: r}
	// data, err := ioutil.ReadAll(r.Body)
	// if err == nil && data != nil {
	// 	json.Unmarshal(data, errorResponse)
	// }
	// switch {
	// case r.StatusCode == http.StatusUnauthorized:
	// 	return error.Error("unauthroized") // TODO
	// case r.StatusCode == http.StatusForbidden && r.Header.Get(headerRateRemaining) == "0":
	// 	return &RateLimitError{
	// 		Rate:     parseRate(r),
	// 		Response: errorResponse.Response,
	// 		Message:  errorResponse.Message,
	// 	}
	// default:
	// 	return errorResponse
	// }
	return nil
}

type APIError struct {
	ErrorCode    string `json:"errorCode"`
	ErrorSummary string `json:"errorSummary"`
	ErrorLink    string `json:"errorLink"`
	ErrorID      string `json:"errorId"`
	ErrorCauses  []struct {
		ErrorSummary string `json:"errorSummary"`
	} `json:"errorCauses"`
}

type ErrorResponse struct {
	Response    *http.Response //
	ErrorDetail APIError
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("HTTP Method: %v - URL: %v: - HTTP Status Code: %d, OKTA Error Code: %v, OKTA Error Summary: %v",
		r.Response.Request.Method, r.Response.Request.URL, r.Response.StatusCode, r.ErrorDetail.ErrorCode, r.ErrorDetail.ErrorSummary)
}

// Code stolen from Github api libary
// Stringify attempts to create a reasonable string representation of types in
// the library.  It does things like resolve pointers to their values
// and omits struct fields with nil values.
func Stringify(message interface{}) string {
	var buf bytes.Buffer
	v := reflect.ValueOf(message)
	stringifyValue(&buf, v)
	return buf.String()
}

// stringifyValue was heavily inspired by the goprotobuf library.

func stringifyValue(w io.Writer, val reflect.Value) {
	if val.Kind() == reflect.Ptr && val.IsNil() {
		w.Write([]byte("<nil>"))
		return
	}

	v := reflect.Indirect(val)

	switch v.Kind() {
	case reflect.String:
		fmt.Fprintf(w, `"%s"`, v)
	case reflect.Slice:
		w.Write([]byte{'['})
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				w.Write([]byte{' '})
			}

			stringifyValue(w, v.Index(i))
		}

		w.Write([]byte{']'})
		return
	case reflect.Struct:
		if v.Type().Name() != "" {
			w.Write([]byte(v.Type().String()))
		}
		w.Write([]byte{'{'})

		var sep bool
		for i := 0; i < v.NumField(); i++ {
			fv := v.Field(i)
			if fv.Kind() == reflect.Ptr && fv.IsNil() {
				continue
			}
			if fv.Kind() == reflect.Slice && fv.IsNil() {
				continue
			}

			if sep {
				w.Write([]byte(", "))
			} else {
				sep = true
			}

			w.Write([]byte(v.Type().Field(i).Name))
			w.Write([]byte{':'})
			stringifyValue(w, fv)
		}

		w.Write([]byte{'}'})
	default:
		if v.CanInterface() {
			fmt.Fprint(w, v.Interface())
		}
	}
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash.  If
// specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set(headerAuthorization, fmt.Sprintf(headerAuthorizationFormat, c.apiKey))

	if body != nil {
		req.Header.Set("Content-Type", mediaTypeJSON)
	}

	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	return req, nil
}

// addOptions adds the parameters in opt as URL query parameters to s.  opt
// must be a struct whose fields may contain "url" tags.
func addOptions(s string, opt interface{}) (string, error) {
	v := reflect.ValueOf(opt)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opt)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}
