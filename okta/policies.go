package okta

import (
	"time"
)

// PoliciesService handles communication with the Policy data related
// methods of the OKTA API.
type PoliciesService service

// Policy represents the Policy Object from the OKTA API
type Policy struct {
	ID          string      `json:"id,omitempty"`
	Type        string      `json:"type"` // OKTA_SIGN_ON, PASSWORD, MFA_ENROLL, OAUTH_AUTHORIZATION_POLICY
	Name        string      `json:"name"`
	System      string      `json:"system,omitempty"`      // true or false D=false
	Description string      `json:"description,omitempty"` // D=null
	Priority    int         `json:"priority,omitempty"`
	Status      string      `json:"status,omitempty"` // ACTIVE or INACTIVE D=ACTIVE
	Conditions  conditions  `json:"conditions,omitempty"`
	Settings    settings    `json:"settings,omitempty"`
	Created     time.Time   `json:"created,omitempty"`
	LastUpdated time.Time   `json:"lastUpdated,omitempty"`
	Links       policyLinks `json:"_links,omitempty"`
}

// the conditions that must be met during policy or rule evaluation
type conditions struct {
	People struct {
		Groups groups `json:"groups,omitempty"`
		Users  users  `json:"users,omitempty"`
	} `json:"people,omitempty"`
	AuthContext  string       `json:"authType,omitempty"` // ANY or RADIUS
	Network      network      `json:"network,omitempty"`
	AuthProvider authProvider `json:"authProvider,omitempty"`
}

// set of users to be included or excluded
type users struct {
	Include []string `json:"include,omitempty"`
	Exclude []string `json:"exclude,omitempty"`
}

// set of groups to be included or excluded
type groups struct {
	Include []string `json:"include,omitempty"`
	Exclude []string `json:"exclude,omitempty"`
}

// network selection mode, and a set of network zones to be included or excluded
type network struct {
	Connection string   `json:"connection,omitempty"` // ANYWHERE, ZONE, ON_NETWORK, or OFF_NETWORK
	Include    []string `json:"include,omitempty"`    // required if connect type is ZONE
	Exclude    []string `json:"exclude,omitempty"`    // required if connect type is ZONE
}

// specifies an authentication for users
type authProvider struct {
	Provider string   `json:"provider"`          // Okta or Active Directory D=Okta
	Include  []string `json:"include,omitempty"` // Include all AD integrations
}

// policy level settings for the particular policy type
type settings struct {
	factors struct {
		GoogleOtp    mfaFactor `json:"google_otp,omitempty"`
		OktaOtp      mfaFactor `json:"okta_otp,omitempty"`
		OktaPush     mfaFactor `json:"okta_push,omitempty"`
		OktaQuestion mfaFactor `json:"okta_question,omitempty"`
		OktaSms      mfaFactor `json:"okta_sms,omitempty"`
		RsaToken     mfaFactor `json:"rsa_token,omitempty"`
		SymantecVip  mfaFactor `json:"symantec_vip,omitempty"`
	} `json:"factors,omitempty"`
	password struct {
		password   password   `json:"password,omitempty"`
		recovery   recovery   `json:"recovery,omitempty"`
		delegation delegation `json:"lockout,omitempty"`
	} `json:"password,omitempty"`
}

type password struct {
	complexity complexity `json:"complexity,omitempty"`
	age        age        `json:"age,omitempty"`
	lockout    lockout    `json:"lockout,omitempty"`
}

type complexity struct {
	MinLength         int        `json:"minLength,omitempty"`         // D=8
	MinLowerCase      int        `json:"minLowerCase,omitempty"`      // D=1
	MinUpperCase      int        `json:"minUpperCase,omitempty"`      // D=1
	MinNumber         int        `json:"minNumber,omitempty"`         // D=1
	MinSymbol         int        `json:"minSymbol,omitempty"`         // D=1
	ExcludeUsername   bool       `json:"excludeUsername,omitempty"`   // D=true
	ExcludeAttributes []string   `json:"excludeAttributes,omitempty"` // firstname and lastname
	Dictionary        dictionary `json:"dictionary,omitempty"`
}

