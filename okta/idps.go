package okta

import (
	"time"
)

type IdentityProvidersService service

type IdentityProvider struct {
	ID          string    `json:"id"`
	Type        string    `json:"type"`
	Status      string    `json:"status"`
	Name        string    `json:"name"`
	Created     time.Time `json:"created"`
	LastUpdated time.Time `json:"lastUpdated"`
	Protocol    struct {
		Type      string `json:"type"`
		Endpoints struct {
			Authorization struct {
				Url     string `json:"url"`
				Binding string `json:"binding"`
			} `json:"authorization"`
			Token struct {
				Url     string `json:"url"`
				Binding string `json:"binding"`
			}
		} `json:"endpoints"`
		Scopes      []string `json:"scopes"`
		Credentials struct {
			Client struct {
				ClientID     string `json:"client_id"`
				ClientSecret string `json:"client_secret"`
			} `json:"client"`
		} `json:"credentials"`
	} `json:"protocol"`
	Policy struct {
		Provisioning struct {
			Action        string `json:"action"`
			ProfileMaster bool   `json:"profileMaster"`
			Groups        struct {
				Action string `json:"action"`
			} `json:"groups"`
			Conditions struct {
				Deprovisioned struct {
					Action string `json:"action"`
				} `json:"deprovisioned"`
				Suspended struct {
					Action string `json:"action"`
				} `json:"suspended"`
			} `json:"conditions"`
		} `json:"provisioning"`
		AccountLink struct {
			Filter string `json:"filter"`
			Action string `json:"action"`
		} `json:"accountLink"`
		Subject struct {
			UserNameTemplate struct {
				Template string `json:"template"`
			} `json:"userNameTemplate"`
			Filter    string `json:"filter"`
			MatchType string `json:"matchType"`
		} `json:"subject"`
		MaxClockSkew int `json:"maxClockSkew"`
	} `json:"policy"`
	Links struct {
		Authorize struct {
			Href      string `json:"href"`
			Templated bool   `json:"templated"`
			Hints     struct {
				Allow []string `json:"allow"`
			} `json:"hints"`
		} `json:"authorize"`
		ClientRedirectUri struct {
			Href  string `json:"href"`
			Hints struct {
				Allow []string `json:"allow"`
			} `json:"hints"`
		} `json:"clientRedirectUri"`
	} `json:"_links"`
}
