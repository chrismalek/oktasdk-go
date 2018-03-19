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
	ID          string    `json:"id,omitempty"`
	Type        string    `json:"type,omitempty"`
	Name        string    `json:"name,omitempty"`
	System      bool      `json:"system,omitempty"`
	Description string    `json:"description,omitempty"`
	Priority    int       `json:"priority,omitempty"`
	Status      string    `json:"status,omitempty"`
	Created     time.Time `json:"created,omitempty"`
	LastUpdated time.Time `json:"lastUpdated,omitempty"`
	Conditions  struct {
		*People       `json:"people,omitempty"`
		AuthType      *string `json:"authType,omitempty"`
		*Network      `json:"network,omitempty"`
		*AuthProvider `json:"authProvider,omitempty"`
	} `json:"conditions,omitempty"`
	Settings struct {
		*Factors    `json:"factors,omitempty"`
		*Password   `json:"password,omitempty"`
		*Recovery   `json:"recovery,omitempty"`
		*Delegation `json:"delegation,omitempty"`
	} `json:"settings,omitempty"`
	Links `json:"_links,omitempty"`
}

// Mfa policy factors obj
type Factors struct {
	GoogleOtp struct {
		Consent `json:"consent,omitempty"`
		Enroll  `json:"enroll,omitempty"`
	} `json:"google_otp,omitempty"`
	OktaOtp struct {
		Consent `json:"consent,omitempty"`
		Enroll  `json:"enroll,omitempty"`
	} `json:"okta_otp,omitempty"`
	OktaPush struct {
		Consent `json:"consent,omitempty"`
		Enroll  `json:"enroll,omitempty"`
	} `json:"okta_push,omitempty"`
	OktaQuestion struct {
		Consent `json:"consent,omitempty"`
		Enroll  `json:"enroll,omitempty"`
	} `json:"okta_question,omitempty"`
	OktaSms struct {
		Consent `json:"consent,omitempty"`
		Enroll  `json:"enroll,omitempty"`
	} `json:"okta_sms,omitempty"`
	RsaToken struct {
		Consent `json:"consent,omitempty"`
		Enroll  `json:"enroll,omitempty"`
	} `json:"rsa_token,omitempty"`
	SymantecVip struct {
		Consent `json:"consent,omitempty"`
		Enroll  `json:"enroll,omitempty"`
	} `json:"symantec_vip,omitempty"`
}

// Mfa policy factor consent obj
type Consent struct {
	Terms struct {
		Format *string `json:"format,omitempty"`
		Value  *string `json:"value,omitempty"`
	} `json:"terms,omitempty"`
	Type *string `json:"type,omitempty"`
}

// Mfa policy factor enroll obj
type Enroll struct {
	Self *string `json:"self,omitempty"`
}

// password policy obj
type Password struct {
	Complexity struct {
		MinLength         *int      `json:"minLength,omitempty"`
		MinLowerCase      *int      `json:"minLowerCase,omitempty"`
		MinUpperCase      *int      `json:"minUpperCase,omitempty"`
		MinNumber         *int      `json:"minNumber,omitempty"`
		MinSymbol         *int      `json:"minSymbol,omitempty"`
		ExcludeUsername   *bool     `json:"excludeUsername,omitempty"`
		ExcludeAttributes *[]string `json:"excludeAttributes,omitempty"`
		Dictionary        struct {
			Common struct {
				Exclude *bool `json:"exclude,omitempty"`
			} `json:"common,omitempty"`
		} `json:"dictionary,omitempty"`
	} `json:"complexity,omitempty"`
	Age struct {
		MaxAgeDays     *int `json:"maxAgeDays,omitempty"`
		ExpireWarnDays *int `json:"expireWarnDays,omitempty"`
		MinAgeMinutes  *int `json:"minAgeMinutes,omitempty"`
		HistoryCount   *int `json:"historyCount,omitempty"`
	} `json:"age,omitempty"`
	Lockout struct {
		MaxAttempts         *int  `json:"maxAttempts,omitempty"`
		AutoUnlockMinutes   *int  `json:"autoUnlockMinutes,omitempty"`
		ShowLockoutFailures *bool `json:"showLockoutFailures,omitempty"`
	} `json:"lockout,omitempty"`
}

