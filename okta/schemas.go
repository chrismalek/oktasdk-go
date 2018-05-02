package okta

import (
	"fmt"
	"time"
)

// SchemasService handles communication with the Schema data related
// methods of the OKTA API.
type SchemasService service

type Schema struct {
	ID          string    `json:"id"`
	Schema      string    `json:"$schema"`
	Name        string    `json:"name"`
	Title       string    `json:"title"`
	Created     time.Time `json:"created"`
	LastUpdated time.Time `json:"lastUpdated"`
	Definitions struct {
		Base struct {
			ID         string                 `json:"id"`
			Type       string                 `json:"type"`
			Properties map[string]interface{} `json:"properties"`
			Required   []string               `json:"required"`
		} `json:"base"`
		Custom struct {
			ID         string                 `json:"id"`
			Type       string                 `json:"type"`
			Properties map[string]interface{} `json:"properties"`
			Required   []string               `json:"required"`
		} `json:"custom"`
	} `json:"definitions"`
	Type       string `json:"type"`
	Properties struct {
		Profile struct {
			AllOf []interface{} `json:"allOf"`
		} `json:"profile"`
	} `json:"properties"`
}

type BaseSubSchema struct {
	Base struct {
		ID         string `json:"id"`
		Type       string `json:"type"`
		Properties struct {
			Title       string        `json:"title"`
			Type        string        `json:"type"`
			Format      string        `json:"format,omitempty"`
			Required    bool          `json:"required,omitempty"`
			Mutability  string        `json:"mutablity,omitempty"`
			Scope       string        `json:"scope,omitempty"`
			MinLength   int           `json:"minLength,omitempty"`
			MaxLength   int           `json:"maxLength,omitempty"`
			Permissions []Permissions `json:"permissions"`
			Master      struct {
				Type string `json:"type"`
			} `json:"master"`
		} `json:"properties"`
		Required []string `json:"required"`
	} `json:"base"`
}

type CustomSubSchema struct {
	Custom struct {
		ID         string `json:"id"`
		Type       string `json:"type"`
		Properties struct {
			Title       string        `json:"title"`
			Type        string        `json:"type"`
			Description string        `json:"description,omitempty"`
			Format      string        `json:"format,omitempty"`
			Required    bool          `json:"required,omitempty"`
			Mutability  string        `json:"mutablity,omitempty"`
			Scope       string        `json:"scope,omitempty"`
			MinLength   int           `json:"minLength,omitempty"`
			MaxLength   int           `json:"maxLength,omitempty"`
			Items       []interface{} `json:"items,omitempty"`
			Union       string        `json:"union,omitempty"`
			Enum        []string      `json:"enum,omitempty"`
			OneOf       []interface{} `json:"oneOf,omitempty"`
			Permissions []Permissions `json:"permissions"`
			Master      struct {
				Type string `json:"type"`
			} `json:"master"`
		} `json:"properties"`
		Required []string `json:"required"`
	} `json:"custom"`
}

type Permissions struct {
	Principal string `json:"principal"`
	Action    string `json:"action"`
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

func (s *SchemasService) GetUserSchemaIndex(scope string) ([]string, *Response, error) {
	var prop map[string]interface{}
	var index []string
	schema, resp, err := s.client.Schemas.GetUserSchema()
	if err != nil {
		return nil, resp, err
	}
	switch {
	case scope == "base":
		prop = schema.Definitions.Base.Properties
	case scope == "custom":
		prop = schema.Definitions.Custom.Properties
	default:
		return nil, resp, fmt.Errorf("[ERROR] GetUserSchemaIndex input string var supports values \"base\" or \"custom\"")
	}
	for key := range prop {
		index = append(index, key)
	}
	return index, resp, err
}

func (s *SchemasService) GetUserBaseSubSchema(title string) (*BaseSubSchema, *Response, error) {
	subSchema := new(BaseSubSchema)
	schema, resp, err := s.client.Schemas.GetUserSchema()
	list, resp, err := s.client.Schemas.GetUserSchemaIndex("base")
	for _, key := range list {
		if key == title {
			subSchema.Base.ID = "#base"
			subSchema.Base.Type = "object"
			subSchema.Base.Required = schema.Definitions.Base.Required
			temp := schema.Definitions.Base.Properties[title]
			for k, v := range temp.(map[string]interface{}) {
				switch {
				case k == "title":
					subSchema.Base.Properties.Title = v.(string)
				case k == "type":
					subSchema.Base.Properties.Type = v.(string)
				case k == "format":
					subSchema.Base.Properties.Format = v.(string)
				case k == "required":
					subSchema.Base.Properties.Required = v.(bool)
				case k == "mutability":
					subSchema.Base.Properties.Mutability = v.(string)
				case k == "scope":
					subSchema.Base.Properties.Scope = v.(string)
				case k == "minLength":
					subSchema.Base.Properties.MinLength = int(v.(float64))
				case k == "maxLength":
					subSchema.Base.Properties.MaxLength = int(v.(float64))
				case k == "permissions":
					perms := new(Permissions)
					for _, j := range v.([]interface{}) {
						for h, x := range j.(map[string]interface{}) {
							switch {
							case h == "principal":
								perms.Principal = x.(string)
							case h == "action":
								perms.Action = x.(string)
							}
						}
					}
					subSchema.Base.Properties.Permissions = append(subSchema.Base.Properties.Permissions, *perms)
				case k == "master":
					for g, z := range v.(map[string]interface{}) {
						switch {
						case g == "type":
							subSchema.Base.Properties.Master.Type = z.(string)
						}
					}
				}
			}
			break
		}
	}
	return subSchema, resp, err
}
