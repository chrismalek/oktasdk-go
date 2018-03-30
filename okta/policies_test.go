package okta

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

var testInputPassPolicy *PasswordPolicy
var testInputSignonPolicy *SignOnPolicy
var testPassPolicy *Policy
var testSignonPolicy *Policy

var testPassPolicies *policies
var testSignonPolicies *policies

var testInputPassRule *PasswordRule
var testInputSignonRule *SignOnRule
var testPassRule *Rule
var testSignonRule *Rule

func setupTestPolicies() {

	hmm, _ := time.Parse("2006-01-02T15:04:05.000Z", "2018-02-16T19:59:05.000Z")

	// password policy
	testPassPolicy = &Policy{
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
	testPassPolicy.Conditions.AuthProvider.Provider = "OKTA"
	testPassPolicy.Settings.Recovery.Factors.OktaEmail.Status = "ACTIVE"
	testPassPolicy.Settings.Recovery.Factors.RecoveryQuestion.Status = "ACTIVE"
	testPassPolicy.Settings.Password.Complexity.MinLength = 12
	testPassPolicy.Settings.Password.Age.HistoryCount = 5
	testPassPolicy.Links.Self.Href = "https://your-domain.okta.com/api/v1/policies/00pedv3qclXeC2aFH0h7"
	testPassPolicy.Links.Self.Hints.Allow = []string{"GET PUT DELETE"}
	testPassPolicy.Links.Deactivate.Href = "https://your-domain.okta.com/api/v1/policies/00pedv3qclXeC2aFH0h7/lifecycle/deactivate"
	testPassPolicy.Links.Deactivate.Hints.Allow = []string{"POST"}
	testPassPolicy.Links.Rules.Href = "https://your-domain.okta.com/api/v1/policies/00pedv3qclXeC2aFH0h7/rules"
	testPassPolicy.Links.Rules.Hints.Allow = []string{"GET POST"}

	// signon policy
	testSignonPolicy = &Policy{
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
	testSignonPolicy.Links.Self.Href = "https://your-domain.okta.com/api/v1/policies/00pedv3qclXeC2aFH0h7"
	testSignonPolicy.Links.Self.Hints.Allow = []string{"GET PUT DELETE"}
	testSignonPolicy.Links.Deactivate.Href = "https://your-domain.okta.com/api/v1/policies/00pedv3qclXeC2aFH0h7/lifecycle/deactivate"
	testSignonPolicy.Links.Deactivate.Hints.Allow = []string{"POST"}
	testSignonPolicy.Links.Rules.Href = "https://your-domain.okta.com/api/v1/policies/00pedv3qclXeC2aFH0h7/rules"
	testSignonPolicy.Links.Rules.Hints.Allow = []string{"GET POST"}

	// input password policy
	testInputPassPolicy = &PasswordPolicy{
		Type:        "PASSWORD",
		Name:        "PasswordPolicy",
		System:      false,
		Description: "Unit Test Password Policy",
		Priority:    2,
		Status:      "ACTIVE",
		Created:     hmm,
		LastUpdated: hmm,
	}
	testInputPassPolicy.Conditions.AuthProvider.Provider = "OKTA"
	testInputPassPolicy.Settings.Recovery.Factors.OktaEmail.Status = "ACTIVE"
	testInputPassPolicy.Settings.Recovery.Factors.RecoveryQuestion.Status = "ACTIVE"
	testInputPassPolicy.Settings.Password.Complexity.MinLength = 12
	testInputPassPolicy.Settings.Password.Age.HistoryCount = 5

	// input signon policy
	testInputSignonPolicy = &SignOnPolicy{
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
	testPassPolicies = new(policies)
	testPassPolicies.Policies = append(testPassPolicies.Policies, *testPassPolicy)

	// slice of signon policies
	testSignonPolicies = new(policies)
	testSignonPolicies.Policies = append(testSignonPolicies.Policies, *testSignonPolicy)

}

func TestPolicyGet(t *testing.T) {

	setup()
	defer teardown()
	setupTestPolicies()

	temp, err := json.Marshal(testPassPolicy)
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
	if !reflect.DeepEqual(outputpolicy, testPassPolicy) {
		t.Errorf("client.Policies.GetPolicy returned \n\t%+v, want \n\t%+v\n", outputpolicy, testPassPolicy)
	}
}

func testGetPoliciesByType(t *testing.T, policytype string, policies *policies) {

	setup()
	defer teardown()
	setupTestPolicies()

	temp, err := json.Marshal(policies.Policies)
	if err != nil {
		t.Errorf("Policies.GetPoliciesByType json Marshall returned error: %v", err)
	}
	policyTestJSONString := string(temp)

	// apparently query params don't constitute a route w mux
	//mux.HandleFunc("/policies?type={policytype}", func(w http.ResponseWriter, r *http.Request) {
	mux.HandleFunc("/policies", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testAuthHeader(t, r)
		fmt.Fprint(w, policyTestJSONString)
	})

	outputpolicies, _, err := client.Policies.GetPoliciesByType(policytype)
	if err != nil {
		t.Errorf("Policies.GetPoliciesByType returned error: %v", err)
	}
	if !reflect.DeepEqual(outputpolicies, policies) {
		t.Errorf("client.Policies.GetPoliciesByType returned \n\t%+v, want \n\t%+v\n", outputpolicies, policies.Policies)
	}
}

func TestGetPoliciesByType(t *testing.T) {
	t.Run("Password", func(t *testing.T) {
		testGetPoliciesByType(t, "PASSWORD", testPassPolicies)
	})
	t.Run("SignOn", func(t *testing.T) {
		testGetPoliciesByType(t, "OKTA_SIGN_ON", testSignonPolicies)
	})
}

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
		testPolicyCreate(t, testInputPassPolicy, testPassPolicy)
	})
	t.Run("SignOn", func(t *testing.T) {
		testPolicyCreate(t, testInputSignonPolicy, testSignonPolicy)
	})
}

