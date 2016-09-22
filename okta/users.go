package okta

import (
	"fmt"
	"time"
)

type UsersService service

type User struct {
	Activated   string `json:"activated"`
	Created     string `json:"created"`
	Credentials struct {
		Password struct{} `json:"password"`
		Provider struct {
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"provider"`
		RecoveryQuestion struct {
			Question string `json:"question"`
		} `json:"recovery_question"`
	} `json:"credentials"`
	ID              string  `json:"id"`
	LastLogin       string  `json:"lastLogin"`
	LastUpdated     string  `json:"lastUpdated"`
	PasswordChanged string  `json:"passwordChanged"`
	Profile         profile `json:"profile"`
	Status          string  `json:"status"`
	StatusChanged   string  `json:"statusChanged"`
	Links           struct {
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
	} `json:"_links"`

	factors []UserMFAFactor
}

type UserMFAFactor struct {
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

type profile struct {
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

func (u User) String() string {
	return Stringify(u)
}

func (s *UsersService) GetByID(id string) (*User, *Response, error) {
	u := fmt.Sprintf("user/%v", id)
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
