package okta

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

var testinputpasspolicy *PasswordPolicy
var testinputsignonpolicy *SignOnPolicy
var testpasspolicy *Policy
var testsignonpolicy *Policy

var testpasspolicies *policies
var testsignonpolicies *policies

var testinputpassrule *PasswordRule
var testinputsignonrule *SignOnRule
var testpassrule *Rule
var testsignonrule *Rule

func setupTestPolicies() {

	hmm, _ := time.Parse("2006-01-02T15:04:05.000Z", "2018-02-16T19:59:05.000Z")

	// password policy
	testpasspolicy = &Policy{
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
	testpasspolicy.Conditions.AuthProvider.Provider = "OKTA"
	testpasspolicy.Settings.Recovery.Factors.OktaEmail.Status = "ACTIVE"
	testpasspolicy.Settings.Recovery.Factors.RecoveryQuestion.Status = "ACTIVE"
	testpasspolicy.Settings.Password.Complexity.MinLength = 12
	testpasspolicy.Settings.Password.Age.HistoryCount = 5
	testpasspolicy.Links.Self.Href = "https://your-domain.okta.com/api/v1/policies/00pedv3qclXeC2aFH0h7"
	testpasspolicy.Links.Self.Hints.Allow = []string{"GET PUT DELETE"}
	testpasspolicy.Links.Deactivate.Href = "https://your-domain.okta.com/api/v1/policies/00pedv3qclXeC2aFH0h7/lifecycle/deactivate"
	testpasspolicy.Links.Deactivate.Hints.Allow = []string{"POST"}
	testpasspolicy.Links.Rules.Href = "https://your-domain.okta.com/api/v1/policies/00pedv3qclXeC2aFH0h7/rules"
	testpasspolicy.Links.Rules.Hints.Allow = []string{"GET POST"}

	// signon policy
	testsignonpolicy = &Policy{
		ID:          "00pedv3qclXeC2aFH0h7",
		Type:        "OKTA_SIGN_ON",
		Name:        "SignOnPolicy",
		System:      false,
		Description: "Unit Test SignOn Policy",
		Priority:    2,
		Status:      "ACTIVE",
		Created:     hmm,
		LastUpdated: hmm,
	}
	testsignonpolicy.Links.Self.Href = "https://your-domain.okta.com/api/v1/policies/00pedv3qclXeC2aFH0h7"
	testsignonpolicy.Links.Self.Hints.Allow = []string{"GET PUT DELETE"}
	testsignonpolicy.Links.Deactivate.Href = "https://your-domain.okta.com/api/v1/policies/00pedv3qclXeC2aFH0h7/lifecycle/deactivate"
	testsignonpolicy.Links.Deactivate.Hints.Allow = []string{"POST"}
	testsignonpolicy.Links.Rules.Href = "https://your-domain.okta.com/api/v1/policies/00pedv3qclXeC2aFH0h7/rules"
	testsignonpolicy.Links.Rules.Hints.Allow = []string{"GET POST"}

	// input password policy
	testinputpasspolicy = &PasswordPolicy{
		Type:        "PASSWORD",
		Name:        "PasswordPolicy",
		System:      false,
		Description: "Unit Test Password Policy",
		Priority:    2,
		Status:      "ACTIVE",
		Created:     hmm,
		LastUpdated: hmm,
	}
	testinputpasspolicy.Conditions.AuthProvider.Provider = "OKTA"
	testinputpasspolicy.Settings.Recovery.Factors.OktaEmail.Status = "ACTIVE"
	testinputpasspolicy.Settings.Recovery.Factors.RecoveryQuestion.Status = "ACTIVE"
	testinputpasspolicy.Settings.Password.Complexity.MinLength = 12
	testinputpasspolicy.Settings.Password.Age.HistoryCount = 5

	// input signon policy
	testinputsignonpolicy = &SignOnPolicy{
		Type:        "OKTA_SIGN_ON",
		Name:        "SignOnPolicy",
		System:      false,
		Description: "Unit Test SignOn Policy",
		Priority:    2,
		Status:      "ACTIVE",
		Created:     hmm,
		LastUpdated: hmm,
	}

	// slice of password policies
	testpasspolicies = new(policies)
	testpasspolicies.Policies = append(testpasspolicies.Policies, *testpasspolicy)

	// slice of signon policies
	testsignonpolicies = new(policies)
	testsignonpolicies.Policies = append(testsignonpolicies.Policies, *testsignonpolicy)

}

func TestPolicyGet(t *testing.T) {

	setup()
	defer teardown()
	setupTestPolicies()

	temp, err := json.Marshal(testpasspolicy)
	if err != nil {
		t.Errorf("Polices.GetPolicy json Marshall returned error: %v", err)
	}
	policyTestJSONString := string(temp)

	mux.HandleFunc("/policies/00pedv3qclXeC2aFH0h7", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testAuthHeader(t, r)
		fmt.Fprint(w, policyTestJSONString)
	})

	outputpolicy, _, err := client.Policies.GetPolicy("00pedv3qclXeC2aFH0h7")
	if err != nil {
		t.Errorf("Policies.GetPolicy returned error: %v", err)
	}
	if !reflect.DeepEqual(outputpolicy, testpasspolicy) {
		t.Errorf("client.Policies.GetPolicy returned \n\t%+v, want \n\t%+v\n", outputpolicy, testpasspolicy)
	}
}

