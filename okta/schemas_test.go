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
                    "title": "First name",
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
	// reflect.DeepEqual results in false even when the structs are identical
	// possibly because of the fields that contain time values https://stackoverflow.com/a/45222521
	// suggested fix may be to add https://github.com/google/go-cmp package
	orig := testSchema.Definitions.Base.Properties[0].Index
	final := schemaStruct.Definitions.Base.Properties[0].Index
	if !reflect.DeepEqual(final, orig) {
		t.Errorf("client.Schemas.userSchema returned \n\t%+v, want \n\t%+v\n", final, orig)
	}
	orig = testSchema.Definitions.Custom.Properties[0].Index
	final = schemaStruct.Definitions.Custom.Properties[0].Index
	if !reflect.DeepEqual(final, orig) {
		t.Errorf("client.Schemas.userSchema returned \n\t%+v, want \n\t%+v\n", final, orig)
	}
}

func TestSubSchemaIndexGet(t *testing.T) {

	setup()
	defer teardown()
	setupTestSchemas()

	mux.HandleFunc("/meta/schemas/user/default", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testAuthHeader(t, r)
		fmt.Fprint(w, schemaTestJSONString)
	})

	index, _, err := client.Schemas.GetUserSubSchemaIndex("custom")
	if err != nil {
		t.Errorf("SchemaIndex.Get returned error: %v", err)
	}
	exists := false
	for _, v := range index {
		if v == testCustomSubSchema.Index {
			exists = true
		}
	}
	if exists == false {
		t.Errorf("SchemaIndex.Get did not return valid data in slice: %+v", index)
	}
}

func TestSubSchemaCustomGet(t *testing.T) {

	setup()
	defer teardown()
	setupTestSchemas()

	mux.HandleFunc("/meta/schemas/user/default", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testAuthHeader(t, r)
		fmt.Fprint(w, schemaTestJSONString)
	})

	propMap, _, err := client.Schemas.GetUserSubSchemaPropMap("custom", "testSubSchema")
	if err != nil {
		t.Errorf("SchemaPropMap.Get returned error: %v", err)
	}
	if _, ok := propMap["title"]; !ok {
		t.Errorf("SchemaPropMap.Get did not return a valid map: %+v", propMap)
	} else if propMap["title"] != "test subschema" {
		t.Errorf("SchemaPropMap.Get did not return valid data in map: %+v", propMap)
	}

	custom, err := client.Schemas.GetUserCustomSubSchema("testSubSchema", propMap)
	if err != nil {
		t.Errorf("SchemaCustomSubSchema.Get returned error: %v", err)
	}
	if !reflect.DeepEqual(custom, testCustomSubSchema) {
		t.Errorf("client.Schemas.GetUserCustomSubSchema returned \n\t%+v, want \n\t%+v\n", custom, testCustomSubSchema)
	}
}

func TestSubSchemaBaseGet(t *testing.T) {

	setup()
	defer teardown()
	setupTestSchemas()

	mux.HandleFunc("/meta/schemas/user/default", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testAuthHeader(t, r)
		fmt.Fprint(w, schemaTestJSONString)
	})

	propMap, _, err := client.Schemas.GetUserSubSchemaPropMap("base", "firstName")
	if err != nil {
		t.Errorf("SchemaPropMap.Get returned error: %v", err)
	}
	if _, ok := propMap["title"]; !ok {
		t.Errorf("SchemaPropMap.Get did not return a valid map: %+v", propMap)
	} else if propMap["title"] != "First name" {
		t.Errorf("SchemaPropMap.Get did not return valid data in map: %+v", propMap)
	}

	custom, err := client.Schemas.GetUserBaseSubSchema("firstName", propMap)
	if err != nil {
		t.Errorf("SchemaCustomSubSchema.Get returned error: %v", err)
	}
	if !reflect.DeepEqual(custom, testBaseSubSchema) {
		t.Errorf("client.Schemas.GetUserBaseSubSchema returned \n\t%+v, want \n\t%+v\n", custom, testBaseSubSchema)
	}
}

func TestUpdateCustomSubSchema(t *testing.T) {

	setup()
	defer teardown()
	setupTestSchemas()

	mux.HandleFunc("/meta/schemas/user/default", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testAuthHeader(t, r)
		fmt.Fprint(w, schemaTestJSONString)
	})

	schema, _, err := client.Schemas.UpdateUserCustomSubSchema(*testCustomSubSchema)
	if err != nil {
		t.Errorf("CustomSubSchema.Update returned error: %v", err)
	}
	// reflect.DeepEqual results in false even when the structs are identical
	// possibly because of the fields that contain time values https://stackoverflow.com/a/45222521
	// suggested fix may be to add https://github.com/google/go-cmp package
	orig := testSchema.Definitions.Custom.Properties[0].Index
	final := schema.Definitions.Custom.Properties[0].Index
	if !reflect.DeepEqual(final, orig) {
		t.Errorf("client.Schemas.GetUserBaseSubSchema returned \n\t%+v, want \n\t%+v\n", final, orig)
	}
}

func TestUpdateBaseSubSchema(t *testing.T) {

	setup()
	defer teardown()
	setupTestSchemas()

	mux.HandleFunc("/meta/schemas/user/default", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testAuthHeader(t, r)
		fmt.Fprint(w, schemaTestJSONString)
	})

	schema, _, err := client.Schemas.UpdateUserBaseSubSchema(*testBaseSubSchema)
	if err != nil {
		t.Errorf("BaseSubSchema.Update returned error: %v", err)
	}
	// reflect.DeepEqual results in false even when the structs are identical
	// possibly because of the fields that contiain time values https://stackoverflow.com/a/45222521
	// suggested fix may be to add https://github.com/google/go-cmp package
	orig := testSchema.Definitions.Base.Properties[0].Index
	final := schema.Definitions.Base.Properties[0].Index
	if !reflect.DeepEqual(final, orig) {
		t.Errorf("client.Schemas.UpdateUserCustomSubSchema returned \n\t%+v, want \n\t%+v\n", final, orig)
	}
}

func TestDeleteCustomSubSchema(t *testing.T) {

	setup()
	defer teardown()
	setupTestSchemas()

	mux.HandleFunc("/meta/schemas/user/default", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testAuthHeader(t, r)
		fmt.Fprint(w, schemaTestJSONString)
	})

	schema, _, err := client.Schemas.DeleteUserCustomSubSchema("testSubSchema")
	if err != nil {
		t.Errorf("CustomSubSchema.Delete returned error: %v", err)
	}
	// reflect.DeepEqual results in false even when the structs are identical
	// possibly because of the fields that contain time values https://stackoverflow.com/a/45222521
	// suggested fix may be to add https://github.com/google/go-cmp package
	orig := testSchema.Definitions.Custom.Properties[0].Index
	final := schema.Definitions.Custom.Properties[0].Index
	if !reflect.DeepEqual(final, orig) {
		t.Errorf("client.Schemas.GetUserBaseSubSchema returned \n\t%+v, want \n\t%+v\n", final, orig)
	}
}
