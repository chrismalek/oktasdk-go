package okta

import (
	//	"encoding/json"
	//	"fmt"
	//	"net/http"
	//	"reflect"
	//	"testing"
	"time"
)

var testSchema *Schema
var testBaseSubSchema *BaseSubSchema
var testCustomSubSchema *CustomSubSchema
var testPermissions *Permissions
var testOneOf *OneOf

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
	testCustomSubSchema.Permissions = append(testBaseSubSchema.Permissions, *testPermissions)

	testSchema = &Schema{
		ID:          "https://dev-XXXX.oktapreview.com/meta/schemas/user/default",
		Schema:      "http://json-schema.org/draft-04/schema#",
		Name:        "user",
		Title:       "User",
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

//func TestUserGet(t *testing.T) {
//
//	setup()
//	defer teardown()
//	setupTestUsers()
//
//	temp, err := json.Marshal(testuser)
//	if err != nil {
//		t.Errorf("Users.Get json Marshall returned error: %v", err)
//	}
//	userTestJSONString := string(temp)
//
//	mux.HandleFunc("/users/00ub0oNGTSWTBKOLGLNR", func(w http.ResponseWriter, r *http.Request) {
//		testMethod(t, r, "GET")
//		testAuthHeader(t, r)
//		fmt.Fprint(w, userTestJSONString)
//	})
//
//	user, _, err := client.Users.GetByID("00ub0oNGTSWTBKOLGLNR")
//	if err != nil {
//		t.Errorf("Users.Get returned error: %v", err)
//	}
//	if !reflect.DeepEqual(user, testuser) {
//		// fmt.Printf("pretty---\n%v\n---\n", pretty.Diff(user, testuser))
//		t.Errorf("client.Users.GetByID returned \n\t%+v, want \n\t%+v\n", user, testuser)
//	}
//}
