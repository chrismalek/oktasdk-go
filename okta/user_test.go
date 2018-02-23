package okta

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

var testuser *User
var testroles *userRoles

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

func setupTestRoles() {

	hmm, _ := time.Parse("2006-01-02T15:04:05.000Z", "2018-02-16T19:59:05.000Z")

	testrole := new(userRole)
	testrole.ID = "KVJUKUS7IFCE2SKO"
	testrole.Label = "User Administrator"
	testrole.Type = "USER_ADMIN"
	testrole.Status = "ACTIVE"
	testrole.Created = hmm
	testrole.LastUpdated = hmm

	testroles = new(userRoles)
	testroles.Role = append(testroles.Role, *testrole)

}

func TestUserGet(t *testing.T) {

	setup()
	defer teardown()
	setupTestUsers()

	temp, err := json.Marshal(testuser)
	if err != nil {
		t.Errorf("Users.Get json Marshall returned error: %v", err)
	}
	userTestJSONString := string(temp)

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

func TestUserUpdate(t *testing.T) {

	setup()
	defer teardown()
	setupTestUsers()

	// our updateuser struct with profile changes to pass into the Update function
	updateuser := &NewUser{
		Profile: userProfile{Login: "isaac.brock@example.com",
			FirstName: "Herschel",
			LastName:  "Brock",
			Email:     "isaac.brock@example.com",
		},
	}

	temp, err := json.Marshal(updateuser)
	if err != nil {
		t.Errorf("Users.Update json Marshall returned error: %v", err)
	}
	updateTestJSONString := string(temp)

	mux.HandleFunc("/users/00ub0oNGTSWTBKOLGLNR", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testAuthHeader(t, r)
		fmt.Fprint(w, updateTestJSONString)
	})

	user, _, err := client.Users.Update(*updateuser, "00ub0oNGTSWTBKOLGLNR")
	if err != nil {
		t.Errorf("Users.Update returned error: %v", err)
	}
	if !reflect.DeepEqual(user.Profile.FirstName, updateuser.Profile.FirstName) {
		t.Errorf("client.Users.Update returned \n\t%+v, want \n\t%+v\n", user.Profile.FirstName, updateuser.Profile.FirstName)
	}
}

func TestUserDelete(t *testing.T) {

	setup()
	defer teardown()
	setupTestUsers()

	// user delete only works when user status is DEPROVISIONED
	testuser.Status = "DEPROVISIONED"

	mux.HandleFunc("/users/00ub0oNGTSWTBKOLGLNR", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		testAuthHeader(t, r)
		fmt.Fprint(w, "")
	})

	_, err := client.Users.Delete("00ub0oNGTSWTBKOLGLNR")
	if err != nil {
		t.Errorf("Users.Delete returned error: %v", err)
	}
}

func TestListRoles(t *testing.T) {

	setup()
	defer teardown()
	setupTestRoles()

	temp, err := json.Marshal(testroles.Role)
	if err != nil {
		t.Errorf("Users.ListRoles json Marshall returned error: %v", err)
	}
	roleTestJSONString := string(temp)

	mux.HandleFunc("/users/00ub0oNGTSWTBKOLGLNR/roles", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testAuthHeader(t, r)
		fmt.Fprint(w, roleTestJSONString)
	})

	roles, _, err := client.Users.ListRoles("00ub0oNGTSWTBKOLGLNR")
	if err != nil {
		t.Errorf("Users.ListRoles returned error: %v", err)
	}
	if !reflect.DeepEqual(roles, testroles) {
		t.Errorf("client.Users.ListRoles returned \n\t%+v, want \n\t%+v\n", roles, testroles)
	}
}

func TestAssignRole(t *testing.T) {

	setup()
	defer teardown()
	setupTestRoles()

	type testRoleType struct {
		Type string `json:"type"`
	}

	mux.HandleFunc("/users/00ub0oNGTSWTBKOLGLNR/roles", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testAuthHeader(t, r)
		testBody(t, r, testRoleType{
			Type: "USER_ADMIN",
		})
		fmt.Fprint(w, "")
	})

	_, err := client.Users.AssignRole("00ub0oNGTSWTBKOLGLNR", "USER_ADMIN")
	if err != nil {
		t.Errorf("Users.AssignRole returned error: %v", err)
	}
}

func TestUnAssignRole(t *testing.T) {

	setup()
	defer teardown()
	setupTestRoles()

	mux.HandleFunc("/users/00ub0oNGTSWTBKOLGLNR/roles/KVJUKUS7IFCE2SKO", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		testAuthHeader(t, r)
		fmt.Fprint(w, "")
	})

	_, err := client.Users.UnAssignRole("00ub0oNGTSWTBKOLGLNR", "KVJUKUS7IFCE2SKO")
	if err != nil {
		t.Errorf("Users.UnAssignRole returned error: %v", err)
	}
}

//  Test User Search Query Parameter Generation
// Test Pagination
//
