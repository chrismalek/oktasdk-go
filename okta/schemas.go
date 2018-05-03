package okta

import (
	"fmt"
	"time"
)

// SchemasService handles communication with the Schema data related
// methods of the OKTA API.
type SchemasService service

type Schema struct {
	ID          string
	Schema      string
	Name        string
	Title       string
	Created     time.Time
	LastUpdated time.Time
	Definitions struct {
		Base struct {
			ID   string
			Type string
			Properties []BaseSubSchema
			Required []string
		}
		Custom struct {
			ID   string
			Type string
			Properties []CustomSubSchema
			Required []string
		}
	}
	Type string
}

type BaseSubSchema struct {
	Index       string
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
}

type CustomSubSchema struct {
	Index       string
	Title       string `json:"title"`
	Type        string `json:"type"`
	Description string `json:"description,omitempty"`
	Format      string `json:"format,omitempty"`
	Required    bool   `json:"required,omitempty"`
	Mutability  string `json:"mutablity,omitempty"`
	Scope       string `json:"scope,omitempty"`
	MinLength   int    `json:"minLength,omitempty"`
	MaxLength   int    `json:"maxLength,omitempty"`
	Items       struct {
		Type string `json:"type,omitempty"`
	} `json:"items,omitempty"`
	Union       string        `json:"union,omitempty"`
	Enum        []string      `json:"enum,omitempty"`
	OneOf       []OneOf       `json:"oneOf,omitempty"`
	Permissions []Permissions `json:"permissions"`
	Master      struct {
		Type string `json:"type"`
	} `json:"master"`
}

type Permissions struct {
	Principal string `json:"principal"`
	Action    string `json:"action"`
}

type OneOf struct {
	Const string `json:"const"`
	Title string `json:"title"`
}

func (s *SchemasService) GetRawUserSchema() (map[string]interface{}, *Response, error) {
	u := fmt.Sprintf("meta/schemas/user/default")
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var temp map[string]interface{}
	resp, err := s.client.Do(req, &temp)
	if err != nil {
		return nil, resp, err
	}
	return temp, resp, err
}

func (s *SchemasService) GetUserSchema() (*Schema, *Response, error) {
	temp, resp, err := s.client.Schemas.GetRawUserSchema()
	if err != nil {
		return nil, resp, err
	}

	layout := "2006-01-02T15:04:05.000Z"
	create, err := time.Parse(layout, temp["created"].(string))
	if err != nil {
		return nil, resp, err
	}
	update, _ := time.Parse(layout, temp["lastUpdated"].(string))
	if err != nil {
		return nil, resp, err
	}

	schema := new(Schema)
	schema.ID = temp["id"].(string)
	schema.Schema = temp["$schema"].(string)
	schema.Name = temp["name"].(string)
	schema.Title = temp["title"].(string)
	schema.Created = create
	schema.LastUpdated = update
	schema.Type = temp["type"].(string)
	for k, v := range temp["definitions"].(map[string]interface{}) {
		switch k {
		case "base":
			for j, c := range v.(map[string]interface{}) {
				switch j {
				case "id":
					schema.Definitions.Base.ID = c.(string)
				case "type":
					schema.Definitions.Base.Type = c.(string)
				case "required":
					var req []string
					for _, j := range c.([]interface{}) {
						req = append(req, j.(string))
					}
					schema.Definitions.Base.Required = req
				case "properties":
					for g, _ := range c.(map[string]interface{}) {
						sub, resp, err := s.client.Schemas.GetUserBaseSubSchema(g)
						if err != nil {
							return nil, resp, err
						}
						schema.Definitions.Base.Properties = append(schema.Definitions.Base.Properties, *sub)
					}
				}
			}
		case "custom":
			for j, c := range v.(map[string]interface{}) {
				switch j {
				case "id":
					schema.Definitions.Custom.ID = c.(string)
				case "type":
					schema.Definitions.Custom.Type = c.(string)
				case "required":
					var req []string
					for _, j := range c.([]interface{}) {
						req = append(req, j.(string))
					}
					schema.Definitions.Custom.Required = req
				case "properties":
					for g, _ := range c.(map[string]interface{}) {
						sub, resp, err := s.client.Schemas.GetUserCustomSubSchema(g)
						if err != nil {
							return nil, resp, err
						}
						schema.Definitions.Custom.Properties = append(schema.Definitions.Custom.Properties, *sub)
					}
				}
			}
		}
	}
	return schema, resp, err
}

func (s *SchemasService) GetPropertiesMap(scope string) (map[string]interface{}, *Response, error) {
	if scope != "base" && scope != "custom" {
		return nil, nil, fmt.Errorf("[ERROR] SubSchema Properties Map input string var supports values \"base\" or \"custom\"")
	}
	temp, resp, err := s.client.Schemas.GetRawUserSchema()
	if err != nil {
		return nil, resp, err
	}
	for k, v := range temp["definitions"].(map[string]interface{}) {
		if k == scope {
			for j, c := range v.(map[string]interface{}) {
				if j == "properties" {
					return c.(map[string]interface{}) , nil, nil
				}
			}
		}
	}
	return nil, nil, nil
}

