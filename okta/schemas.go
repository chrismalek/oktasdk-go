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
			ID         string
			Type       string
			Properties []BaseSubSchema
			Required   []string
		}
		Custom struct {
			ID         string
			Type       string
			Properties []CustomSubSchema
			Required   []string
		}
	}
	Type string
}

type BaseSubSchema struct {
	Index       string        // json do not export
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
	Index       string // json do not export
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
			for k2, v2 := range v.(map[string]interface{}) {
				switch k2 {
				case "id":
					schema.Definitions.Base.ID = v2.(string)
				case "type":
					schema.Definitions.Base.Type = v2.(string)
				case "required":
					var req []string
					for _, v3 := range v2.([]interface{}) {
						req = append(req, v3.(string))
					}
					schema.Definitions.Base.Required = req
				case "properties":
					for k3, v3 := range v2.(map[string]interface{}) {
						sub, err := s.client.Schemas.GetUserBaseSubSchema(k3, v3.(map[string]interface{}))
						if err != nil {
							return nil, resp, err
						}
						schema.Definitions.Base.Properties = append(schema.Definitions.Base.Properties, *sub)
					}
				}
			}
		case "custom":
			for k2, v2 := range v.(map[string]interface{}) {
				switch k2 {
				case "id":
					schema.Definitions.Custom.ID = v2.(string)
				case "type":
					schema.Definitions.Custom.Type = v2.(string)
				case "required":
					var req []string
					for _, v3 := range v2.([]interface{}) {
						req = append(req, v3.(string))
					}
					schema.Definitions.Custom.Required = req
				case "properties":
					for k3, v3 := range v2.(map[string]interface{}) {
						sub, err := s.client.Schemas.GetUserCustomSubSchema(k3, v3.(map[string]interface{}))
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

func (s *SchemasService) userSubSchemaPropMap(scope string) (map[string]interface{}, *Response, error) {
	if scope != "base" && scope != "custom" {
		return nil, nil, fmt.Errorf("[ERROR] SubSchema Properties Map input string var supports values \"base\" or \"custom\"")
	}
	temp, resp, err := s.client.Schemas.GetRawUserSchema()
	if err != nil {
		return nil, resp, err
	}
	for k, v := range temp["definitions"].(map[string]interface{}) {
		if k == scope {
			for k2, v2 := range v.(map[string]interface{}) {
				if k2 == "properties" {
					return v2.(map[string]interface{}), nil, nil
				}
			}
		}
	}
	return nil, nil, nil
}

func (s *SchemasService) GetUserSubSchemaPropMap(scope string, title string) (map[string]interface{}, *Response, error) {
	prop, resp, err := s.client.Schemas.userSubSchemaPropMap(scope)
	if err != nil {
		return nil, resp, err
	}
	if v, ok := prop[title]; ok {
		return v.(map[string]interface{}), resp, err
	} else {
		return nil, resp, fmt.Errorf("[ERROR] GetUserSubSchemaPropMap subschema %v not found", title)
	}

	return nil, resp, err
}

func (s *SchemasService) GetUserSubSchemaIndex(scope string) ([]string, *Response, error) {
	var index []string
	prop, resp, err := s.client.Schemas.userSubSchemaPropMap(scope)
	if err != nil {
		return nil, resp, err
	}
	for key := range prop {
		index = append(index, key)
	}
	return index, resp, err
}

func (s *SchemasService) GetUserBaseSubSchema(title string, obj map[string]interface{}) (*BaseSubSchema, error) {
	subSchema := new(BaseSubSchema)
	subSchema.Index = title
	if v, ok := obj["title"]; ok {
		subSchema.Title = v.(string)
	} else {
		// if we cant find a title field, we'll assume this obj map is not correct
		return nil, fmt.Errorf("[ERROR] GetUserBaseSubSchema interface map parsing error")
	}
	if v, ok := obj["type"]; ok {
		subSchema.Type = v.(string)
	}
	if v, ok := obj["format"]; ok {
		subSchema.Format = v.(string)
	}
	if v, ok := obj["required"]; ok {
		subSchema.Required = v.(bool)
	}
	if v, ok := obj["mutability"]; ok {
		subSchema.Mutability = v.(string)
	}
	if v, ok := obj["scope"]; ok {
		subSchema.Scope = v.(string)
	}
	if v, ok := obj["minLength"]; ok {
		subSchema.MinLength = int(v.(float64))
	}
	if v, ok := obj["maxLength"]; ok {
		subSchema.MaxLength = int(v.(float64))
	}
	if v, ok := obj["permissions"]; ok {
		perms := make([]Permissions, len(v.([]interface{})))
		for k2, v2 := range v.([]interface{}) {
			for k3, v3 := range v2.(map[string]interface{}) {
				switch k3 {
				case "principal":
					perms[k2].Principal = v3.(string)
				case "action":
					perms[k2].Action = v3.(string)
				}
			}
		}
		subSchema.Permissions = perms
	}
	if v, ok := obj["master"]; ok {
		for k2, v2 := range v.(map[string]interface{}) {
			switch k2 {
			case "type":
				subSchema.Master.Type = v2.(string)
			}
		}
	}
	return subSchema, nil
}

func (s *SchemasService) GetUserCustomSubSchema(title string, obj map[string]interface{}) (*CustomSubSchema, error) {
	subSchema := new(CustomSubSchema)
	subSchema.Index = title
	if v, ok := obj["title"]; ok {
		subSchema.Title = v.(string)
	} else {
		// if we cant find a title field, we'll assume this obj map is not correct
		return nil, fmt.Errorf("[ERROR] GetUserCustomSubSchema interface map parsing error")
	}
	if v, ok := obj["type"]; ok {
		subSchema.Type = v.(string)
	}
	if v, ok := obj["description"]; ok {
		subSchema.Description = v.(string)
	}
	if v, ok := obj["format"]; ok {
		subSchema.Format = v.(string)
	}
	if v, ok := obj["required"]; ok {
		subSchema.Required = v.(bool)
	}
	if v, ok := obj["mutability"]; ok {
		subSchema.Mutability = v.(string)
	}
	if v, ok := obj["scope"]; ok {
		subSchema.Scope = v.(string)
	}
	if v, ok := obj["minLength"]; ok {
		subSchema.MinLength = int(v.(float64))
	}
	if v, ok := obj["maxLength"]; ok {
		subSchema.MaxLength = int(v.(float64))
	}
	if v, ok := obj["items"]; ok {
		for k2, v2 := range v.(map[string]interface{}) {
			switch k2 {
			case "type":
				subSchema.Items.Type = v2.(string)
			}
		}
	}
	if v, ok := obj["union"]; ok {
		subSchema.Union = v.(string)
	}
	if v, ok := obj["enum"]; ok {
		// assuming here all enum values are strings, I hope i'm right
		subSchema.Enum = v.([]string)
	}
	if v, ok := obj["oneOf"]; ok {
		oneof := make([]OneOf, len(v.([]interface{})))
		for k2, v2 := range v.([]interface{}) {
			for k3, v3 := range v2.(map[string]interface{}) {
				switch k3 {
				case "const":
					oneof[k2].Const = v3.(string)
				case "title":
					oneof[k2].Title = v3.(string)
				}
			}
		}
		subSchema.OneOf = oneof
	}
	if v, ok := obj["permissions"]; ok {
		perms := make([]Permissions, len(v.([]interface{})))
		for k2, v2 := range v.([]interface{}) {
			for k3, v3 := range v2.(map[string]interface{}) {
				switch k3 {
				case "principal":
					perms[k2].Principal = v3.(string)
				case "action":
					perms[k2].Action = v3.(string)
				}
			}
		}
		subSchema.Permissions = perms
	}
	if v, ok := obj["master"]; ok {
		for k2, v2 := range v.(map[string]interface{}) {
			switch k2 {
			case "type":
				subSchema.Master.Type = v2.(string)
			}
		}
	}
	return subSchema, nil
}
