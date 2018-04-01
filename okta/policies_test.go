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

var testPassRules *rules
var testSignonRules *rules

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

func setupTestRules() {

	hmm, _ := time.Parse("2006-01-02T15:04:05.000Z", "2018-02-16T19:59:05.000Z")

	// password rule
	testPassRule = &Rule{
		ID:          "0predz80vvMTwva7T0h7",
		Type:        "PASSWORD",
		Status:      "ACTIVE",
		Name:        "PasswordRule",
		Priority:    2,
		System:      false,
		Created:     hmm,
		LastUpdated: hmm,
	}
	testPassRule.Conditions.Network.Connection = "ANYWHERE"
	testPassRule.Actions.PasswordChange.Access = "ALLOW"
	testPassRule.Actions.SelfServicePasswordReset.Access = "ALLOW"
	testPassRule.Actions.SelfServiceUnlock.Access = "ALLOW"
	testPassRule.Links.Self.Href = "https://your-domain.okta.com/api/v1/policies/00pedv3qclXeC2aFH0h7"
	testPassRule.Links.Self.Hints.Allow = []string{"GET PUT DELETE"}
	testPassRule.Links.Deactivate.Href = "https://your-domain.okta.com/api/v1/policies/00pedv3qclXeC2aFH0h7/lifecycle/deactivate"
	testPassRule.Links.Deactivate.Hints.Allow = []string{"POST"}
	testPassRule.Links.Rules.Href = "https://your-domain.okta.com/api/v1/policies/00pedv3qclXeC2aFH0h7/rules"
	testPassRule.Links.Rules.Hints.Allow = []string{"GET POST"}

	// input password rule
	testInputPassRule = &PasswordRule{
		ID:          "0predz80vvMTwva7T0h7",
		Type:        "PASSWORD",
		Status:      "ACTIVE",
		Name:        "PasswordRule",
		Priority:    2,
		System:      false,
		Created:     hmm,
		LastUpdated: hmm,
	}
	testInputPassRule.Conditions.Network.Connection = "ANYWHERE"
	testInputPassRule.Actions.PasswordChange.Access = "ALLOW"
	testInputPassRule.Actions.SelfServicePasswordReset.Access = "ALLOW"
	testInputPassRule.Actions.SelfServiceUnlock.Access = "ALLOW"

	// signon rule
	testSignonRule = &Rule{
		ID:          "0predz80vvMTwva7T0h7",
		Type:        "OKTA_SIGN_ON",
		Status:      "ACTIVE",
		Name:        "SignOnRule",
		Priority:    2,
		System:      false,
		Created:     hmm,
		LastUpdated: hmm,
	}
	testSignonRule.Conditions.Network.Connection = "ANYWHERE"
	testSignonRule.Actions.SignOn.Access = "ALLOW"
	testSignonRule.Actions.SignOn.Session.MaxSessionIdleMinutes = 120
	testSignonRule.Links.Self.Href = "https://your-domain.okta.com/api/v1/policies/00pedv3qclXeC2aFH0h7"
	testSignonRule.Links.Self.Hints.Allow = []string{"GET PUT DELETE"}
	testSignonRule.Links.Deactivate.Href = "https://your-domain.okta.com/api/v1/policies/00pedv3qclXeC2aFH0h7/lifecycle/deactivate"
	testSignonRule.Links.Deactivate.Hints.Allow = []string{"POST"}
	testSignonRule.Links.Rules.Href = "https://your-domain.okta.com/api/v1/policies/00pedv3qclXeC2aFH0h7/rules"
	testSignonRule.Links.Rules.Hints.Allow = []string{"GET POST"}

	// input signon rule
	testInputSignonRule = &SignOnRule{
		ID:          "0predz80vvMTwva7T0h7",
		Type:        "OKTA_SIGN_ON",
		Status:      "ACTIVE",
		Name:        "SignOnRule",
		Priority:    2,
		System:      false,
		Created:     hmm,
		LastUpdated: hmm,
	}
	testInputSignonRule.Conditions.Network.Connection = "ANYWHERE"
	testInputSignonRule.Actions.SignOn.Access = "ALLOW"
	testInputSignonRule.Actions.SignOn.Session.MaxSessionIdleMinutes = 120

	// slice of password rules
	testPassRules = new(rules)
	testPassRules.Rules = append(testPassRules.Rules, *testPassRule)

	// slice of signon rules
	testSignonRules = new(rules)
	testSignonRules.Rules = append(testSignonRules.Rules, *testSignonRule)

}

