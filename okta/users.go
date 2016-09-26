package okta

import (
	"fmt"
	"net/url"
	"time"
)

const (
	profileEmailFilter       = "profile.email"
	profileLoginFilter       = "profile.login"
	profileStatusFilter      = "status"
	profileIDFilter          = "id"
	profileFirstNameFilter   = "profile.firstName"
	profileLastNameFilter    = "profile.lastName"
	profileLastUpdatedFilter = "lastUpdated"
	// UserStatusActive is a  constant to represent OKTA User State returned by the API
	UserStatusActive = "ACTIVE"
	// UserStatusStaged is a  constant to represent OKTA User State returned by the API
	UserStatusStaged = "STAGED"
	// UserStatusProvisioned is a  constant to represent OKTA User State returned by the API
	UserStatusProvisioned = "PROVISIONED"
	// UserStatusRecovery is a  constant to represent OKTA User State returned by the API
	UserStatusRecovery = "RECOVERY"
	// UserStatusLockedOut is a  constant to represent OKTA User State returned by the API
	UserStatusLockedOut = "LOCKED_OUT"
	// UserStatusPasswordExpired is a  constant to represent OKTA User State returned by the API
	UserStatusPasswordExpired = "PASSWORD_EXPIRED"
	// UserStatusSuspended is a  constant to represent OKTA User State returned by the API
	UserStatusSuspended = "SUSPENDED"
	// UserStatusDeprovisioned is a  constant to represent OKTA User State returned by the API
	UserStatusDeprovisioned = "DEPROVISIONED"

	oktaFilterTimeFormat = "2006-01-02T15:05:05.000Z"
)

// UsersService handles communication with the User data related
// methods of the OKTA API.
type UsersService service

type provider struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type recoveryQuestion struct {
	Question string `json:"question"`
}

type credentials struct {
	Password         struct{}         `json:"password"`
	Provider         provider         `json:"provider"`
	RecoveryQuestion recoveryQuestion `json:"recovery_question"`
}

