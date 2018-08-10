package okta

import (
	"fmt"
	"net/url"
	"time"
)


type IdentityProvidersService service

type IdentityProvider struct {
	ID      string	  `json:"id"`
	Type    string	  `json:"type"`
	Status  string	  `json:"status"`
	Name    string	  `json:"name"`
	Created time.Time `json:"created"`
	LastUpdated time.Time `json:"lastUpdated"`
	Protocol struct {
		Type string `json:"type"`
		Endpoints struct {
			Authorization struct {
				Url string `json:"url"`
				Binding string `json:"binding"`
			} `json:"authorization"`
			Token struct {
				Url string `json:"url"`
				Binding string `json:"binding"`
			}
		} `json:"endpoints"`
		Scopes []interface{} `json:"scopes"`
		Credentials struct {
			Client struct {
				ClientID string `json:"client_id"`
				ClientSecret string `json:"client_secret"`
			} `json:"client"`
		} `json:"credentials"`
	} `json:"protocol"`
}