func TestGetPolicy(t *testing.T) {

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

func TestiGetRule(t *testing.T) {

	setup()
	defer teardown()
	setupTestRules()

	temp, err := json.Marshal(testPassRule)
	if err != nil {
		t.Errorf("Polices.GetPolicyRule json Marshall returned error: %v", err)
	}
	ruleTestJSONString := string(temp)

	mux.HandleFunc("/policies/00pedv3qclXeC2aFH0h7rules/0predz80vvMTwva7T0h7", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testAuthHeader(t, r)
		fmt.Fprint(w, ruleTestJSONString)
	})

	outputrule, _, err := client.Policies.GetPolicyRule("00pedv3qclXeC2aFH0h7", "0predz80vvMTwva7T0h7")
	if err != nil {
		t.Errorf("Policies.GetPolicyRule returned error: %v", err)
	}
	if !reflect.DeepEqual(outputrule, testPassRule) {
		t.Errorf("client.Policies.GetPolicy returned \n\t%+v, want \n\t%+v\n", outputrule, testPassRule)
	}
}

func testGetPoliciesByType(t *testing.T, policytype string, policies *policies) {

	setup()
	defer teardown()

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
		t.Errorf("client.Policies.GetPoliciesByType returned \n\t%+v, want \n\t%+v\n", outputpolicies, policies)
	}
}

func TestGetPoliciesByType(t *testing.T) {
	setupTestPolicies()

	t.Run("Password", func(t *testing.T) {
		testGetPoliciesByType(t, "PASSWORD", testPassPolicies)
	})
	t.Run("SignOn", func(t *testing.T) {
		testGetPoliciesByType(t, "OKTA_SIGN_ON", testSignonPolicies)
	})
}

func testGetPolicyRules(t *testing.T, rules *rules) {

	setup()
	defer teardown()

	temp, err := json.Marshal(rules.Rules)
	if err != nil {
		t.Errorf("Policies.GetPolicyRules json Marshall returned error: %v", err)
	}
	ruleTestJSONString := string(temp)

	mux.HandleFunc("/policies/00pedv3qclXeC2aFH0h7/rules", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testAuthHeader(t, r)
		fmt.Fprint(w, ruleTestJSONString)
	})

	outputrules, _, err := client.Policies.GetPolicyRules("00pedv3qclXeC2aFH0h7")
	if err != nil {
		t.Errorf("Policies.GetPolicyRules returned error: %v", err)
	}
	if !reflect.DeepEqual(outputrules, rules) {
		t.Errorf("client.Policies.GetPolicyRules returned \n\t%+v, want \n\t%+v\n", outputrules, rules)
	}
}

func TestGetPolicyRules(t *testing.T) {
	setupTestRules()

	t.Run("Password", func(t *testing.T) {
		testGetPolicyRules(t, testPassRules)
	})
	t.Run("SignOn", func(t *testing.T) {
		testGetPolicyRules(t, testSignonRules)
	})
}

func testPolicyCreate(t *testing.T, inputpolicy interface{}, policy *Policy) {

	setup()
	defer teardown()

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
	setupTestPolicies()

	t.Run("Password", func(t *testing.T) {
		testPolicyCreate(t, testInputPassPolicy, testPassPolicy)
	})
	t.Run("SignOn", func(t *testing.T) {
		testPolicyCreate(t, testInputSignonPolicy, testSignonPolicy)
	})
}

func testRuleCreate(t *testing.T, inputrule interface{}, rule *Rule) {

	setup()
	defer teardown()

	temp, err := json.Marshal(rule)
	if err != nil {
		t.Errorf("Policies.CreateiPolicyRule json Marshall returned error: %v", err)
	}
	ruleTestJSONString := string(temp)

	mux.HandleFunc("/policies/00pedv3qclXeC2aFH0h7/rules", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testAuthHeader(t, r)
		fmt.Fprint(w, ruleTestJSONString)
	})

	outputrule, _, err := client.Policies.CreatePolicyRule("00pedv3qclXeC2aFH0h7", inputrule)
	if err != nil {
		t.Errorf("Policies.CreatePolicyRule returned error: %v", err)
	}
	if !reflect.DeepEqual(outputrule, rule) {
		t.Errorf("client.Policies.CreatePolicy returned \n\t%+v, want \n\t%+v\n", outputrule, rule)
	}
}

func TestRuleCreate(t *testing.T) {
	setupTestRules()

	t.Run("Password", func(t *testing.T) {
		testRuleCreate(t, testInputPassRule, testPassRule)
	})
	t.Run("SignOn", func(t *testing.T) {
		testRuleCreate(t, testInputSignonRule, testSignonRule)
	})
}

func testPolicyUpdate(t *testing.T, updatepolicy interface{}, policy *Policy) {

	setup()
	defer teardown()

	temp, err := json.Marshal(policy)
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
	if !reflect.DeepEqual(outputpolicy.Name, policy.Name) {
		t.Errorf("client.Policies.UpdatePolicy returned \n\t%+v, want \n\t%+v\n", outputpolicy.Name, policy.Name)
	}
}

