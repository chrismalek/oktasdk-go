package okta

import (
	"fmt"
	"time"
)

// PoliciesService handles communication with the Policy data related
// methods of the OKTA API.
type PoliciesService service

// Return the Policy object. Used to create & update policies
func (p *PoliciesService) Policy() Policy {
	return Policy{}
}

// Return the Rule object. Used to create & update rules
func (p *PoliciesService) Rule() Rule {
	return Rule{}
}

// Policy represents the Policy Object from the OKTA API
type Policy struct {
	ID          string      `json:"id,omitempty"`
	Type        string      `json:"type"`
	Name        string      `json:"name"`
	System      bool        `json:"system,omitempty"`
	Description string      `json:"description,omitempty"`
	Priority    int         `json:"priority,omitempty"`
	Status      string      `json:"status,omitempty"`
	Conditions  conditions  `json:"conditions,omitempty"`
	Settings    settings    `json:"settings,omitempty"`
	Created     time.Time   `json:"created,omitempty"`
	LastUpdated time.Time   `json:"lastUpdated,omitempty"`
	Links       policyLinks `json:"_links,omitempty"`
}

type policies struct {
	Policies []Policy `json:"-,omitempty"`
}

type conditions struct {
	People struct {
		Groups groups `json:"groups,omitempty"`
		Users  users  `json:"users,omitempty"`
	} `json:"people,omitempty"`
	AuthContext  string       `json:"authType,omitempty"`
	Network      network      `json:"network,omitempty"`
	AuthProvider authProvider `json:"authProvider,omitempty"`
}

type users struct {
	Include []string `json:"include,omitempty"`
	Exclude []string `json:"exclude,omitempty"`
}

type groups struct {
	Include []string `json:"include,omitempty"`
	Exclude []string `json:"exclude,omitempty"`
}

type network struct {
	Connection string   `json:"connection,omitempty"`
	Include    []string `json:"include,omitempty"`
	Exclude    []string `json:"exclude,omitempty"`
}

type authProvider struct {
	Provider string   `json:"provider"`
	Include  []string `json:"include,omitempty"`
}

type settings struct {
	Factors struct {
		GoogleOtp    mfaFactor `json:"google_otp,omitempty"`
		OktaOtp      mfaFactor `json:"okta_otp,omitempty"`
		OktaPush     mfaFactor `json:"okta_push,omitempty"`
		OktaQuestion mfaFactor `json:"okta_question,omitempty"`
		OktaSms      mfaFactor `json:"okta_sms,omitempty"`
		RsaToken     mfaFactor `json:"rsa_token,omitempty"`
		SymantecVip  mfaFactor `json:"symantec_vip,omitempty"`
	} `json:"factors,omitempty"`
	Password struct {
		password   password   `json:"password,omitempty"`
		recovery   recovery   `json:"recovery,omitempty"`
		delegation delegation `json:"lockout,omitempty"`
	} `json:"password,omitempty"`
}

type password struct {
	Complexity complexity `json:"complexity,omitempty"`
	Age        age        `json:"age,omitempty"`
	Lockout    lockout    `json:"lockout,omitempty"`
}

type complexity struct {
	MinLength         int        `json:"minLength,omitempty"`
	MinLowerCase      int        `json:"minLowerCase,omitempty"`
	MinUpperCase      int        `json:"minUpperCase,omitempty"`
	MinNumber         int        `json:"minNumber,omitempty"`
	MinSymbol         int        `json:"minSymbol,omitempty"`
	ExcludeUsername   bool       `json:"excludeUsername,omitempty"`
	ExcludeAttributes []string   `json:"excludeAttributes,omitempty"`
	Dictionary        dictionary `json:"dictionary,omitempty"`
}

type dictionary struct {
	Common common `json:"common,omitempty"`
}

type common struct {
	Exclude bool `json:"excllude,omitempty"`
}

type age struct {
	MaxAgeDays     int `json:"maxAgeDays,omitempty"`
	ExpireWarnDays int `json:"expireWarnDays,omitempty"`
	MinAgeMinutes  int `json:"minAgeMinutes,omitempty"`
	HistoryCount   int `json:"historyCount,omitempty"`
}

