package okta

import (
	"fmt"
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

// GetIdentityProvider: Get an IdP
// Requires IdentityProvider ID from IdentityProvider object
func (p *IdentityProvidersService) GetIdentityProvider(id string) (*IdentityProvider, *Response, error) {
	u := fmt.Sprintf("idps/%v", id)
	req, err := p.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	idp := new(IdentityProvider)
	resp, err := p.client.Do(req, idp)
	if err != nil {
		return nil, resp, err
	}

	return idp, resp, err
}

// CreateIdentityprovider: Create a identityprovider
// You must pass in the Identityprovider object created from the desired input identityprovider
func (p *IdentityProvidersService) CreateIdentityProvider(idp interface{}) (*IdentityProvider, *Response, error) {
	u := fmt.Sprintf("idps")
	req, err := p.client.NewRequest("POST", u, idp)
	if err != nil {
		return nil, nil, err
	}

	newIdp := new(IdentityProvider)
	resp, err := p.client.Do(req, newIdp)
	if err != nil {
		return nil, resp, err
	}

	return newIdp, resp, err
}

// UpdateIdentityProvider: Update a policy
// Requires IdentityProvider ID from IdentityProvider object & IdentityProvider object from the desired input policy
func (p *IdentityProvidersService) UpdateIdentityProvider(id string, idp interface{}) (*IdentityProvider, *Response, error) {
	u := fmt.Sprintf("idps/%v", id)
	req, err := p.client.NewRequest("PUT", u, idp)
	if err != nil {
		return nil, nil, err
	}

	updateIdentityProvider := new(IdentityProvider)
	resp, err := p.client.Do(req, updateIdentityProvider)
	if err != nil {
		return nil, resp, err
	}

	return updateIdentityProvider, resp, err
}
