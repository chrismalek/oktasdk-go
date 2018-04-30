package okta

import (
	"fmt"
	"time"
)

// SchemasService handles communication with the Schema data related
// methods of the OKTA API.
type SchemasService service

type Schema struct {
	ID          string      `json:"id"`
	Schema      string      `json:"$schema"`
	Name        string      `json:"name"`
	Title       string      `json:"title"`
	Created     time.Time   `json:"created"`
	LastUpdated time.Time   `json:"lastUpdated"`
	Definitions Definitions `json:"definitions"`
	Type        string      `json:"type"`
	Properties  Properties  `json:"properties"`
}

type Definitions struct {
	Base struct {
		ID         string                 `json:"id"`
		Type       string                 `json:"type"`
		Properties map[string]interface{} `json:"properties"`
		Required   []string               `json:"required"`
	} `json:"base"`
	Custom struct {
		ID   string `json:"id"`
		Type string `json:"type"`
		Properties map[string]interface{} `json:"properties"`
		Required   []string             `json:"required"`
	} `json:"custom"`
}

type Properties struct {
	Profile struct {
		AllOf []interface{} `json:"allOf"`
	} `json:"profile"`
}

type SubSchema struct {
	Title       string        `json:"title"`
	Type        string        `json:"type"`
	Description string        `json:"description,omitempty"`
	Format      string        `json:"format,omitempty"`
	Required    bool          `json:"required,omitempty"`
	Mutability  string        `json:"mutablity,omitempty"`
	Scope       string        `json:"scope,omitempty"`
	MinLength   int           `json:"minLength,omitempty"`
	MaxLength   int           `json:"maxLength,omitempty"`
	Enum        []string      `json:"enum,omitempty"`
	OneOf       []interface{} `json:"oneOf,omitempty"`
	Permissions []struct {
		Principal string `json:"principal"`
		Action    string `json:"action"`
	} `json:"permissions"`
	Master struct {
		Type string `json:"type,omitempty"`
	} `json:"master,omitempty"`
}

func (s *SchemasService) GetUserSchema() (*Schema, *Response, error) {
	u := fmt.Sprintf("meta/schemas/user/default")
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	schema := new(Schema)
	resp, err := s.client.Do(req, schema)
	if err != nil {
		return nil, resp, err
	}

	return schema, resp, err
}