//func testGetPoliciesByType(t *testing.T, policytype string, policies *policies) {
//
//	setup()
//	defer teardown()
//	setupTestPolicies()
//
//	temp, err := json.Marshal(policies.Policies)
//	if err != nil {
//		t.Errorf("Policies.GetPoliciesByType json Marshall returned error: %v", err)
//	}
//	policyTestJSONString := string(temp)
//
//	uri := fmt.Sprintf("/policies?type=%v", policytype)
//	mux.HandleFunc(uri, func(w http.ResponseWriter, r *http.Request) {
//		testMethod(t, r, "GET")
//		testAuthHeader(t, r)
//		fmt.Fprint(w, policyTestJSONString)
//	})
//
//	t.Errorf("%+v", policies.Policies)
//	t.Errorf("%+v", policyTestJSONString)
//
//	outputpolicies, _, err := client.Policies.GetPolicesByType(policytype)
//	if err != nil {
//		t.Errorf("Policies.GetPoliciesByType returned error: %v", err)
//	}
//	if !reflect.DeepEqual(outputpolicies, policies.Policies) {
//		t.Errorf("client.Policies.GetPoliciesByType returned \n\t%+v, want \n\t%+v\n", outputpolicies, policies.Policies)
//	}
//}

//func TestGetPoliciesByType(t *testing.T) {
//	t.Run("Password", func(t *testing.T) {
//		testGetPoliciesByType(t, "PASSWORD", testpasspolicies)
//	})
//	t.Run("SignOn", func(t *testing.T) {
//		testGetPoliciesByType(t, "OKTA_SIGN_ON", testsignonpolicies)
//	})
//}

func testPolicyCreate(t *testing.T, inputpolicy interface{}, policy *Policy) {

	setup()
	defer teardown()
	setupTestPolicies()

	temp, err := json.Marshal(policy)
	if err != nil {
		t.Errorf("Policies.CreatePolicy json Marshall returned error: %v", err)
	}
	policyTestJSONString := string(temp)

	mux.HandleFunc("/policies", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testAuthHeader(t, r)
		fmt.Fprint(w, policyTestJSONString)
	})

	outputpolicy, _, err := client.Policies.CreatePolicy(inputpolicy)
	if err != nil {
		t.Errorf("Policies.CreatePolicy returned error: %v", err)
	}
	if !reflect.DeepEqual(outputpolicy, policy) {
		t.Errorf("client.Policies.CreatePolicy returned \n\t%+v, want \n\t%+v\n", outputpolicy, policy)
	}
}

func TestPolicyCreate(t *testing.T) {
	t.Run("Password", func(t *testing.T) {
		testPolicyCreate(t, testinputpasspolicy, testpasspolicy)
	})
	t.Run("SignOn", func(t *testing.T) {
		testPolicyCreate(t, testinputsignonpolicy, testsignonpolicy)
	})
}

func testPolicyUpdate(t *testing.T, updatepolicy interface{}, policy *Policy) (*Policy, *Policy) {

	setup()
	defer teardown()
	setupTestPolicies()

	temp, err := json.Marshal(updatepolicy)
	if err != nil {
		t.Errorf("Policies.UpdatePolicy json Marshall returned error: %v", err)
	}
	updateTestJSONString := string(temp)

	mux.HandleFunc("/policies/00pedv3qclXeC2aFH0h7", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		testAuthHeader(t, r)
		fmt.Fprint(w, updateTestJSONString)
	})

	outputpolicy, _, err := client.Policies.UpdatePolicy("00pedv3qclXeC2aFH0h7", updatepolicy)
	if err != nil {
		t.Errorf("Policies.UpdatePolicy returned error: %v", err)
	}
	return outputpolicy, policy
}

