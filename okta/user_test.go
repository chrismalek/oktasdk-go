package okta

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
	// "github.com/kr/pretty"
)

var userTestJSONString = `

{
  "id": "00ub0oNGTSWTBKOLGLNR",
  "status": "ACTIVE",
  "created": "2013-06-24T16:39:18.000Z",
  "activated": "2013-06-24T16:39:19.000Z",
  "statusChanged": "2013-06-24T16:39:19.000Z",
  "lastLogin": "2013-06-24T17:39:19.000Z",
  "lastUpdated": "2013-06-27T16:35:28.000Z",
  "passwordChanged": "2013-06-24T16:39:19.000Z",
  "profile": {
    "login": "isaac.brock@example.com",
    "firstName": "Isaac",
    "lastName": "Brock",
    "nickName": "issac",
    "displayName": "Isaac Brock",
    "email": "isaac.brock@example.com",
    "secondEmail": "isaac@example.org",
    "preferredLanguage": "en-US",
    "organization": "Okta",
    "title": "Director",
    "division": "R&D",
    "department": "Engineering",
    "costCenter": "10",
    "employeeNumber": "187",
    "mobilePhone": "+1-555-415-1337",
    "primaryPhone": "+1-555-514-1337",
    "streetAddress": "301 Brannan St.",
    "city": "San Francisco",
    "state": "CA",
    "zipCode": "94107",
    "countryCode": "US"
  },
  "credentials": {
    "password": {},
    "recovery_question": {
      "question": "Who's a major player in the cowboy scene?"
    },
    "provider": {
      "type": "OKTA",
      "name": "OKTA"
    }
  },
  "_links": {
    "resetPassword": {
      "href": "https://your-domain.okta.com/api/v1/users/00ub0oNGTSWTBKOLGLNR/lifecycle/reset_password"
    },
    "resetFactors": {
      "href": "https://your-domain.okta.com/api/v1/users/00ub0oNGTSWTBKOLGLNR/lifecycle/reset_factors"
    },
    "expirePassword": {
      "href": "https://your-domain.okta.com/api/v1/users/00ub0oNGTSWTBKOLGLNR/lifecycle/expire_password"
    },
    "forgotPassword": {
      "href": "https://your-domain.okta.com/api/v1/users/00ub0oNGTSWTBKOLGLNR/credentials/forgot_password"
    },
    "changeRecoveryQuestion": {
      "href": "https://your-domain.okta.com/api/v1/users/00ub0oNGTSWTBKOLGLNR/credentials/change_recovery_question"
    },
    "deactivate": {
      "href": "https://your-domain.okta.com/api/v1/users/00ub0oNGTSWTBKOLGLNR/lifecycle/deactivate"
    },
    "changePassword": {
      "href": "https://your-domain.okta.com/api/v1/users/00ub0oNGTSWTBKOLGLNR/credentials/change_password"
    }
  }
}
`

var testuser *User

func setupTestUsers() {

	testuser = &User{
		ID:              "00ub0oNGTSWTBKOLGLNR",
		Status:          "ACTIVE",
		Created:         "2013-06-24T16:39:18.000Z",
		Activated:       "2013-06-24T16:39:19.000Z",
		StatusChanged:   "2013-06-24T16:39:19.000Z",
		LastLogin:       "2013-06-24T17:39:19.000Z",
		LastUpdated:     "2013-06-27T16:35:28.000Z",
		PasswordChanged: "2013-06-24T16:39:19.000Z",
		Profile: userProfile{Login: "isaac.brock@example.com",
			FirstName:         "Isaac",
			LastName:          "Brock",
			NickName:          "issac",
			DisplayName:       "Isaac Brock",
			Email:             "isaac.brock@example.com",
			SecondEmail:       "isaac@example.org",
			PreferredLanguage: "en-US",
			Organization:      "Okta",
			Title:             "Director",
			Division:          "R&D",
			Department:        "Engineering",
			CostCenter:        "10",
			EmployeeNumber:    "187",
			MobilePhone:       "+1-555-415-1337",
			PrimaryPhone:      "+1-555-514-1337",
			StreetAddress:     "301 Brannan St.",
			City:              "San Francisco",
			State:             "CA",
			ZipCode:           "94107",
			CountryCode:       "US"},
		Credentials: credentials{
			Password:         &passwordValue{},
			RecoveryQuestion: &recoveryQuestion{Question: "Who's a major player in the cowboy scene?"},
			Provider: &provider{Type: "OKTA",
				Name: "OKTA"},
		},
	}

	testuser.Links.ChangePassword.Href = "https://your-domain.okta.com/api/v1/users/00ub0oNGTSWTBKOLGLNR/credentials/change_password"
	testuser.Links.ResetFactors.Href = "https://your-domain.okta.com/api/v1/users/00ub0oNGTSWTBKOLGLNR/lifecycle/reset_factors"
	testuser.Links.ChangeRecoveryQuestion.Href = "https://your-domain.okta.com/api/v1/users/00ub0oNGTSWTBKOLGLNR/credentials/change_recovery_question"
	testuser.Links.Deactivate.Href = "https://your-domain.okta.com/api/v1/users/00ub0oNGTSWTBKOLGLNR/lifecycle/deactivate"
	testuser.Links.ExpirePassword.Href = "https://your-domain.okta.com/api/v1/users/00ub0oNGTSWTBKOLGLNR/lifecycle/expire_password"
	testuser.Links.ForgotPassword.Href = "https://your-domain.okta.com/api/v1/users/00ub0oNGTSWTBKOLGLNR/credentials/forgot_password"
	testuser.Links.ResetPassword.Href = "https://your-domain.okta.com/api/v1/users/00ub0oNGTSWTBKOLGLNR/lifecycle/reset_password"
}

func TestUserGet(t *testing.T) {

	setup()
	defer teardown()
	setupTestUsers()

	mux.HandleFunc("/users/00ub0oNGTSWTBKOLGLNR", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testAuthHeader(t, r)
		fmt.Fprint(w, userTestJSONString)
	})

	user, _, err := client.Users.GetByID("00ub0oNGTSWTBKOLGLNR")
	if err != nil {
		t.Errorf("Users.Get returned error: %v", err)
	}
	if !reflect.DeepEqual(user, testuser) {
		// fmt.Printf("pretty---\n%v\n---\n", pretty.Diff(user, testuser))
		t.Errorf("client.Users.GetByID returned \n\t%+v, want \n\t%+v\n", user, testuser)
	}
}

//  Test User Search Query Parameter Generation
// Test Pagination
//