type lockout struct {
	MaxAttempts         int  `json:"maxAttempts,omitempty"`
	AutoUnlockMinutes   int  `json:"autoUnlockMinutes,omitempty"`
	ShowLockoutFailures bool `json:"showLockoutFailures,omitempty"`
}

type recovery struct {
	Factors struct {
		RecoveryQuestion policyRecoveryQuestion `json:"recovery_question,omitempty"`
		OktaEmail        oktaEmail              `json:"okta_email,omitempty"`
		OktaSms          oktaSms                `json:"okta_sms,omitempty"`
	} `json:"factors,omitempty"`
}

type policyRecoveryQuestion struct {
	Status     string                     `json:"status"`
	Properties recoveryQuestionProperties `json:"properties,omitempty"`
}

type recoveryQuestionProperties struct {
	Complexity recoveryQuestionPropertiesComplexity `json:"complexity,omitempty"`
}

type recoveryQuestionPropertiesComplexity struct {
	MinLength int `json:"minLength,omitempty"`
}

type oktaEmail struct {
	Status     string              `json:"status"`
	Properties oktaEmailProperties `json:"properties,omitempty"`
}

type oktaEmailProperties struct {
	RecoveryToken oktaEmailPropertiesRecoveryToken `json:"recoveryToken,omitempty"`
}

type oktaEmailPropertiesRecoveryToken struct {
	TokenLifetimeMinutes int `json:"tokenLifetimeMinutes,omitempty"`
}

type oktaSms struct {
	Status string `json:"status,omitempty"`
}

type delegation struct {
	Options options `json:"options,omitempty"`
}

type options struct {
	SkipUnlock bool `json:"skipUnlock,omitempty"`
}

type mfaFactor struct {
	Consent factorConsent `json:"consent,omitempty"`
	Enroll  factorEnroll  `json:"enroll,omitempty"`
}

type factorConsent struct {
	Terms factorConsentTerms `json:"terms,omitempty"`
	Type  string             `json:"type,omitempty"`
}

type factorEnroll struct {
	Self string `json:"self,omitempty"`
}

type factorConsentTerms struct {
	Format string `json:"format,omitempty"`
	Value  string `json:"value,omitempty"`
}

// TODO add MFA conditions objects
// https://developer.okta.com/docs/api/resources/policy#policy-conditions-1

type policyLinks struct {
	Self struct {
		Href  string `json:"href"`
		Hints struct {
			Allow []string `json:"allow"`
		} `json:"hints"`
	} `json:"self"`
	Activate struct {
		Href  string `json:"href"`
		Hints struct {
			Allow []string `json:"allow"`
		} `json:"hints"`
	} `json:"activate",omitempty`
	Deactivate struct {
		Href  string `json:"href"`
		Hints struct {
			Allow []string `json:"allow"`
		} `json:"hints"`
	} `json:"deactivate,omitempty"`
	Rules struct {
		Href  string `json:"href"`
		Hints struct {
			Allow []string `json:"allow"`
		} `json:"hints"`
	} `json:"rules,omitempty"`
}

// Rule represents the Rule Object from the OKTA API
type Rule struct {
	ID          string     `json:"id,omitempty"`
	Type        string     `json:"type"`
	Status      string     `json:"status,omitempty"`
	Priority    int        `json:"priority,omitempty"`
	System      string     `json:"system,omitempty"`
	Created     time.Time  `json:"created,omitempty"`
	LastUpdated time.Time  `json:"lastUpdated,omitempty"`
	Conditions  conditions `json:"conditions,omitempty"`
	Actions     actions    `json:"actions,omitempty"`
	Links       ruleLinks  `json:"_links,omitempty"`
}

type rules struct {
	Rules []Rule `json:"-,omitempty"`
}