func TestPolicyUpdate(t *testing.T) {

	testtemppasspolicy := testpasspolicy
	testtemppasspolicy.Name = "PasswordPolicyUpdated"
	testtemppasspolicy.Description = "Unit Test Password Policy Updated"
	testtemppasspolicy.Settings.Password.Complexity.MinLength = 14
	testtemppasspolicy.Settings.Password.Age.HistoryCount = 8

	testupdatepasspolicy := &PasswordPolicy{
		Type:        "PASSWORD",
		Name:        "PasswordPolicyUpdated",
		Description: "Unit Test Password Policy Updated",
		Status:      "ACTIVE",
	}
	testupdatepasspolicy.Settings.Password.Complexity.MinLength = 14
	testupdatepasspolicy.Settings.Password.Age.HistoryCount = 8

	t.Run("Password", func(t *testing.T) {
		output, policy := testPolicyUpdate(t, testupdatepasspolicy, testtemppasspolicy)
		before := output.Settings.Password.Complexity.MinLength
		after := policy.Settings.Password.Complexity.MinLength
		if !reflect.DeepEqual(before, after) {
			t.Errorf("client.Policies.UpdatePolicy returned \n\t%+v, want \n\t%+v\n", before, after)
		}
	})

	testtempsignonpolicy := testsignonpolicy
	testtempsignonpolicy.Name = "SignOnPolicyUpdated"
	testtempsignonpolicy.Description = "Unit Test SignOn Policy Updated"

	testupdatesignonpolicy := &PasswordPolicy{
		Type:        "OKTA_SIGN_ON",
		Name:        "SignOnPolicyUpdated",
		Description: "Unit Test SignOn Policy Updated",
		Status:      "ACTIVE",
	}

	t.Run("SignOn", func(t *testing.T) {
		output, policy := testPolicyUpdate(t, testupdatesignonpolicy, testtempsignonpolicy)
		before := output.Name
		after := policy.Name
		if !reflect.DeepEqual(before, after) {
			t.Errorf("client.Policies.UpdatePolicy returned \n\t%+v, want \n\t%+v\n", before, after)
		}
	})
}

func TestPolicyUpdatePeopleCondition(t *testing.T) {

	setup()
	defer teardown()
	setupTestPolicies()

	t.Run("Password", func(t *testing.T) {
		err := testinputpasspolicy.PeopleCondition("groups", "include", []string{"00ge0t33mvT5q62O40h7"})
		if err != nil {
			t.Errorf("client.PasswordPolicyPeopleCondition returned error: %v", err)
		}
		if testinputpasspolicy.Conditions.People.Groups.Include == nil {
			t.Errorf("client.PasswordPolicy.PeopleCondition returned a nil value")
		}
	})
	t.Run("SignOn", func(t *testing.T) {
		err := testinputsignonpolicy.PeopleCondition("groups", "include", []string{"00ge0t33mvT5q62O40h7"})
		if err != nil {
			t.Errorf("client.SignOnPolicy.PeopleCondition returned error: %v", err)
		}
		if testinputpasspolicy.Conditions.People.Groups.Include == nil {
			t.Errorf("client.SignOnPolicy.PeopleCondition returned a nil value")
		}
	})
}

func TestPolicyDelete(t *testing.T) {

	setup()
	defer teardown()
	setupTestPolicies()

	mux.HandleFunc("/policies/00pedv3qclXeC2aFH0h7", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		testAuthHeader(t, r)
		fmt.Fprint(w, "")
	})

	_, err := client.Policies.DeletePolicy("00pedv3qclXeC2aFH0h7")
	if err != nil {
		t.Errorf("Policies.DeletePolicy returned error: %v", err)
	}
}

func TestActivatePolicy(t *testing.T) {

	setup()
	defer teardown()
	setupTestPolicies()

	mux.HandleFunc("/policies/00pedv3qclXeC2aFH0h7/lifecycle/activate", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testAuthHeader(t, r)
		fmt.Fprint(w, "")
	})

	_, err := client.Policies.ActivatePolicy("00pedv3qclXeC2aFH0h7")
	if err != nil {
		t.Errorf("Policies.ActivatePolicy returned error: %v", err)
	}
}

func TestDeactivatePolicy(t *testing.T) {

	setup()
	defer teardown()
	setupTestPolicies()

	mux.HandleFunc("/policies/00pedv3qclXeC2aFH0h7/lifecycle/deactivate", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testAuthHeader(t, r)
		fmt.Fprint(w, "")
	})

	_, err := client.Policies.DeactivatePolicy("00pedv3qclXeC2aFH0h7")
	if err != nil {
		t.Errorf("Policies.DeactivatePolicy returned error: %v", err)
	}
}
