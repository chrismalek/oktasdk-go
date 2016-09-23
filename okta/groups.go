package okta

import (
	"fmt"
	"time"
)

type Group struct {
	ID                    string    `json:"id"`
	Created               time.Time `json:"created"`
	LastUpdated           time.Time `json:"lastUpdated"`
	LastMembershipUpdated time.Time `json:"lastMembershipUpdated"`
	ObjectClass           []string  `json:"objectClass"`
	Type                  string    `json:"type"`
	Profile               struct {
		Name                       string `json:"name"`
		Description                string `json:"description"`
		SamAccountName             string `json:"samAccountName"`
		Dn                         string `json:"dn"`
		WindowsDomainQualifiedName string `json:"windowsDomainQualifiedName"`
		ExternalID                 string `json:"externalId"`
	} `json:"profile"`
	Links struct {
		Logo []struct {
			Name string `json:"name"`
			Href string `json:"href"`
			Type string `json:"type"`
		} `json:"logo"`
		Users struct {
			Href string `json:"href"`
		} `json:"users"`
		Apps struct {
			Href string `json:"href"`
		} `json:"apps"`
	} `json:"_links"`
}

func (g Group) String() string {
	// return Stringify(g)
	return fmt.Sprintf("ID: {%v} - Type: {%v} - Group Name: {%v}\n", g.ID, g.Type, g.Profile.Name)
}
