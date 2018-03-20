package okta

import (
	"fmt"
	"time"
)

// PoliciesService handles communication with the Policy data related
// methods of the OKTA API.
type PoliciesService service

// Return the InputPolicy object. Used to create & update policies
func (p *PoliciesService) InputPolicy() InputPolicy {
	return InputPolicy{}
}

// Return the SignOnPolicy object. Used to create & update the signon policy
// this policy is a separate struct because it does not include settings
func (p *PoliciesService) SignOnPolicy() SignOnPolicy {
	return SignOnPolicy{}
}

// Return the Rule object. Used to create & update rules
func (p *PoliciesService) InputRule() InputRule {
	return InputRule{}
}

// Return the Password object.
// Used to create & update the password policy obj password settings
func (p *PoliciesService) Password() Password {
	return Password{}
}

// Return the Recovery object.
// Used to create & update the recovery policy obj recovery settings
func (p *PoliciesService) Recovery() Recovery {
	return Recovery{}
}

// Return the Delegation object.
// Used to create & update the delegation policy obj delegation settings
func (p *PoliciesService) Delegation() Delegation {
	return Delegation{}
}

// Return the SignOn object.
// Used to create & update the signon rules obj signon actions
func (p *PoliciesService) SignOn() SignOn {
	return SignOn{}
}

// Return the Enroll object.
// Used to create & update the enroll rules obj enroll actions
func (p *PoliciesService) Enroll() Enroll {
	return Enroll{}
}

// Return the PasswordAction object.
// Used to create & update the passwordaction rules obj passwordaction actions
func (p *PoliciesService) PasswordAction() PasswordAction {
	return PasswordAction{}
}

// InputPolicy represents the Policy Object from the OKTA API
// used to create or update a password policy
type InputPolicy struct {
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
		People struct {
			*Groups `json:"groups,omitempty"`
			*Users  `json:"users,omitempty"`
		} `json:"people,omitempty"`
		AuthProvider `json:"authProvider,omitempty"`
	} `json:"conditions,omitempty"`
	Settings struct {
		*Factors    `json:"factors,omitempty"`
		*Password   `json:"password,omitempty"`
		*Recovery   `json:"recovery,omitempty"`
		*Delegation `json:"delegation,omitempty"`
	} `json:"settings,omitempty"`
}

// PeopleConditionUsers updates the People Users condition for the InputPolicy
// requires inputs string "include" or "exclude" & a string slice of Okta users
func (p *InputPolicy) PeopleConditionUsers(clude string, values []string) error {
	var pop *Users
	if p.Conditions.People.Users == nil {
		pop = new(Users)
	} else {
		pop = p.Conditions.People.Users
	}
	switch {
	case clude == "include":
		pop.Include = &values
	case clude == "exclude":
		pop.Exclude = &values
	default:
		return fmt.Errorf("[ERROR] clude input var supports values \"include\" or \"exclude\"")
	}
	p.Conditions.People.Users = pop
	return nil
}

// PeopleConditionGroups updates the People Groups condition for the InputPolicy
// requires inputs string "include" or "exclude" & a string slice of Okta groups
func (p *InputPolicy) PeopleConditionGroups(clude string, values []string) error {
	var pop *Groups
	if p.Conditions.People.Groups == nil {
		pop = new(Groups)
	} else {
		pop = p.Conditions.People.Groups
	}
	switch {
	case clude == "include":
		pop.Include = &values
	case clude == "exclude":
		pop.Exclude = &values
	default:
		return fmt.Errorf("[ERROR] clude input var supports values \"include\" or \"exclude\"")
	}
	p.Conditions.People.Groups = pop
	return nil
}

// PasswordSettings updates the password settings for the InputPolicy
// requires input the Password struct with the values to update
func (p *InputPolicy) PasswordSettings(settings Password) {
	var pop *Password
	if p.Settings.Password == nil {
		pop = new(Password)
	} else {
		pop = p.Settings.Password
	}
	pop = &settings
	p.Settings.Password = pop
}

// RecoverySettings updates the recovery settings for the InputPolicy
// requires input the Recovery struct with the values to update
func (p *InputPolicy) RecoverySettings(settings Recovery) {
	var pop *Recovery
	if p.Settings.Recovery == nil {
		pop = new(Recovery)
	} else {
		pop = p.Settings.Recovery
	}
	pop = &settings
	p.Settings.Recovery = pop
}

