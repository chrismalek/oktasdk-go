package okta

import (
	"fmt"
	"net/url"
	"time"
)

// GroupsService handles communication with the Groups data related
// methods of the OKTA API.
type GroupsService service

// TODO: GroupSearch

// Group represents the Group Object from the OKTA API
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

// GetByID gets a group from OKTA by the Gropu ID. An error is returned if the group is not found
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

// GetUsers returns the members in a group
//   Pass in an optional GroupFilterOptions struct to filter the results
//   The Users in the group are returned
func (g *GroupsService) GetUsers(groupID string, opt *GroupFilterOptions) (users []User, resp *Response, err error) {
	pagesRetreived := 0
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
	resp, err = g.client.Do(req, &users)

	if err != nil {
		return nil, resp, err
	}

	pagesRetreived++
	if (opt.NumberOfPages > 0 && pagesRetreived < opt.NumberOfPages) || opt.GetAllPages {

		for {

			if pagesRetreived == opt.NumberOfPages {
				break
			}
			if resp.NextURL != nil {

				var userPage []User
				pageOpts := new(GroupFilterOptions)
				pageOpts.NextURL = resp.NextURL
				pageOpts.Limit = opt.Limit
				pageOpts.NumberOfPages = 1

				userPage, resp, err = g.GetUsers(groupID, pageOpts)
				if err != nil {
					return users, resp, err
				} else {
					users = append(users, userPage...)
					pagesRetreived++
				}
			} else {
				break
			}

		}
	}

	return users, resp, err
}

// GroupFilterOptions is a struct that you populate which will limit or control group fetches and searches
//  The values here will coorelate to the search filtering allowed in the OKTA API. These values are turned into Query Parameters
type GroupFilterOptions struct {
	Limit         int      `url:"limit,omitempty"`
	NextURL       *url.URL `url:"-"`
	GetAllPages   bool     `url:"-"`
	NumberOfPages int      `url:"-"`
}
