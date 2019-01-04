package okta

import (
	"fmt"
	"net/url"
	"time"
)

const (
	eventTypeFilter = "eventType"
)

// LogProps is map that that okta uses everywhere
type LogProps map[string]interface{}

// LogGeo represents geo data for the log
type LogGeo struct {
	GeoLoc struct {
		Lat float64 `json:"lat"` // Optional
		Lon float64 `json:"lon"` // Optional

	} `json:"geolocation"` // Optional
	City       string `json:"city"`       // Optional
	State      string `json:"state"`      // Optional
	Country    string `json:"country"`    // Optional
	PostalCode string `json:"postalCode"` // Optional
}

// LogClient represents the client data from the log
type LogClient struct {
	UserAgent struct {
		Raw     string `json:"rawUserAgent"`
		OS      string `json:"os"` // Optional
		Browser string `json:"browser"`
	} `json:"userAgent"` // optional
	GeoContext LogGeo `json:"geographicalContext"` // Optional
	Zone       string `json:"zone"`                // Optional
	IPAdress   string `json:"ipAddress"`           // Optional
	Device     string `json:"device"`              // Optional
	ID         string `json:"id"`                  // Optional
}

// LogSecContext represents security info related to the ip of the log
type LogSecContext struct {
	AsNum   int    `json:"asNumber"` // Optional
	AsOrg   string `json:"asOrg"`    // Optional
	ISP     string `json:"isp"`      // Optional
	Domain  string `json:"domain"`   // Optional
	IsProxy bool   `json:"isProxy"`  // Optional
}

// LogIP represents ip info of the log
type LogIP struct {
	IP         string `json:"ip"`                  // Optional
	GeoContext LogGeo `json:"geographicalContext"` // Optional
	Version    string `json:"version"`             // one of V4, V6 Optional
	Source     string `json:"source"`              // Optional
}

// LogAuthContext contains metadata about how the actor was authenticated
type LogAuthContext struct {
	AuthProvider string `json:"authenticationProvider"` // one of OKTA_AUTHENTICATION_PROVIDER, ACTIVE_DIRECTORY, LDAP, FEDERATION, SOCIAL, FACTOR_PROVIDER, Optional
	CredProvider string `json:"credentialProvider"`     // one of OKTA_CREDENTIAL_PROVIDER, RSA, SYMANTEC, GOOGLE, DUO, YUBIKEY, Optional
	CredType     string `json:"credentialType"`         // one of OTP, SMS, PASSWORD, ASSERTION, IWA, EMAIL, OAUTH2, JWT, Optional
	Issuer       struct {
		ID   string `json:"id"`   // Optional
		Type string `json:"type"` // Optional

	} `json:"issuer"` // Optional
	SessID string `json:"externalSessionId"` // Optional
	Inter  string `json:"interface"`         // Optional i.e. Outlook, Office365, wsTrust
}

// LogActor Describes the user, app, client, or other entity (actor) who performed an action on a target
type LogActor struct {
	ID          string   `json:"id"`
	Type        string   `json:"type"`
	AlternateID string   `json:"alternateId"` // Optional - this usually seems to be a user email
	DisplayName string   `json:"displayName"` // Optional
	Details     LogProps `json:"detailEntry"`
}

func (a *LogActor) String() string {
	return fmt.Sprintf("id: %s, type: %s, alt_id: %s, display_name: %s, details: %v", a.ID, a.Type, a.AlternateID, a.DisplayName, a.Details)
}

// Log describes a single logged action or “event” performed by a set of actors for a set of targets.
type Log struct {
	UUID      string    `json:"uuid"`
	Published time.Time `json:"published"`
	EventType string    `json:"eventType"`
	Version   string    `json:"version"`
	//  one of DEBUG, INFO, WARN, ERROR
	Severity        string    `json:"severity"`
	LegacyEventType string    `json:"legacyEventType"` // Optional
	DisplayMessage  string    `json:"displayMessage"`  // Optional
	Actor           LogActor  `json:"actor"`
	Client          LogClient `json:"client"` // option
	Outcome         struct {
		Results string `json:"result"` // one of: SUCCESS, FAILURE, SKIPPED, UNKNOWN, Required
		Reason  string `json:"reason"` // Optional
	} `json:"outcome"` //  Optional
	// Targets The entity upon which an actor performs an action. Targets may be anything: an app user, a login token or anything else.
	Targets     []LogActor `json:"target"`
	Transaction struct {
		ID      string   `json:"id"`   // Optional
		Type    string   `json:"type"` // one of "WEB", "JOB", Optional
		Details LogProps `json:"detail"`
	} `json:"transaction"` // Optional
	DebugContext struct {
		// this always has a requestURI too
		DebugData LogProps `json:"debugData"`
	} `json:"debugContext"` // Optional
	AuthContext LogAuthContext `json:"authenticationContext"`
	SecContext  LogSecContext  `json:"securityContext"` // Optional
	Req         struct {
		IPChain []LogIP `json:"ipChain"` // Optional
	} `json:"request"` // Optional
}

// LogListFilterOptions is a struct that you can populate which will "filter" log searches. the exported field should allow you to do different filters based on what is allowed in the OKTA API.
type LogListFilterOptions struct {
	NextURL   *url.URL  `url:"-"`
	Limit     int       `url:"limit,omitempty"`
	Since     time.Time `url:"since,omitempty"`
	SortOrder string    `url:"sortOrder,omitempty"`
	Query     string    `url:"q,omitempty"`

	EventType string `url:"-"`

	FilterString string `url:"filter,omitempty"`
}

// LogsService handles communication with the log methods of the OKTA API
type LogsService service

// ListWithFilters will use the input list with filters to find logs return all log results that match that query. ListWithFilter doesn't support getting all logs at once bec it can be a ton of data
func (s *LogsService) ListWithFilters(opt *LogListFilterOptions) ([]Log, *Response, error) {
	var u string
	var err error

	if opt.NextURL != nil {
		u = opt.NextURL.String()
	} else {
		if opt.Limit == 0 {
			opt.Limit = defaultLimit
		}

		if opt.EventType != "" {
			opt.FilterString = appendToFilterString(opt.FilterString, eventTypeFilter, FilterEqualOperator, opt.EventType)
		}

		u, err = addOptions("logs", opt)
	}

	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("s.client.NewRequest %v", err)
	}
	var logs []Log
	resp, err := s.client.Do(req, &logs)
	if err != nil {
		return nil, resp, err
	}

	return logs, resp, nil
}
