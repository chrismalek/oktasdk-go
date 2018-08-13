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
}

func TestIdentityProvider(t *testing.T) {
	setup()
	defer teardown()
	setupTestIdentityProvider()
}