type userProfile struct {
	Email       string `json:"email"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Login       string `json:"login"`
	MobilePhone string `json:"mobilePhone"`
	SecondEmail string `json:"secondEmail"`
	PsEmplid    string `json:"psEmplid"`
	NickName    string `json:"nickname"`
	DisplayName string `json:"displayName"`

	ProfileURL        string `json:"profileUrl"`
	PreferredLanguage string `json:"preferredLanguage"`
	UserType          string `json:"userType"`
	Organization      string `json:"organization"`
	Title             string `json:"title"`
	Division          string `json:"division"`
	Department        string `json:"department"`
	CostCenter        string `json:"costCenter"`
	EmployeeNumber    string `json:"employeeNumber"`
	PrimaryPhone      string `json:"primaryPhone"`
	StreetAddress     string `json:"streetAddress"`
	City              string `json:"city"`
	State             string `json:"state"`
	ZipCode           string `json:"zipCode"`
	CountryCode       string `json:"countryCode"`
}

type userLinks struct {
	ChangePassword struct {
		Href string `json:"href"`
	} `json:"changePassword"`
	ChangeRecoveryQuestion struct {
		Href string `json:"href"`
	} `json:"changeRecoveryQuestion"`
	Deactivate struct {
		Href string `json:"href"`
	} `json:"deactivate"`
	ExpirePassword struct {
		Href string `json:"href"`
	} `json:"expirePassword"`
	ForgotPassword struct {
		Href string `json:"href"`
	} `json:"forgotPassword"`
	ResetFactors struct {
		Href string `json:"href"`
	} `json:"resetFactors"`
	ResetPassword struct {
		Href string `json:"href"`
	} `json:"resetPassword"`
}

// User is a struct that represents a user object from OKTA.
type User struct {
	Activated       string      `json:"activated"`
	Created         string      `json:"created"`
	Credentials     credentials `json:"credentials"`
	ID              string      `json:"id"`
	LastLogin       string      `json:"lastLogin"`
	LastUpdated     string      `json:"lastUpdated"`
	PasswordChanged string      `json:"passwordChanged"`
	Profile         userProfile `json:"profile"`
	Status          string      `json:"status"`
	StatusChanged   string      `json:"statusChanged"`
	Links           userLinks   `json:"_links"`
	MFAFactors      []userMFAFactor
	Groups          []Group
}

type userMFAFactor struct {
	ID          string    `json:"id"`
	FactorType  string    `json:"factorType"`
	Provider    string    `json:"provider"`
	VendorName  string    `json:"vendorName"`
	Status      string    `json:"status"`
	Created     time.Time `json:"created"`
	LastUpdated time.Time `json:"lastUpdated"`
	Profile     struct {
		CredentialID string `json:"credentialId"`
	} `json:"profile"`
}

func (u User) String() string {
	return Stringify(u)
	// return fmt.Sprintf("ID: %v \tLogin: %v", u.ID, u.Profile.Login)
}

// GetByID returns a user object for a specific OKTA ID.
// Generally the id input string is the cryptic OKTA key value from User.ID. However, the OKTA API may accept other values like "me", or login shortname
func (s *UsersService) GetByID(id string) (*User, *Response, error) {
	u := fmt.Sprintf("users/%v", id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	user := new(User)
	resp, err := s.client.Do(req, user)
	if err != nil {
		return nil, resp, err
	}

	return user, resp, err
}

// UserListFilterOptions is a struct that you can populate which will "filter" user searches
// the exported struct fields should allow you to do different filters based on what is allowed in the OKTA API.
//  The filter OKTA API is limited in the fields it can search
//  NOTE: In the current form you can't add parenthesis and ordering
// OKTA API Supports only a limited number of properties:
// status, lastUpdated, id, profile.login, profile.email, profile.firstName, and profile.lastName.
// http://developer.okta.com/docs/api/resources/users.html#list-users-with-a-filter
type UserListFilterOptions struct {
	Limit         int    `url:"limit,omitempty"`
	EmailEqualTo  string `url:"-"`
	LoginEqualTo  string `url:"-"`
	StatusEqualTo string `url:"-"`
	IDEqualTo     string `url:"-"`

	FirstNameEqualTo string `url:"-"`
	LastNameEqualTo  string `url:"-"`
	//  API documenation says you can search with "starts with" but these don't work

	// FirstNameStartsWith    string    `url:"-"`
	// LastNameStartsWith     string    `url:"-"`
	LastUpdatedGreaterThan time.Time `url:"-"`
	LastUpdatedLessThan    time.Time `url:"-"`
	// This will be built by internal - may not need to export
	FilterString string   `url:"filter,omitempty"`
	NextURL      *url.URL `url:"-"`
}

// PopulateGroups will populate the groups a user is a member of. You pass in a pointer to an existing users
func (s *UsersService) PopulateGroups(user *User) (*Response, error) {
	u := fmt.Sprintf("users/%v/groups", user.ID)
	req, err := s.client.NewRequest("GET", u, nil)

	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, &user.Groups)
	if err != nil {
		return resp, err
	}

	return resp, err
}

// List users with status of LOCKED_OUT
// filter=status eq "LOCKED_OUT"
// List users updated after 06/01/2013 but before 01/01/2014
// filter=lastUpdated gt "2013-06-01T00:00:00.000Z" and lastUpdated lt "2014-01-01T00:00:00.000Z"
// List users updated after 06/01/2013 but before 01/01/2014 with a status of ACTIVE
// filter=lastUpdated gt "2013-06-01T00:00:00.000Z" and lastUpdated lt "2014-01-01T00:00:00.000Z" and status eq "ACTIVE"
// TODO - Currently no way to do parenthesis
// List users updated after 06/01/2013 but with a status of LOCKED_OUT or RECOVERY
// filter=lastUpdated gt "2013-06-01T00:00:00.000Z" and (status eq "LOCKED_OUT" or status eq "RECOVERY")

// OTKA API docs: http://developer.okta.com/docs/api/resources/users.html#list-users-with-a-filter

func appendToFilterString(currFilterString string, appendFilterKey string, appendFilterOperator string, appendFilterValue string) (rs string) {
	if currFilterString != "" {
		rs = fmt.Sprintf("%v and %v %v \"%v\"", currFilterString, appendFilterKey, appendFilterOperator, appendFilterValue)
	} else {
		rs = fmt.Sprintf("%v %v \"%v\"", appendFilterKey, appendFilterOperator, appendFilterValue)
	}

	return rs
}

// ListWithFilter will use the input UserListFilterOptions to find users and return a paged result set
func (s *UsersService) ListWithFilter(opt *UserListFilterOptions) ([]*User, *Response, error) {
	var u string
	var err error

	if opt.NextURL != nil {
		u = opt.NextURL.String()
		fmt.Printf("ListWithFilter NextURL: %v\n", u)
	} else {
		if opt.EmailEqualTo != "" {
			opt.FilterString = appendToFilterString(opt.FilterString, profileEmailFilter, filterEqualOperator, opt.EmailEqualTo)
		}
		if opt.LoginEqualTo != "" {
			opt.FilterString = appendToFilterString(opt.FilterString, profileLoginFilter, filterEqualOperator, opt.LoginEqualTo)
		}

		if opt.StatusEqualTo != "" {
			opt.FilterString = appendToFilterString(opt.FilterString, profileStatusFilter, filterEqualOperator, opt.StatusEqualTo)
		}

		if opt.IDEqualTo != "" {
			opt.FilterString = appendToFilterString(opt.FilterString, profileIDFilter, filterEqualOperator, opt.IDEqualTo)
		}

		if opt.FirstNameEqualTo != "" {
			opt.FilterString = appendToFilterString(opt.FilterString, profileFirstNameFilter, filterEqualOperator, opt.FirstNameEqualTo)
		}

		if opt.LastNameEqualTo != "" {
			opt.FilterString = appendToFilterString(opt.FilterString, profileLastNameFilter, filterEqualOperator, opt.LastNameEqualTo)
		}

		//  API documenation says you can search with "starts with" but these don't work
		// if opt.FirstNameStartsWith != "" {
		// 	opt.FilterString = appendToFilterString(opt.FilterString, profileFirstNameFilter, filterStartsWithOperator, opt.FirstNameStartsWith)
		// }

		// if opt.LastNameStartsWith != "" {
		// 	opt.FilterString = appendToFilterString(opt.FilterString, profileLastNameFilter, filterStartsWithOperator, opt.LastNameStartsWith)
		// }

		if !opt.LastUpdatedGreaterThan.IsZero() {
			opt.FilterString = appendToFilterString(opt.FilterString, profileLastUpdatedFilter, filterGreaterThanOperator, opt.LastUpdatedGreaterThan.UTC().Format(oktaFilterTimeFormat))
		}

		if !opt.LastUpdatedLessThan.IsZero() {
			fmt.Printf("-------LastUpdatedLessThan in 8601: %v \n", opt.LastUpdatedLessThan.UTC().Format(oktaFilterTimeFormat))

			opt.FilterString = appendToFilterString(opt.FilterString, profileLastUpdatedFilter, filterLessThanOperator, opt.LastUpdatedLessThan.UTC().Format(oktaFilterTimeFormat))
		}

		if opt.Limit == 0 {
			opt.Limit = defaultLimit
		}

		u, err = addOptions("users", opt)

	}

	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	users := new([]*User)
	resp, err := s.client.Do(req, users)
	if err != nil {
		return nil, resp, err
	}

	return *users, resp, err
}
