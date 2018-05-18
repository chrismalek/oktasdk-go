package okta

import (
	"encoding/json"
	"fmt"
	//"github.com/davecgh/go-spew/spew"
	"net/http"
	"reflect"
	"testing"
	"time"
)

var testSchema *Schema
var testBaseSubSchema *BaseSubSchema
var testCustomSubSchema *CustomSubSchema
var testPermissions *Permissions
var testOneOf *OneOf

var schemaTestJSONString = `
{
    "id": "https://dev-XXXX.oktapreview.com/meta/schemas/user/default",
    "$schema": "http://json-schema.org/draft-04/schema#",
    "name": "user",
    "title": "Default Okta User",
    "type": "object",
    "lastUpdated": "2018-02-16T19:59:05.000Z",
    "created": "2018-02-16T19:59:05.000Z",
    "definitions": {
        "base": {
            "id": "#base",
            "type": "object",
            "properties": {
                "firstName": {
                    "title": "first name",
                    "type": "string",
                    "required": true,
                    "format": "firstname",
                    "mutability": "READ_WRITE",
                    "scope": "NONE",
                    "minLength": 1,
                    "maxLength": 50,
                    "master":
                        {
                            "type": "PROFILE_MASTER"
                        },
                    "permissions": [
                        {
                            "principal": "SELF",
                            "action": "READ_WRITE"
                        }
                    ]
                }
            },
            "required": [
                "login",
                "firstName",
                "lastName",
                "email"
            ]
        },
        "custom": {
            "id": "#custom",
            "type": "object",
            "properties": {
                "testSubSchema": {
                    "title": "test subschema",
                    "type": "array",
                    "description": "test subschema",
                    "required": false,
                    "mutability": "READ_WRITE",
                    "scope": "NONE",
                    "union": "DISABLE",
                    "enum": ["S", "M", "L", "XL"],
                    "items":
                        {
                            "type": "string"
                        },
                    "master":
                        {
                            "type": "PROFILE_MASTER"
                        },
                    "oneOf": [
                        {
                            "const": "S",
                            "title": "Small"
                        }
                    ],
                    "permissions": [
                        {
                            "principal": "SELF",
                            "action": "READ_WRITE"
                        }
                    ]
                }
            },
            "required": []
        }
    },
    "properties": {
        "profile": {
            "allOf": [
                {
                    "$ref": "#/definitions/base"
                },
                {
                    "$ref": "#/definitions/custom"
                }
            ]
        }
    }
}
`

func setupTestSchemas() {

	hmm, _ := time.Parse("2006-01-02T15:04:05.000Z", "2018-02-16T19:59:05.000Z")

	testPermissions = &Permissions{
		Principal: "SELF",
		Action:    "READ_WRITE",
	}

	testOneOf = &OneOf{
		Const: "S",
		Title: "Small",
	}

	testBaseSubSchema = &BaseSubSchema{
		Index:      "firstName",
		Title:      "First name",
		Type:       "string",
		Format:     "firstname",
		Required:   true,
		Mutability: "READ_WRITE",
		Scope:      "NONE",
		MinLength:  1,
		MaxLength:  50,
	}
	testBaseSubSchema.Master.Type = "PROFILE_MASTER"
	testBaseSubSchema.Permissions = append(testBaseSubSchema.Permissions, *testPermissions)

	testCustomSubSchema = &CustomSubSchema{
		Index:       "testSubSchema",
		Title:       "test subschema",
		Type:        "array",
		Description: "test subschema",
		Format:      "",
		Required:    false,
		Mutability:  "READ_WRITE",
		Scope:       "NONE",
		MinLength:   0,
		MaxLength:   0,
		Union:       "DISABLE",
		Enum:        []string{"S", "M", "L", "XL"},
	}
	testCustomSubSchema.Items.Type = "string"
	testCustomSubSchema.Master.Type = "PROFILE_MASTER"
	testCustomSubSchema.OneOf = append(testCustomSubSchema.OneOf, *testOneOf)
	testCustomSubSchema.Permissions = append(testCustomSubSchema.Permissions, *testPermissions)

	testSchema = &Schema{
		ID:          "https://dev-XXXX.oktapreview.com/meta/schemas/user/default",
		Schema:      "http://json-schema.org/draft-04/schema#",
		Name:        "user",
		Title:       "Default Okta User",
		Created:     hmm,
		LastUpdated: hmm,
		Type:        "object",
	}
	testSchema.Definitions.Base.ID = "#base"
	testSchema.Definitions.Base.Type = "object"
	testSchema.Definitions.Base.Properties = append(testSchema.Definitions.Base.Properties, *testBaseSubSchema)
	testSchema.Definitions.Base.Required = []string{"login", "firstName", "lastName", "email"}
	testSchema.Definitions.Custom.ID = "#custom"
	testSchema.Definitions.Custom.Type = "object"
	testSchema.Definitions.Custom.Properties = append(testSchema.Definitions.Custom.Properties, *testCustomSubSchema)
	testSchema.Definitions.Custom.Required = []string{}
}

func TestSchemaGet(t *testing.T) {

	setup()
	defer teardown()
	setupTestSchemas()

	var outMap map[string]interface{}
	json.Unmarshal([]byte(schemaTestJSONString), &outMap)

	mux.HandleFunc("/meta/schemas/user/default", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testAuthHeader(t, r)
		fmt.Fprint(w, schemaTestJSONString)
	})

	schema, _, err := client.Schemas.GetRawUserSchema()
	if err != nil {
		t.Errorf("SchemaRaw.Get returned error: %v", err)
	}
	if !reflect.DeepEqual(schema, outMap) {
		t.Errorf("client.Schemas.GetRawUserSchema returned \n\t%+v, want \n\t%+v\n", schema, outMap)
	}
	schemaStruct, err := client.Schemas.userSchema(schema)
	if err != nil {
		t.Errorf("UserSchema.Get returned error: %v", err)
	}
	if !reflect.DeepEqual(schemaStruct, testSchema) {
		t.Errorf("client.Schemas.userSchema returned \n\t%+v, want \n\t%+v\n", schemaStruct, testSchema)
	}
}