type actions struct {
	signon struct {
		Access                  string  `json:"access"`
		RequireFactor           bool    `json:"requireFactor,omitempty"`
		FactorPromptMode        string  `json:"factorPromptMode,omitempty"`
		RememberDeviceByDefault bool    `json:"rememberDeviceByDefault,omitempty"`
		FactorLifetime          int     `json:"factorLifetime,omitempty"`
		Session                 session `json:"session,omitempty"`
	} `json:"signon,omitempty"`
	enroll struct {
		Self string `json:"self,omitempty"`
	} `json:"enroll,omitempty"`
	PasswordChange           passwordAction `json:"passwordChange,omitempty"`
	SelfServicePasswordReset passwordAction `json:"selfServicePasswordReset,omitempty"`
	SelfServiceUnlock        passwordAction `json:"selfServiceUnlock,omitempty"`
}

type passwordAction struct {
	Access string `json:"access,omitempty"`
}

type session struct {
	MaxSessionIdleMinutes     int  `json:"maxSessionIdleMinutes,omitempty"`
	MaxSessionLifetimeMinutes int  `json:"maxSessionLifetimeMinutes,omitempty"`
	UsePersistentCookie       bool `json:"usePersistentCookie,,omitempty"`
}

type ruleLinks struct {
	Self struct {
		Href  string `json:"href"`
		Hints struct {
			Allow []string `json:"allow"`
		} `json:"hints"`
	} `json:"self"`
	Activate struct {
		Href  string `json:"href"`
		Hints struct {
			Allow []string `json:"allow"`
		} `json:"hints"`
	} `json:"activate",omitempty`
	Deactivate struct {
		Href  string `json:"href"`
		Hints struct {
			Allow []string `json:"allow"`
		} `json:"hints"`
	} `json:"deactivate,omitempty"`
	Rules struct {
		Href  string `json:"href"`
		Hints struct {
			Allow []string `json:"allow"`
		} `json:"hints"`
	} `json:"rules,omitempty"`
}

// Get a policy
// Requires Policy ID from Policy object
func (p *PoliciesService) GetPolicy(id string) (*Policy, *Response, error) {
	u := fmt.Sprintf("policies/%v", id)
	req, err := p.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	policy := new(Policy)
	resp, err := p.client.Do(req, policy)
	if err != nil {
		return nil, resp, err
	}

	return policy, resp, err
}

// Get all policies by type
// Allowed types are OKTA_SIGN_ON, PASSWORD, MFA_ENROLL, or OAUTH_AUTHORIZATION_POLICY
// TODO: MFA_ENROLL obj build out not complete
func (p *PoliciesService) GetPolicesByType(policyType string) (*policies, *Response, error) {
	u := fmt.Sprintf("policies?type=%v", policyType)
	req, err := p.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	policy := make([]Policy, 0)
	resp, err := p.client.Do(req, &policy)
	if err != nil {
		return nil, resp, err
	}
	if len(policy) > 0 {
		myPolicies := new(policies)
		for _, v := range policy {
			myPolicies.Policies = append(myPolicies.Policies, v)
		}
		return myPolicies, resp, err
	}

	return nil, resp, err
}

// Delete a policy
// Requires Policy ID from Policy object
func (p *PoliciesService) DeletePolicy(id string) (*Response, error) {
	u := fmt.Sprintf("policies/%v", id)
	req, err := p.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}
	resp, err := p.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, err
}

// Create a policy
// You must pass in the Policy object created from Policies.Policy()
func (p *PoliciesService) CreatePolicy(policy Policy) (*Policy, *Response, error) {
	u := fmt.Sprintf("policies")
	req, err := p.client.NewRequest("POST", u, policy)
	if err != nil {
		return nil, nil, err
	}

	newPolicy := new(Policy)
	resp, err := p.client.Do(req, newPolicy)
	if err != nil {
		return nil, resp, err
	}

	return newPolicy, resp, err
}

// Update a policy
// You must pass in the Policy object from Policies.Policy()
// This endpoint uses a PUT so I'm going to assume partial updates are not supported
func (p *PoliciesService) UpdatePolicy(policy Policy) (*Policy, *Response, error) {
	u := fmt.Sprintf("policies")
	req, err := p.client.NewRequest("POST", u, policy)
	if err != nil {
		return nil, nil, err
	}

	updatePolicy := new(Policy)
	resp, err := p.client.Do(req, updatePolicy)
	if err != nil {
		return nil, resp, err
	}

	return updatePolicy, resp, err
}

