package okta

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

var testIdentityProvider *IdentityProvider

func setupTestIdentityProvider() {
	hmm, _ := time.Parse("2006-01-02T15:04:05.000Z", "2018-02-16T19:59:05.000Z")

	testIdentityProvider = &IdentityProvider{
		Type:        "GOOGLE",
		Status:      "ACTIVE",
		Name:        "Google",
		Created:     hmm,
		LastUpdated: hmm,
	}
	testIdentityProvider.Protocol.Type = "OIDC"
	testIdentityProvider.Protocol.Endpoints.Authorization.Url = "https://accounts.google.com/o/oauth2/auth"
	testIdentityProvider.Protocol.Endpoints.Authorization.Binding = "HTTP-REDIRECT"
	testIdentityProvider.Protocol.Endpoints.Token.Url = "https://www.googleapis.com/oauth2/v3/token"
	testIdentityProvider.Protocol.Endpoints.Token.Binding = "HTTP-POST"
	testIdentityProvider.Protocol.Scopes = []string{"profile email openid"}
	testIdentityProvider.Protocol.Credentials.Client.ClientID = "your-client-id"
	testIdentityProvider.Protocol.Credentials.Client.ClientSecret = "your-client-secret"
	testIdentityProvider.Policy.Provisioning.Action = "AUTO"
	testIdentityProvider.Policy.Provisioning.ProfileMaster = true
	testIdentityProvider.Policy.Provisioning.Groups.Action = "NONE"
	testIdentityProvider.Policy.Provisioning.Conditions.Deprovisioned.Action = "NONE"
	testIdentityProvider.Policy.Provisioning.Conditions.Suspended.Action = "NONE"
	// testIdentityProvider.Policy.AccountLink.Filter = null
	testIdentityProvider.Policy.AccountLink.Action = "AUTO"
	testIdentityProvider.Policy.Subject.UserNameTemplate.Template = "idpuser.userPrincipalName"
	// testIdentityProvider.Policy.Subject.Filter = null
	testIdentityProvider.Policy.Subject.MatchType = "USERNAME"
	testIdentityProvider.Policy.MaxClockSkew = 0
	testIdentityProvider.Links.Authorize.Href = "https://{yourOktaDomain}/oauth2/v1/authorize?idp=0oa62bfdiumsUndnZ0h7&client_id={clientId}&response_type={responseType}&response_mode={responseMode}&scope={scopes}&redirect_uri={redirectUri}&state={state}"
	testIdentityProvider.Links.Authorize.Templated = true
	testIdentityProvider.Links.Authorize.Hints.Allow = []string{"GET"}
	testIdentityProvider.Links.ClientRedirectUri.Href = "https://{yourOktaDomain}/oauth2/v1/authorize/callback"
	testIdentityProvider.Links.ClientRedirectUri.Hints.Allow = []string{"POST"}
}

func TestGetIdentityProvider(t *testing.T) {
	setup()
	defer teardown()
	setupTestIdentityProvider()
	testIdentityProvider.ID = "0oa62bfdiumsUndnZ0h7"

	temp, err := json.Marshal(testIdentityProvider)
	if err != nil {
		t.Errorf("IdentityProviders.GetIdentityProvider json Marshall returned error: %v", err)
	}
	IdentityProviderTestJSONString := string(temp)

	mux.HandleFunc("/idps/0oa62bfdiumsUndnZ0h7", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testAuthHeader(t, r)
		fmt.Fprint(w, IdentityProviderTestJSONString)
	})

	outputIdentityProvider, _, err := client.IdentityProviders.GetIdentityProvider("0oa62bfdiumsUndnZ0h7")
	if err != nil {
		t.Errorf("IdentityProviders.GetIdentityProvider returned error: %v", err)
	}
	if !reflect.DeepEqual(outputIdentityProvider, testIdentityProvider) {
		t.Errorf("client.IdentityProviders.GetIdentityProvider returned \n\t%+v, want \n\t%+v\n", outputIdentityProvider, testIdentityProvider)
	}
}

func TestIdentityProviderCreate(t *testing.T) {

	setup()
	defer teardown()
	setupTestIdentityProvider()

	temp, err := json.Marshal(testIdentityProvider)

	if err != nil {
		t.Errorf("IdentityProviders.CreateIdentityProvider json Marshall returned error: %v", err)
	}

	IdentityProviderTestJSONString := string(temp)

	mux.HandleFunc("/idps", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testAuthHeader(t, r)
		fmt.Fprint(w, IdentityProviderTestJSONString)
	})

	outputIdentityProvider, _, err := client.IdentityProviders.CreateIdentityProvider(testIdentityProvider)
	if err != nil {
		t.Errorf("IdentityProvider.CreateIdentityProvider returned error: %v", err)
	}
	if !reflect.DeepEqual(outputIdentityProvider, testIdentityProvider) {
		t.Errorf("client.IdentityProviders.CreateIdentityProvider returned \n\t%+v, want \n\t%+v\n", outputIdentityProvider, testIdentityProvider)
	}
}

func TestIdentityProviderUpdate(t *testing.T) {

	setup()
	defer teardown()
	setupTestIdentityProvider()
	testIdentityProvider.ID = "0oa62bfdiumsUndnZ0h7"

	testIdentityProvider.Name = "Something Completely Different"
	temp, err := json.Marshal(testIdentityProvider)
	if err != nil {
		t.Errorf("IdentityProviders.UpdateIdentityProvider json Marshall returned error: %v", err)
	}
	updateTestJSONString := string(temp)

	mux.HandleFunc("/idps/0oa62bfdiumsUndnZ0h7", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		testAuthHeader(t, r)
		fmt.Fprint(w, updateTestJSONString)
	})

	outputIdentityProvider, _, err := client.IdentityProviders.UpdateIdentityProvider("0oa62bfdiumsUndnZ0h7", testIdentityProvider)
	if err != nil {
		t.Errorf("IdentityProviders.UpdateIdentityProvider returned error: %v", err)
	}
	if !reflect.DeepEqual(outputIdentityProvider.Name, testIdentityProvider.Name) {
		t.Errorf("client.IdentityProviders.UpdateIdentityProvider returned \n\t%+v, want \n\t%+v\n", outputIdentityProvider.Name, testIdentityProvider.Name)
	}
}
