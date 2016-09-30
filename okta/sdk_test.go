package okta

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
	"time"
)

var (
	// mux is the HTTP request multiplexer used with the test server.
	mux *http.ServeMux

	// client is the  client being tested.
	client *Client

	// server is a test HTTP server used to provide mock API responses.
	server *httptest.Server

	orgName string
	token   string
)

const (
	testServerOrg = "test-org"
	testToken     = "marked.swishy.eighteen.noticing.styptic"
)

// setup sets up a test HTTP server along with a okta.Client that is
// configured to talk to that test server.  Tests should register handlers on
// mux which provide mock responses for the API method being tested.
func setup() {

	// test server
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	//  client configured to use test server
	// non-production server
	client = NewClient(nil, testServerOrg, testToken, false)

	client.BaseURL, _ = url.Parse(server.URL)

}

func teardown() {
	server.Close()

}

func TestVerifyPreviewClientSetup(t *testing.T) {
	client := NewClient(nil, testServerOrg, testToken, false)
	wantURL := fmt.Sprintf("https://%v.oktapreview.com/api/v1/", testServerOrg)
	if client.BaseURL.String() != wantURL {
		t.Errorf("client.BaseURL should be %v but got %v", client.BaseURL, wantURL)
	}

}

func TestVerifyProdClientSetup(t *testing.T) {
	client := NewClient(nil, testServerOrg, testToken, true)
	wantURL := fmt.Sprintf("https://%v.okta.com/api/v1/", testServerOrg)
	if client.BaseURL.String() != wantURL {
		t.Errorf("client.BaseURL should be %v but got %v", client.BaseURL, wantURL)
	}

}

func testAuthHeader(t *testing.T, r *http.Request) {
	want := fmt.Sprintf("SSWS %v", testToken)
	if value := r.Header.Get("Authorization"); want != value {
		t.Errorf("Authorization Header %s, want: %s", value, want)
	}
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func TestRateLimitRateExceededError(t *testing.T) {
	setup()
	defer teardown()
	// Returning a rate response which is below our threshold
	// The SDK should through and error dependin on how the client is configured.
	headerRateLimitWant := 1200
	headerRateRemainingWant := 20

	mux.HandleFunc("/users/me", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add(headerRateLimit, strconv.Itoa(headerRateLimitWant))
		w.Header().Add(headerRateRemaining, strconv.Itoa(headerRateRemainingWant))
		thirtySecondsFromNow := time.Now().Add(30 * time.Second)
		w.Header().Add(headerRateReset, strconv.FormatInt(thirtySecondsFromNow.Unix(), 10))
	})

	client = NewClient(nil, testServerOrg, testToken, false)
	client.BaseURL, _ = url.Parse(server.URL)
	client.PauseOnRateLimit = false

	if client.mostRecentRate.Remaining != 0 {
		t.Errorf("client.mostRecentRate.Remaining should be initialized as Zero. Got: %v\n", client.mostRecentRate.Remaining)
	}

	u := fmt.Sprintf("users/me")
	req, err := client.NewRequest("GET", u, nil)

	if err != nil {
		t.Errorf("Error Creating Client: %v\n", err)
	}

	_, err = client.Do(req, nil)

	if err != nil {
		t.Errorf("Error doing GET Test: %v\n", err)
	}

	if client.mostRecentRate.Remaining != headerRateRemainingWant {

		t.Errorf("client.mostRecentRate.Remaining was not cached. Expected %v, Got: %v", client.mostRecentRate.Remaining, headerRateRemainingWant)
	}

	if client.mostRecentRate.RatePerMinuteLimit != headerRateLimitWant {
		t.Errorf("client.mostRecentRate.RatePerMinuteLimit was not cached. Expected %v, Got: %v", client.mostRecentRate.RatePerMinuteLimit, headerRateLimitWant)
	}
	// Second Call should return an error becasue it has cached the values
	_, err = client.Do(req, nil)

	if err == nil {
		t.Errorf("Expected Rate Limit Error To be Returned. However, No Error was created.\n")
	}

}
