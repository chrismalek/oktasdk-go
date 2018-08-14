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
		ID:          "0oa62bfdiumsUndnZ0h7",
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

func TestPolicyCreate(t *testing.T) {
	setupTestIdentityProvider()

	t.Run("Password", func(t *testing.T) {
		testPolicyCreate(t, testInputPassPolicy, testPassPolicy)
	})
	t.Run("SignOn", func(t *testing.T) {
		testPolicyCreate(t, testInputSignonPolicy, testSignonPolicy)
	})
}

func testRuleCreate(t *testing.T, inputrule interface{}, rule *Rule) {

	setup()
	defer teardown()

	temp, err := json.Marshal(rule)
	if err != nil {
		t.Errorf("Policies.CreateiPolicyRule json Marshall returned error: %v", err)
	}
	ruleTestJSONString := string(temp)

	mux.HandleFunc("/policies/00pedv3qclXeC2aFH0h7/rules", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testAuthHeader(t, r)
		fmt.Fprint(w, ruleTestJSONString)
	})

	outputrule, _, err := client.Policies.CreatePolicyRule("00pedv3qclXeC2aFH0h7", inputrule)
	if err != nil {
		t.Errorf("Policies.CreatePolicyRule returned error: %v", err)
	}
	if !reflect.DeepEqual(outputrule, rule) {
		t.Errorf("client.Policies.CreatePolicy returned \n\t%+v, want \n\t%+v\n", outputrule, rule)
	}
}

func TestRuleCreate(t *testing.T) {
	setupTestRules()

	t.Run("Password", func(t *testing.T) {
		testRuleCreate(t, testInputPassRule, testPassRule)
	})
	t.Run("SignOn", func(t *testing.T) {
		testRuleCreate(t, testInputSignonRule, testSignonRule)
	})
}
