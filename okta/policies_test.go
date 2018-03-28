package okta

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

var testpasswordpolicy *PasswordPolicy
var testsingonpolicy *SignOnPolicy
var testmfapolicy *MfaPolicy
var testpolicy *Policy
var testpasswordrule *PasswordRule
var testsingonrule *SignOnRule
var testmfarule *MfaRule
var testrule *Rule

func setupTestPolicies() {

	hmm, _ := time.Parse("2006-01-02T15:04:05.000Z", "2018-02-16T19:59:05.000Z")

	testpolicy = &Policy{
		ID:          "00pedv3qclXeC2aFH0h7",
		Type:        "PASSWORD",
		Name:        "PasswordPolicy",
		System:      false,
		Description: "Unit Test Password Policy",
		Priority:    2,
		Status:      "ACTIVE",
		Created:     hmm,
		LastUpdated: hmm,
	}
	testpolicy.Conditions.People.Groups.Include = []string{"00ge0t33mvT5q62O40h7"}
	testpolicy.Conditions.AuthProvider.Provider = "OKTA"
	testpolicy.Settings.Recovery.Factors.OktaEmail.Status = "ACTIVE"
	testpolicy.Settings.Recovery.Factors.RecoveryQuestion.Status = "ACTIVE"
	testpolicy.Settings.Password.Complexity.MinLength = 12
	testpolicy.Settings.Password.Age.HistoryCount = 5
	testpolicy.Links.Self.Href = "https://your-domain.okta.com/api/v1/policies/00pedv3qclXeC2aFH0h7"
	testpolicy.Links.Self.Hints.Allow = []string{"GET PUT DELETE"}
	testpolicy.Links.Deactivate.Href = "https://your-domain.okta.com/api/v1/policies/00pedv3qclXeC2aFH0h7/lifecycle/deactivate"
	testpolicy.Links.Deactivate.Hints.Allow = []string{"POST"}
	testpolicy.Links.Rules.Href = "https://your-domain.okta.com/api/v1/policies/00pedv3qclXeC2aFH0h7/rules"
	testpolicy.Links.Rules.Hints.Allow = []string{"GET POST"}

	testpasswordpolicy = &PasswordPolicy{
		Type:        "PASSWORD",
		Name:        "PasswordPolicy",
		System:      false,
		Description: "Unit Test Password Policy",
		Priority:    2,
		Status:      "ACTIVE",
		Created:     hmm,
		LastUpdated: hmm,
	}
	testpolicy.Conditions.People.Groups.Include = []string{"00ge0t33mvT5q62O40h7"}
	testpolicy.Conditions.AuthProvider.Provider = "OKTA"
	testpolicy.Settings.Recovery.Factors.OktaEmail.Status = "ACTIVE"
	testpolicy.Settings.Recovery.Factors.RecoveryQuestion.Status = "ACTIVE"
	testpolicy.Settings.Password.Complexity.MinLength = 12
	testpolicy.Settings.Password.Age.HistoryCount = 5

}

func TestPolicyGet(t *testing.T) {

	setup()
	defer teardown()
	setupTestPolicies()

	temp, err := json.Marshal(testpolicy)
	if err != nil {
		t.Errorf("Policy.Get json Marshall returned error: %v", err)
	}
	policyTestJSONString := string(temp)

	mux.HandleFunc("/policies/00pedv3qclXeC2aFH0h7", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testAuthHeader(t, r)
		fmt.Fprint(w, policyTestJSONString)
	})

	policy, _, err := client.Policies.GetPolicy("00pedv3qclXeC2aFH0h7")
	if err != nil {
		t.Errorf("Policy.Get returned error: %v", err)
	}
	if !reflect.DeepEqual(policy, testpolicy) {
		t.Errorf("client.Policies.GetPolicy returned \n\t%+v, want \n\t%+v\n", policy, testpolicy)
	}
}