// Activate a policy
// Requires Policy ID from Policy object
func (p *PoliciesService) ActivatePolicy(id string) (*Response, error) {
	u := fmt.Sprintf("policies/%v/lifecycle/activate", id)
	req, err := p.client.NewRequest("POST", u, nil)
	if err != nil {
		return nil, err
	}
	resp, err := p.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, err
}

// Deactivate a policy
// Requires Policy ID from Policy object
func (p *PoliciesService) DeactivatePolicy(id string) (*Response, error) {
	u := fmt.Sprintf("policies/%v/lifecycle/deactivate", id)
	req, err := p.client.NewRequest("POST", u, nil)
	if err != nil {
		return nil, err
	}
	resp, err := p.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, err
}

// Get policy rules
// Requires Policy ID from Policy object
func (p *PoliciesService) GetPolicyRules(id string) (*rules, *Response, error) {
	u := fmt.Sprintf("policies/%v/rules", id)
	req, err := p.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	rule := make([]Rule, 0)
	resp, err := p.client.Do(req, &rule)
	if err != nil {
		return nil, resp, err
	}
	if len(rule) > 0 {
		myRules := new(rules)
		for _, v := range rule {
			myRules.Rules = append(myRules.Rules, v)
		}
		return myRules, resp, err
	}

	return nil, resp, err
}

// Create a rule
// Requires Policy ID from Policy object
// You must pass in the Rule object created from Policies.Rule()
func (p *PoliciesService) CreateRule(id string, rule Rule) (*Rule, *Response, error) {
	u := fmt.Sprintf("policies/%v/rules", id)
	req, err := p.client.NewRequest("POST", u, rule)
	if err != nil {
		return nil, nil, err
	}

	newRule := new(Rule)
	resp, err := p.client.Do(req, newRule)
	if err != nil {
		return nil, resp, err
	}

	return newRule, resp, err
}

// Delete a rule
// Requires Policy ID from Policy object and Rule ID from Rule object
func (p *PoliciesService) DeleteRule(policyId string, ruleId string) (*Response, error) {
	u := fmt.Sprintf("policies/%vrules/%v", policyId, ruleId)
	req, err := p.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}
	resp, err := p.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, err
}

// Get a rule
// Requires Policy ID from Policy object and Rule ID from Rule object
func (p *PoliciesService) GetRule(policyId string, ruleId string) (*Rule, *Response, error) {
	u := fmt.Sprintf("policies/%v/rules/%v", policyId, ruleId)
	req, err := p.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	rule := new(Rule)
	resp, err := p.client.Do(req, rule)
	if err != nil {
		return nil, resp, err
	}

	return rule, resp, err
}

// Update a rule
// Requires Policy ID from Policy object and Rule ID from Rule object
// You must pass in the Rule object from Policies.Rule()
// This endpoint uses a PUT so I'm going to assume partial updates are not supported
func (p *PoliciesService) UpdateRule(policyId string, ruleId string, rule Rule) (*Rule, *Response, error) {
	u := fmt.Sprintf("policies/%v/rules/%v", policyId, ruleId)
	req, err := p.client.NewRequest("POST", u, rule)
	if err != nil {
		return nil, nil, err
	}

	updateRule := new(Rule)
	resp, err := p.client.Do(req, updateRule)
	if err != nil {
		return nil, resp, err
	}

	return updateRule, resp, err
}

// Activate a rule
// Requires Policy ID from Policy object and Rule ID from Rule object
func (p *PoliciesService) ActivateRule(policyId string, ruleId string) (*Response, error) {
	u := fmt.Sprintf("policies/%v/rules/%v/lifecycle/activate", policyId, ruleId)
	req, err := p.client.NewRequest("POST", u, nil)
	if err != nil {
		return nil, err
	}
	resp, err := p.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, err
}

// Deactivate a rule
// Requires Policy ID from Policy object and Rule ID from Rule object
func (p *PoliciesService) DeactivateRule(policyId string, ruleId string) (*Response, error) {
	u := fmt.Sprintf("policies/%v/rules/%v/lifecycle/deactivate", policyId, ruleId)
	req, err := p.client.NewRequest("POST", u, nil)
	if err != nil {
		return nil, err
	}
	resp, err := p.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, err
}