func (s *SchemasService) GetUserSchemaIndex(scope string) ([]string, *Response, error) {
	var index []string
	prop, resp, err := s.client.Schemas.GetPropertiesMap(scope)
	if err != nil {
		return nil, resp, err
	}
	for key := range prop {
		index = append(index, key)
	}
	return index, resp, err
}

func (s *SchemasService) GetUserBaseSubSchema(title string) (*BaseSubSchema, *Response, error) {
	exists := false
	schema, resp, err := s.client.Schemas.GetPropertiesMap("base")
	if err != nil {
		return nil, resp, err
	}
	index, resp, err := s.client.Schemas.GetUserSchemaIndex("base")
	if err != nil {
		return nil, resp, err
	}
	for _, field := range index {
		if title == field {
			exists = true
			break
		}
	}
	if exists == false {
		return nil, nil, fmt.Errorf("[ERROR] GetUserBaseSubSchema subschema %v does not exist in Okta", title)
	}
	subSchema := new(BaseSubSchema)
	subSchema.Index = title
	for k, v := range schema[title].(map[string]interface{}) {
		switch k {
		case "title":
			subSchema.Title = v.(string)
		case "type":
			subSchema.Type = v.(string)
		case "format":
			subSchema.Format = v.(string)
		case "required":
			subSchema.Required = v.(bool)
		case "mutability":
			subSchema.Mutability = v.(string)
		case "scope":
			subSchema.Scope = v.(string)
		case "minLength":
			subSchema.MinLength = int(v.(float64))
		case "maxLength":
			subSchema.MaxLength = int(v.(float64))
		case "permissions":
			perms := make([]Permissions, len(v.([]interface{})))
			for c, j := range v.([]interface{}) {
				for h, x := range j.(map[string]interface{}) {
					switch h {
					case "principal":
						perms[c].Principal = x.(string)
					case "action":
						perms[c].Action = x.(string)
					}
				}
			}
			subSchema.Permissions = perms
		case "master":
			for g, z := range v.(map[string]interface{}) {
				switch g {
				case "type":
					subSchema.Master.Type = z.(string)
				}
			}
		}
	}
	return subSchema, resp, err
}

func (s *SchemasService) GetUserCustomSubSchema(title string) (*CustomSubSchema, *Response, error) {
	exists := false
	schema, resp, err := s.client.Schemas.GetPropertiesMap("custom")
	if err != nil {
		return nil, resp, err
	}
	index, resp, err := s.client.Schemas.GetUserSchemaIndex("custom")
	if err != nil {
		return nil, resp, err
	}
	for _, field := range index {
		if title == field {
			exists = true
			break
		}
	}
	if exists == false {
		return nil, nil, fmt.Errorf("[ERROR] GetUserCustomSubSchema subschema %v does not exist in Okta", title)
	}
	subSchema := new(CustomSubSchema)
	subSchema.Index = title
	for k, v := range schema[title].(map[string]interface{}) {
		switch k {
		case "title":
			subSchema.Title = v.(string)
		case "type":
			subSchema.Type = v.(string)
		case "description":
			subSchema.Description = v.(string)
		case "format":
			subSchema.Format = v.(string)
		case "required":
			subSchema.Required = v.(bool)
		case "mutability":
			subSchema.Mutability = v.(string)
		case "scope":
			subSchema.Scope = v.(string)
		case "minLength":
			subSchema.MinLength = int(v.(float64))
		case "maxLength":
			subSchema.MaxLength = int(v.(float64))
		case "items":
			for g, z := range v.(map[string]interface{}) {
				switch g {
				case "type":
					subSchema.Items.Type = z.(string)
				}
			}
		case "union":
			subSchema.Union = v.(string)
		case "enum":
			// assuming here all enum values are strings, I hope i'm right
			subSchema.Enum = v.([]string)
		case "oneOf":
			oneof := make([]OneOf, len(v.([]interface{})))
			for c, j := range v.([]interface{}) {
				for h, x := range j.(map[string]interface{}) {
					switch h {
					case "const":
						oneof[c].Const = x.(string)
					case "title":
						oneof[c].Title = x.(string)
					}
				}
			}
			subSchema.OneOf = oneof
		case "permissions":
			perms := make([]Permissions, len(v.([]interface{})))
			for c, j := range v.([]interface{}) {
				for h, x := range j.(map[string]interface{}) {
					switch h {
					case "principal":
						perms[c].Principal = x.(string)
					case "action":
						perms[c].Action = x.(string)
					}
				}
			}
			subSchema.Permissions = perms
		case "master":
			for g, z := range v.(map[string]interface{}) {
				switch g {
				case "type":
					subSchema.Master.Type = z.(string)
				}
			}
		}
	}
	return subSchema, resp, err
}