func testPolicyCreate(t *testing.T, inputpolicy interface{}, policy *Policy) {

	setup()
	defer teardown()
	setupTestPolicies()

	temp, err := json.Marshal(policy)
	if err != nil {
		t.Errorf("Policy.Create json Marshall returned error: %v", err)
	}
	policyTestJSONString := string(temp)

	mux.HandleFunc("/policies", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testAuthHeader(t, r)
		fmt.Fprint(w, policyTestJSONString)
	})

	outputpolicy, _, err := client.Policies.CreatePolicy(inputpolicy)
	if err != nil {
		t.Errorf("Policy.Create returned error: %v", err)
	}
	if !reflect.DeepEqual(policy, testpolicy) {
		t.Errorf("client.Policies.Create returned \n\t%+v, want \n\t%+v\n", outputpolicy, policy)
	}
}

func TestPolicyCreate(t *testing.T) {
	t.Run("Password", func(t *testing.T) {
		testPolicyCreate(t, testpasswordpolicy, testpolicy)
	})
}

//func TestUserUpdate(t *testing.T) {
//
//	setup()
//	defer teardown()
//	setupTestUsers()
//
//	// our updateuser struct with profile changes to pass into the Update function
//	updateuser := &NewUser{
//		Profile: userProfile{Login: "isaac.brock@example.com",
//			FirstName: "Herschel",
//			LastName:  "Brock",
//			Email:     "isaac.brock@example.com",
//		},
//	}
//
//	temp, err := json.Marshal(updateuser)
//	if err != nil {
//		t.Errorf("Users.Update json Marshall returned error: %v", err)
//	}
//	updateTestJSONString := string(temp)
//
//	mux.HandleFunc("/users/00ub0oNGTSWTBKOLGLNR", func(w http.ResponseWriter, r *http.Request) {
//		testMethod(t, r, "POST")
//		testAuthHeader(t, r)
//		fmt.Fprint(w, updateTestJSONString)
//	})
//
//	user, _, err := client.Users.Update(*updateuser, "00ub0oNGTSWTBKOLGLNR")
//	if err != nil {
//		t.Errorf("Users.Update returned error: %v", err)
//	}
//	if !reflect.DeepEqual(user.Profile.FirstName, updateuser.Profile.FirstName) {
//		t.Errorf("client.Users.Update returned \n\t%+v, want \n\t%+v\n", user.Profile.FirstName, updateuser.Profile.FirstName)
//	}
//}

//func TestUserDelete(t *testing.T) {
//
//	setup()
//	defer teardown()
//	setupTestUsers()
//
//	// user delete only works when user status is DEPROVISIONED
//	testuser.Status = "DEPROVISIONED"
//
//	mux.HandleFunc("/users/00ub0oNGTSWTBKOLGLNR", func(w http.ResponseWriter, r *http.Request) {
//		testMethod(t, r, "DELETE")
//		testAuthHeader(t, r)
//		fmt.Fprint(w, "")
//	})
//
//	_, err := client.Users.Delete("00ub0oNGTSWTBKOLGLNR")
//	if err != nil {
//		t.Errorf("Users.Delete returned error: %v", err)
//	}
//}

//func TestListRoles(t *testing.T) {
//
//	setup()
//	defer teardown()
//	setupTestRoles()
//
//	temp, err := json.Marshal(testroles.Role)
//	if err != nil {
//		t.Errorf("Users.ListRoles json Marshall returned error: %v", err)
//	}
//	roleTestJSONString := string(temp)
//
//	mux.HandleFunc("/users/00ub0oNGTSWTBKOLGLNR/roles", func(w http.ResponseWriter, r *http.Request) {
//		testMethod(t, r, "GET")
//		testAuthHeader(t, r)
//		fmt.Fprint(w, roleTestJSONString)
//	})
//
//	roles, _, err := client.Users.ListRoles("00ub0oNGTSWTBKOLGLNR")
//	if err != nil {
//		t.Errorf("Users.ListRoles returned error: %v", err)
//	}
//	if !reflect.DeepEqual(roles, testroles) {
//		t.Errorf("client.Users.ListRoles returned \n\t%+v, want \n\t%+v\n", roles, testroles)
//	}
//}
