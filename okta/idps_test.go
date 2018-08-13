package okta

import (
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

func TestIdentityProvider(t *testing.T) {
	setup()
	defer teardown()
	setupTestIdentityProvider()
}
