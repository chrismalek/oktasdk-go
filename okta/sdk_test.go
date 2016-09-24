package okta

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
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
	TEST_SERVER_ORG = "test-org"
	TEST_TOKEN      = "marked.swishy.eighteen.noticing.styptic"
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
	client = NewClient(nil, orgName, token, false)

	client.BaseURL, _ = url.Parse(server.URL)

}

func TestVerifyPreviewClientSetup(t *testing.T) {
	client := NewClient(nil, TEST_SERVER_ORG, TEST_TOKEN, false)
	wantURL := fmt.Sprintf("https://%v.oktapreview.com/api/v1/", TEST_SERVER_ORG)
	if client.BaseURL.String() != wantURL {
		t.Errorf("client.BaseURL should be %v but got %v", client.BaseURL, wantURL)
	}

}

func TestVerifyProdClientSetup(t *testing.T) {
	client := NewClient(nil, TEST_SERVER_ORG, TEST_TOKEN, true)
	wantURL := fmt.Sprintf("https://%v.okta.com/api/v1/", TEST_SERVER_ORG)
	if client.BaseURL.String() != wantURL {
		t.Errorf("client.BaseURL should be %v but got %v", client.BaseURL, wantURL)
	}

}