// passowrd policy recover obj
type Recovery struct {
	Factors struct {
		RecoveryQuestion struct {
			Status     *string `json:"status,omitempty"`
			Properties struct {
				Complexity struct {
					MinLength *int `json:"minLength,omitempty"`
				} `json:"complexity,omitempty"`
			} `json:"properties,omitempty"`
		} `json:"recovery_question,omitempty"`
		OktaEmail struct {
			Status     *string `json:"status,omitempty"`
			Properties struct {
				RecoveryToken struct {
					TokenLifetimeMinutes *int `json:"tokenLifetimeMinutes,omitempty"`
				} `json:"recoveryToken,omitempty"`
			} `json:"properties,omitempty"`
		} `json:"okta_email,omitempty"`
		OktaSms struct {
			Status *string `json:"status,omitempty"`
		} `json:"okta_sms,omitempty"`
	} `json:"factors,omitempty"`
}

// password policy delegation obj
type Delegation struct {
	Options struct {
		SkipUnlock *bool `json:"skipUnlock,omitempty"`
	} `json:"options,omitempty"`
}

// policy & rule conditions people obj
type People struct {
	Groups struct {
		Include *[]string `json:"include,omitempty"`
		Exclude *[]string `json:"exclude,omitempty"`
	} `json:"groups,omitempty"`
	Users struct {
		Include *[]string `json:"include,omitempty"`
		Exclude *[]string `json:"exclude,omitempty"`
	} `json:"users,omitempty"`
}

// policy & rule conditions network obj
type Network struct {
	Connection *string   `json:"connection,omitempty"`
	Include    *[]string `json:"include,omitempty"`
	Exclude    *[]string `json:"exclude,omitempty"`
}

// policy & rule authProvider obj
type AuthProvider struct {
	Provider *string   `json:"provider,omitempty"`
	Include  *[]string `json:"include,omitempty"`
}

// GetPolicesByType returns an array of Policy objs
type policies struct {
	Policies []Policy `json:",-"`
}

// Policy & Rule obj use the same links obj
type Links struct {
	Self struct {
		Href  string `json:"href,omitempty"`
		Hints struct {
			Allow []string `json:"allow,omitempty"`
		} `json:"hints,omitempty"`
	} `json:"self,omitempty"`
	Activate struct {
		Href  string `json:"href,omitempty"`
		Hints struct {
			Allow []string `json:"allow,omitempty"`
		} `json:"hints,omitempty"`
	} `json:"activate",omitempty`
	Deactivate struct {
		Href  string `json:"href,omitempty"`
		Hints struct {
			Allow []string `json:"allow,omitempty"`
		} `json:"hints,omitempty"`
	} `json:"deactivate,omitempty"`
	Rules struct {
		Href  string `json:"href,omitempty"`
		Hints struct {
			Allow []string `json:"allow,omitempty"`
		} `json:"hints,omitempty"`
	} `json:"rules,omitempty"`
}

// Rule represents the Rule Object from the OKTA API
type Rule struct {
	ID          string    `json:"id,omitempty"`
	Type        string    `json:"type,omitempty"`
	Status      string    `json:"status,omitempty"`
	Priority    int       `json:"priority,omitempty"`
	System      bool      `json:"system,omitempty"`
	Created     time.Time `json:"created,omitempty"`
	LastUpdated time.Time `json:"lastUpdated,omitempty"`
	Conditions  string    `json:"conditions,omitempty"`
	Actions     struct {
		signon struct {
			Access                  string `json:"access,omitempty"`
			RequireFactor           bool   `json:"requireFactor,omitempty"`
			FactorPromptMode        string `json:"factorPromptMode,omitempty"`
			RememberDeviceByDefault bool   `json:"rememberDeviceByDefault,omitempty"`
			FactorLifetime          int    `json:"factorLifetime,omitempty"`
			Session                 struct {
				MaxSessionIdleMinutes     int  `json:"maxSessionIdleMinutes,omitempty"`
				MaxSessionLifetimeMinutes int  `json:"maxSessionLifetimeMinutes,omitempty"`
				UsePersistentCookie       bool `json:"usePersistentCookie,omitempty"`
			} `json:"session,omitempty"`
		} `json:"signon,omitempty"`
		enroll struct {
			Self string `json:"self,omitempty"`
		} `json:"enroll,omitempty"`
		PasswordChange           passwordAction `json:"passwordChange,omitempty"`
		SelfServicePasswordReset passwordAction `json:"selfServicePasswordReset,omitempty"`
		SelfServiceUnlock        passwordAction `json:"selfServiceUnlock,omitempty"`
	} `json:"actions,omitempty"`
	Links `json:"_links,omitempty"`
}

// rule actions for passwords use the same passwordAction obj
type passwordAction struct {
	Access string `json:"access,omitempty"`
}

// GetPolicyRules returns an array of Rule objs
type rules struct {
	Rules []Rule `json:",-"`
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
func (p *PoliciesService) CreatePolicy(policy interface{}) (*Policy, *Response, error) {
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