// DelegationSettings updates the delegation settings for the InputPolicy
// requires input the Delegation struct with the values to update
func (p *InputPolicy) DelegationSettings(settings Delegation) {
	var pop *Delegation
	if p.Settings.Delegation == nil {
		pop = new(Delegation)
	} else {
		pop = p.Settings.Delegation
	}
	pop = &settings
	p.Settings.Delegation = pop
}

// SignOnPolicy represents the Policy Object from the OKTA API
// used to create or update a signon policy
// this policy is a separate struct because it does not include settings
// TODO: policy signon should use the InputPolicy struct, not sure how atm
type SignOnPolicy struct {
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
		People struct {
			*Groups `json:"groups,omitempty"`
			*Users  `json:"users,omitempty"`
		} `json:"people,omitempty"`
	} `json:"conditions,omitempty"`
}

// SignOnPeopleConditionUsers updates the People Users condition for the SignOnPolicy
// requires inputs string "include" or "exclude" & a string slice of Okta users
func (p *SignOnPolicy) SignOnPeopleConditionUsers(clude string, values []string) error {
	var pop *Users
	if p.Conditions.People.Users == nil {
		pop = new(Users)
	} else {
		pop = p.Conditions.People.Users
	}
	switch {
	case clude == "include":
		pop.Include = &values
	case clude == "exclude":
		pop.Exclude = &values
	default:
		return fmt.Errorf("[ERROR] clude input var supports values \"include\" or \"exclude\"")
	}
	p.Conditions.People.Users = pop
	return nil
}

// SIgnOnPeopleConditionGroups updates the People Groups condition for the SignOnPolicy
// requires inputs string "include" or "exclude" & a string slice of Okta groups
func (p *SignOnPolicy) SignOnPeopleConditionGroups(clude string, values []string) error {
	var pop *Groups
	if p.Conditions.People.Groups == nil {
		pop = new(Groups)
	} else {
		pop = p.Conditions.People.Groups
	}
	switch {
	case clude == "include":
		pop.Include = &values
	case clude == "exclude":
		pop.Exclude = &values
	default:
		return fmt.Errorf("[ERROR] clude input var supports values \"include\" or \"exclude\"")
	}
	p.Conditions.People.Groups = pop
	return nil
}

// Policy represents the complete Policy Object from the OKTA API
// used to return policy data (no need for pointers)
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
		People       `json:"people,omitempty"`
		AuthType     string `json:"authType,omitempty"`
		Network      `json:"network,omitempty"`
		AuthProvider `json:"authProvider,omitempty"`
	} `json:"conditions,omitempty"`
	Settings struct {
		Factors    `json:"factors,omitempty"`
		Password   `json:"password,omitempty"`
		Recovery   `json:"recovery,omitempty"`
		Delegation `json:"delegation,omitempty"`
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
		Format string `json:"format,omitempty"`
		Value  string `json:"value,omitempty"`
	} `json:"terms,omitempty"`
	Type string `json:"type,omitempty"`
}

// Mfa policy & rule factor enroll obj
type Enroll struct {
	Self string `json:"self,omitempty"`
}

// password policy obj
type Password struct {
	Complexity struct {
		MinLength         int      `json:"minLength,omitempty"`
		MinLowerCase      int      `json:"minLowerCase,omitempty"`
		MinUpperCase      int      `json:"minUpperCase,omitempty"`
		MinNumber         int      `json:"minNumber,omitempty"`
		MinSymbol         int      `json:"minSymbol,omitempty"`
		ExcludeUsername   bool     `json:"excludeUsername,omitempty"`
		ExcludeAttributes []string `json:"excludeAttributes,omitempty"`
		Dictionary        struct {
			Common struct {
				Exclude bool `json:"exclude,omitempty"`
			} `json:"common,omitempty"`
		} `json:"dictionary,omitempty"`
	} `json:"complexity,omitempty"`
	Age struct {
		MaxAgeDays     int `json:"maxAgeDays,omitempty"`
		ExpireWarnDays int `json:"expireWarnDays,omitempty"`
		MinAgeMinutes  int `json:"minAgeMinutes,omitempty"`
		HistoryCount   int `json:"historyCount,omitempty"`
	} `json:"age,omitempty"`
	Lockout struct {
		MaxAttempts         int  `json:"maxAttempts,omitempty"`
		AutoUnlockMinutes   int  `json:"autoUnlockMinutes,omitempty"`
		ShowLockoutFailures bool `json:"showLockoutFailures,omitempty"`
	} `json:"lockout,omitempty"`
}

