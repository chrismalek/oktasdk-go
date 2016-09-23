package okta

import (
	"fmt"
	"net/url"
	"time"
)

type GroupsService service
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
	return fmt.Sprintf("Group:(ID: {%v} - Type: {%v} - Group Name: {%v})\n", g.ID, g.Type, g.Profile.Name)
}

func (g *GroupsService) GetByID(groupID string) (*Group, *Response, error) {

	u := fmt.Sprintf("groups/%v", groupID)
	req, err := g.client.NewRequest("GET", u, nil)

	if err != nil {
		return nil, nil, err
	}

	group := new(Group)

	resp, err := g.client.Do(req, group)

	if err != nil {
		return nil, resp, err
	}

	return group, resp, err
}

func (g *GroupsService) GetUsers(groupID string, opt *GroupFilterOptions) (users []User, resp *Response, err error) {

	var u string
	if opt.NextURL != nil {
		u = opt.NextURL.String()
	} else {
		u = fmt.Sprintf("groups/%v/users", groupID)

		if opt.Limit == 0 {
			opt.Limit = defaultLimit
		}

		u, _ = addOptions(u, opt)

	}

	req, err := g.client.NewRequest("GET", u, nil)

	if err != nil {
		return nil, nil, err
	}
	// users := new([]User)
	resp, err = g.client.Do(req, &users)

	if err != nil {
		return nil, resp, err
	}

	return users, resp, err
}

type GroupFilterOptions struct {
	Limit   int      `url:"limit,omitempty"`
	NextURL *url.URL `url:"-"`
}