func TestPolicyUpdate(t *testing.T) {
	setupTestPolicies()

	t.Run("Password", func(t *testing.T) {
		testPolicyUpdate(t, testInputPassPolicy, testPassPolicy)
	})
	t.Run("SignOn", func(t *testing.T) {
		testPolicyUpdate(t, testInputSignonPolicy, testSignonPolicy)
	})
}

func testRuleUpdate(t *testing.T, updaterule interface{}, rule *Rule) {

	setup()
	defer teardown()

	temp, err := json.Marshal(rule)
	if err != nil {
		t.Errorf("Policies.UpdatePolicyRule json Marshall returned error: %v", err)
	}
	updateTestJSONString := string(temp)

	mux.HandleFunc("/policies/00pedv3qclXeC2aFH0h7/rules/0predz80vvMTwva7T0h7", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		testAuthHeader(t, r)
		fmt.Fprint(w, updateTestJSONString)
	})
	outputrule, _, err := client.Policies.UpdatePolicyRule("00pedv3qclXeC2aFH0h7", "0predz80vvMTwva7T0h7", updaterule)
	if err != nil {
		t.Errorf("Policies.UpdatePolicyRule returned error: %v", err)
	}
	if !reflect.DeepEqual(outputrule.Name, rule.Name) {
		t.Errorf("client.Policies.UpdatePolicyRule returned \n\t%+v, want \n\t%+v\n", outputrule.Name, rule.Name)
	}
}

func TestRuleUpdate(t *testing.T) {
	setupTestRules()

	t.Run("Password", func(t *testing.T) {
		testRuleUpdate(t, testInputPassRule, testPassRule)
	})
	t.Run("SignOn", func(t *testing.T) {
		testRuleUpdate(t, testInputSignonRule, testSignonRule)
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

func TestRuleUpdatePeopleCondition(t *testing.T) {

	setup()
	defer teardown()

	setupTestRules()
	t.Run("Password", func(t *testing.T) {
		err := testInputPassRule.PeopleCondition("groups", "include", []string{"00ge0t33mvT5q62O40h7"})
		if err != nil {
			t.Errorf("client.PasswordRulePeopleCondition returned error: %v", err)
		}
		if testInputPassRule.Conditions.People.Groups.Include == nil {
			t.Errorf("client.PasswordRule.PeopleCondition returned a nil value")
		}
	})

	setupTestRules()
	t.Run("SignOn", func(t *testing.T) {
		err := testInputSignonRule.PeopleCondition("groups", "include", []string{"00ge0t33mvT5q62O40h7"})
		if err != nil {
			t.Errorf("client.SignOnRule.PeopleCondition returned error: %v", err)
		}
		if testInputSignonRule.Conditions.People.Groups.Include == nil {
			t.Errorf("client.SignOnRule.PeopleCondition returned a nil value")
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

func TestRuleDelete(t *testing.T) {

	setup()
	defer teardown()
	setupTestRules()

	mux.HandleFunc("/policies/00pedv3qclXeC2aFH0h7/rules/0predz80vvMTwva7T0h7", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		testAuthHeader(t, r)
		fmt.Fprint(w, "")
	})

	_, err := client.Policies.DeletePolicyRule("00pedv3qclXeC2aFH0h7", "0predz80vvMTwva7T0h7")
	if err != nil {
		t.Errorf("Policies.DeletePolicyRule returned error: %v", err)
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

func TestActivateRule(t *testing.T) {

	setup()
	defer teardown()
	setupTestRules()

	mux.HandleFunc("/policies/00pedv3qclXeC2aFH0h7/rules/0predz80vvMTwva7T0h7/lifecycle/activate", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testAuthHeader(t, r)
		fmt.Fprint(w, "")
	})

	_, err := client.Policies.ActivatePolicyRule("00pedv3qclXeC2aFH0h7", "0predz80vvMTwva7T0h7")
	if err != nil {
		t.Errorf("Policies.ActivatePolicyRule returned error: %v", err)
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

func TestDeactivatePolicyRule(t *testing.T) {

	setup()
	defer teardown()
	setupTestRules()

	mux.HandleFunc("/policies/00pedv3qclXeC2aFH0h7/rules/0predz80vvMTwva7T0h7/lifecycle/deactivate", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testAuthHeader(t, r)
		fmt.Fprint(w, "")
	})

	_, err := client.Policies.DeactivatePolicyRule("00pedv3qclXeC2aFH0h7", "0predz80vvMTwva7T0h7")
	if err != nil {
		t.Errorf("Policies.DeactivatePolicyRule returned error: %v", err)
	}
}