// passowrd policy recover obj
type Recovery struct {
	Factors struct {
		RecoveryQuestion struct {
			Status     string `json:"status,omitempty"`
			Properties struct {
				Complexity struct {
					MinLength int `json:"minLength,omitempty"`
				} `json:"complexity,omitempty"`
			} `json:"properties,omitempty"`
		} `json:"recovery_question,omitempty"`
		OktaEmail struct {
			Status     string `json:"status,omitempty"`
			Properties struct {
				RecoveryToken struct {
					TokenLifetimeMinutes int `json:"tokenLifetimeMinutes,omitempty"`
				} `json:"recoveryToken,omitempty"`
			} `json:"properties,omitempty"`
		} `json:"okta_email,omitempty"`
		OktaSms struct {
			Status string `json:"status,omitempty"`
		} `json:"okta_sms,omitempty"`
	} `json:"factors,omitempty"`
}

// password policy delegation obj
type Delegation struct {
	Options struct {
		SkipUnlock bool `json:"skipUnlock,omitempty"`
	} `json:"options,omitempty"`
}

// policy & rule conditions people obj
type People struct {
	Groups struct {
		Include []string `json:"include,omitempty"`
		Exclude []string `json:"exclude,omitempty"`
	} `json:"groups,omitempty"`
	Users struct {
		Include []string `json:"include,omitempty"`
		Exclude []string `json:"exclude,omitempty"`
	} `json:"users,omitempty"`
}

// policy & rule conditions people groups obj
// when creating an obj, Include & Exclude are exclusive
type Groups struct {
	Include *[]string `json:"include,omitempty"`
	Exclude *[]string `json:"exclude,omitempty"`
}

// policy & rule conditions people users obj
// when creating an obj, Include & Exclude are exclusive
type Users struct {
	Include *[]string `json:"include,omitempty"`
	Exclude *[]string `json:"exclude,omitempty"`
}

// policy & rule conditions network obj
type Network struct {
	Connection string   `json:"connection,omitempty"`
	Include    []string `json:"include,omitempty"`
	Exclude    []string `json:"exclude,omitempty"`
}