func testPolicyUpdate(t *testing.T, updatepolicy interface{}, policy *Policy) (*Policy, *Policy) {

	setup()
	defer teardown()

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

	setupTestPolicies()

	testTempPassPolicy := testPassPolicy
	testTempPassPolicy.Name = "PasswordPolicyUpdated"
	testTempPassPolicy.Description = "Unit Test Password Policy Updated"
	testTempPassPolicy.Settings.Password.Complexity.MinLength = 14
	testTempPassPolicy.Settings.Password.Age.HistoryCount = 8

	testUpdatePassPolicy := &PasswordPolicy{
		Type:        "PASSWORD",
		Name:        "PasswordPolicyUpdated",
		Description: "Unit Test Password Policy Updated",
		Status:      "ACTIVE",
	}
	testUpdatePassPolicy.Settings.Password.Complexity.MinLength = 14
	testUpdatePassPolicy.Settings.Password.Age.HistoryCount = 8

	t.Run("Password", func(t *testing.T) {
		output, policy := testPolicyUpdate(t, testUpdatePassPolicy, testTempPassPolicy)
		before := output.Settings.Password.Complexity.MinLength
		after := policy.Settings.Password.Complexity.MinLength
		if !reflect.DeepEqual(before, after) {
			t.Errorf("client.Policies.UpdatePolicy returned \n\t%+v, want \n\t%+v\n", before, after)
		}
	})

	setupTestPolicies()

	testTempSignonPolicy := testSignonPolicy
	testTempSignonPolicy.Name = "SignOnPolicyUpdated"
	testTempSignonPolicy.Description = "Unit Test SignOn Policy Updated"

	testUpdateSignonPolicy := &PasswordPolicy{
		Type:        "OKTA_SIGN_ON",
		Name:        "SignOnPolicyUpdated",
		Description: "Unit Test SignOn Policy Updated",
		Status:      "ACTIVE",
	}

	t.Run("SignOn", func(t *testing.T) {
		output, policy := testPolicyUpdate(t, testUpdateSignonPolicy, testTempSignonPolicy)
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
		err := testInputPassPolicy.PeopleCondition("groups", "include", []string{"00ge0t33mvT5q62O40h7"})
		if err != nil {
			t.Errorf("client.PasswordPolicyPeopleCondition returned error: %v", err)
		}
		if testInputPassPolicy.Conditions.People.Groups.Include == nil {
			t.Errorf("client.PasswordPolicy.PeopleCondition returned a nil value")
		}
	})

	setupTestPolicies()
	t.Run("SignOn", func(t *testing.T) {
		err := testInputSignonPolicy.PeopleCondition("groups", "include", []string{"00ge0t33mvT5q62O40h7"})
		if err != nil {
			t.Errorf("client.SignOnPolicy.PeopleCondition returned error: %v", err)
		}
		if testInputSignonPolicy.Conditions.People.Groups.Include == nil {
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