type dictionary struct {
	common common `json:"common,omitempty"`
}

type common struct {
	exclude bool `json:"excllude,omitempty"` // D=false
}

type age struct {
	MaxAgeDays     int `json:"maxAgeDays,omitempty"`
	ExpireWarnDays int `json:"expireWarnDays,omitempty"`
	MinAgeMinutes  int `json:"minAgeMinutes,omitempty"`
	HistoryCount   int `json:"historyCount,omitempty"`
}

type lockout struct {
	maxAttempts         int  `json:"maxAttempts,omitempty"`
	autoUnlockMinutes   int  `json:"autoUnlockMinutes,omitempty"`
	showLockoutFailures bool `json:"showLockoutFailures,omitempty"` // D=false
}

type recovery struct {
	factors struct {
	} `json:"factors,omitempty"`
}

type delegation struct {
}

type mfaFactor struct {
	Consent factorConsent `json:"consent,omitempty"`
	Enroll  factorEnroll  `json:"enroll,omitempty"`
}

type factorConsent struct {
	Terms factorConsentTerms `json:"terms,omitempty"`
	Type  string             `json:"type,omitempty"` // NONE or TERMS_OF_SERVICE
}

type factorEnroll struct {
	Self string `json:"self,omitempty"` // NOT_ALLOWED, OPTIONAL or REQUIRED D=NOT_ALLOWED
}

type factorConsentTerms struct {
	Format string `json:"format,omitempty"` // TEXT, RTF, MARKDOWN or URL
	Value  string `json:"value,omitempty"`  // D=NONE
}

// TODO add MFA conditions objects
// https://developer.okta.com/docs/api/resources/policy#policy-conditions-1

// TODO: policy links object is not complete?
// https://developer.okta.com/docs/api/resources/policy#LinksObject
type policyLinks struct {
	Self       string `json:"self"`
	Activate   string `json:"activate",omitempty`
	Deactivate string `json:"deactivate,omitempty"`
	Rules      string `json:"rules,omitempty"`
}

// Rules represent the Rules Object from the OKTA API
type Rules struct {
	ID          string     `json:"id,omitempty"`
	Type        string     `json:"type"`             // OKTA_SIGN_ON or PASSWORD or MFA_ENROLL
	Status      string     `json:"status,omitempty"` // ACTIVE or INACTIVE D=Active
	Priority    int        `json:"priority,omitempty"`
	System      string     `json:"system,omitempty"` // true or false D=false
	Created     time.Time  `json:"created,omitempty"`
	LastUpdated time.Time  `json:"lastUpdated,omitempty"`
	Conditions  conditions `json:"conditions,omitempty"`
	Actions     actions    `json:"actions,omitempty"`
	Links       ruleLinks  `json:"_links,omitempty"`
}

type actions struct {
	signon struct {
		Access                  string  `json:"access"`                            // ALLOW or DENY
		RequireFactor           bool    `json:"requireFactor,omitempty"`           // D=false
		FactorPromptMode        string  `json:"factorPromptMode,omitempty"`        // DEVICE, SESSION or ALWAYS
		RememberDeviceByDefault bool    `json:"rememberDeviceByDefault,omitempty"` // D=false
		FactorLifetime          int     `json:"factorLifetime,omitempty"`
		Session                 session `json:"session,omitempty"`
	} `json:"signon,omitempty"`
	enroll struct {
		Self string `json:"self,omitempty"` // CHALLENGE, LOGIN or NEVER
	} `json:"enroll,omitempty"`
}

type session struct {
	MaxSessionIdleMinutes     int  `json:"maxSessionIdleMinutes,omitempty"` // D=120
	MaxSessionLifetimeMinutes int  `json:"maxSessionLifetimeMinutes,omitempty"`
	UsePersistentCookie       bool `json:"usePersistentCookie,,omitempty"` // D=false
}

// TODO: rule links object is not complete
// https://developer.okta.com/docs/api/resources/policy#RulesLinksObject
type ruleLinks struct {
	Self       string `json:"self"`
	Activate   string `json:"activate,omitempty"`
	Deactivate string `json:"deactivate,omitempty"`
}