// policy & rule authProvider obj
type AuthProvider struct {
	Provider string   `json:"provider,omitempty"`
	Include  []string `json:"include,omitempty"`
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

// InputRule represents the Rule Object from the OKTA API
// used to create or update rules
type InputRule struct {
	ID          string    `json:"id,omitempty"`
	Type        string    `json:"type,omitempty"`
	Status      string    `json:"status,omitempty"`
	Priority    int       `json:"priority,omitempty"`
	System      bool      `json:"system,omitempty"`
	Created     time.Time `json:"created,omitempty"`
	LastUpdated time.Time `json:"lastUpdated,omitempty"`
	Conditions  struct {
		People struct {
			*Groups `json:"groups,omitempty"`
			*Users  `json:"users,omitempty"`
		} `json:"people,omitempty"`
		AuthType string `json:"authType,omitempty"`
		Network  struct {
			Connection string `json:"connection,omitempty"`
			// TODO: Include & Exclude not supported as only needed when
			// Connection is "ZONE". zone requires the zone api (not implemented atm)
			Include *[]string `json:"include,omitempty"`
			Exclude *[]string `json:"exclude,omitempty"`
		} `json:"network,omitempty"`
	} `json:"conditions,omitempty"`
	Actions struct {
		*SignOn                  `json:"session,omitempty"`
		*Enroll                  `json:"enroll,omitempty"`
		PasswordChange           *PasswordAction `json:"passwordChange,omitempty"`
		SelfServicePasswordReset *PasswordAction `json:"selfServicePasswordReset,omitempty"`
		SelfServiceUnlock        *PasswordAction `json:"selfServiceUnlock,omitempty"`
	} `json:"actions,omitempty"`
}

// PeopleRuleUsers updates the People Users condition for the InputRule
// requires inputs string "include" or "exclude" & a string slice of Okta users
func (p *InputRule) PeopleRuleUsers(clude string, values []string) error {
	var pop *Users
	if p.Conditions.People.Users == nil {
		pop = new(Users)
	} else {
		pop = p.Conditions.People.Users
	}
	switch {
	case clude == "include":
		pop.Include = &values
	case clude == "exclude":
		pop.Exclude = &values
	default:
		return fmt.Errorf("[ERROR] clude input var supports values \"include\" or \"exclude\"")
	}
	p.Conditions.People.Users = pop
	return nil
}

// PeopleRuleGroups updates the People Groups condition for the InputRule
// requires inputs string "include" or "exclude" & a string slice of Okta groups
func (p *InputRule) PeopleRuleGroups(clude string, values []string) error {
	var pop *Groups
	if p.Conditions.People.Groups == nil {
		pop = new(Groups)
	} else {
		pop = p.Conditions.People.Groups
	}
	switch {
	case clude == "include":
		pop.Include = &values
	case clude == "exclude":
		pop.Exclude = &values
	default:
		return fmt.Errorf("[ERROR] clude input var supports values \"include\" or \"exclude\"")
	}
	p.Conditions.People.Groups = pop
	return nil
}

// SignOnActions updates the signon actions for the InputRule
// requires input the SignOn struct with the values to update
func (p *InputRule) SignOnActions(settings SignOn) {
	var pop *SignOn
	if p.Actions.SignOn == nil {
		pop = new(SignOn)
	} else {
		pop = p.Actions.SignOn
	}
	pop = &settings
	p.Actions.SignOn = pop
}

// EnrollActions updates the enroll actions for the InputRule
// requires input the Enroll struct with the values to update
func (p *InputRule) EnrollActions(settings Enroll) {
	var pop *Enroll
	if p.Actions.Enroll == nil {
		pop = new(Enroll)
	} else {
		pop = p.Actions.Enroll
	}
	pop = &settings
	p.Actions.Enroll = pop
}

// PasswordChangeActions updates the passwordchange actions for the InputRule
// requires input the PasswordAction struct with the values to update
func (p *InputRule) PasswordChangeActions(settings PasswordAction) {
	var pop *PasswordAction
	if p.Actions.PasswordChange == nil {
		pop = new(PasswordAction)
	} else {
		pop = p.Actions.PasswordChange
	}
	pop = &settings
	p.Actions.PasswordChange = pop
}

// SelfServicePasswordResetActions updates the selfservicepasswordreset actions for the InputRule
// requires input the PasswordAction struct with the values to update
func (p *InputRule) SelfServicePasswordResetActions(settings PasswordAction) {
	var pop *PasswordAction
	if p.Actions.SelfServicePasswordReset == nil {
		pop = new(PasswordAction)
	} else {
		pop = p.Actions.SelfServicePasswordReset
	}
	pop = &settings
	p.Actions.SelfServicePasswordReset = pop
}

// SelfServiceUnlockActions updates the selfserviceunlock actions for the InputRule
// requires input the PasswordAction struct with the values to update
func (p *InputRule) SelfServiceUnlockActions(settings PasswordAction) {
	var pop *PasswordAction
	if p.Actions.SelfServiceUnlock == nil {
		pop = new(PasswordAction)
	} else {
		pop = p.Actions.SelfServiceUnlock
	}
	pop = &settings
	p.Actions.SelfServiceUnlock = pop
}

// Rule represents the complete Rule Object from the OKTA API
// used to return rule data (no need for pointers)
type Rule struct {
	ID          string    `json:"id,omitempty"`
	Type        string    `json:"type,omitempty"`
	Status      string    `json:"status,omitempty"`
	Priority    int       `json:"priority,omitempty"`
	System      bool      `json:"system,omitempty"`
	Created     time.Time `json:"created,omitempty"`
	LastUpdated time.Time `json:"lastUpdated,omitempty"`
	Conditions  struct {
		People       `json:"people,omitempty"`
		AuthType     string `json:"authType,omitempty"`
		Network      `json:"network,omitempty"`
		AuthProvider `json:"authProvider,omitempty"`
	} `json:"conditions,omitempty"`
	Actions struct {
		SignOn                   `json:"session,omitempty"`
		Enroll                   `json:"enroll,omitempty"`
		PasswordChange           PasswordAction `json:"passwordChange,omitempty"`
		SelfServicePasswordReset PasswordAction `json:"selfServicePasswordReset,omitempty"`
		SelfServiceUnlock        PasswordAction `json:"selfServiceUnlock,omitempty"`
	} `json:"actions,omitempty"`
	Links `json:"_links,omitempty"`
}

type SignOn struct {
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
}

// rule actions for passwords use the same passwordAction obj
type PasswordAction struct {
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
