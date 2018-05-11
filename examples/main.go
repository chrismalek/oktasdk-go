package main

import (
	"flag"
	"fmt"
	"reflect"
	"runtime"
	"strings"

	"github.com/intello-io/oktasdk-go/okta"
)

var orgName = ""
var apiToken = ""
var ex = ""
var args = "00u5wb3ybyXqakBDa0h7"
var isProductionOKTAORG = false

type fn func()

func main() {
	flag.StringVar(&ex, "example", "all", "strng of the example we want to run")
	flag.StringVar(&orgName, "org_name", "", "org name to submit a request against. The logs command expects a url")
	flag.StringVar(&apiToken, "api_token", "", "the api token to submit a reuqest against")
	flag.Parse()
	fns := []fn{
		nameSearchExample,
		getActiveUsersExampleOnePageAtATime,
		getActiveUsersExampleAllPages,
		getActiveUserUpdatedInLastMonthAllPages,
		getFirstActiveUserRoles,
		createUserNoPassword,
		createUserWithPassword,
		createUserWithRecoveryAndPassword,
		CreateUserThenActivate,
		SetUserPassword,
		deactivateUser,
		// NOTE: we need to handle passing args here
		// getUser,
		// getUserMFAFactor,
		searchForGroupByName,
		getFirst3PageOfOKTAGroupsUpdatedRecently,
		getGroupByID,
		getRandomOKTAGroupUser,
		groupAddAndDelete,
		logListEventType,
		logListQuery,
		logSince,
	}
	for _, fn := range fns {
		fnName := strings.Replace(runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name(), "main.", "", 1)
		if ex == fnName || ex == "all" {
			fmt.Printf("Running %s\n", fnName)
			fn()
		}

	}
}

func printLogArray(logs []okta.Log) {
	for _, l := range logs {
		printLog(l)
	}
}

func printLog(log okta.Log) {
	fmt.Printf("\t LogID: %s. Event type: %s. Display Message: %s. Published At: %s\n", log.UUID, log.EventType, log.DisplayMessage, log.Published)

	var dataToLog []string
	// log the actor
	dataToLog = append(dataToLog, fmt.Sprintf("Actor: %s", log.Actor.String()))
	for _, t := range log.Targets {
		dataToLog = append(dataToLog, fmt.Sprintf("Target: %s", t.String()))
	}

	for _, l := range dataToLog {
		fmt.Printf("\t\t %s\n", l)
	}
}

func printUserArray(users []okta.User) {
	for _, user := range users {
		printUser(user)

	}
}

func printUser(user okta.User) {
	fmt.Printf("\t User: %v \tid: %v\n", user.Profile.Login, user.ID)
	fmt.Printf("\t\t Status: %v - Last Login: %v\n", user.Status, user.LastLogin)

	if user.MFAFactors != nil {
		fmt.Printf("\t\t--- MFA Status ---\n")

		for _, mfa := range user.MFAFactors {
			fmt.Printf("\t\tMFA Factor - Status: %v - Provider: %v - Type: %v\n", mfa.Status, mfa.Provider, mfa.FactorType)
		}
	}
	if user.Groups != nil {
		fmt.Printf("\t\t--- Group Memberships ---\n")
		for _, group := range user.Groups {
			fmt.Printf("\t\tGroup - Name: %v - ID: %v - Last Enrolled: %v\n", group.Profile.Name, group.ID, group.LastMembershipUpdated)
		}
	}

}
func printGroupArray(groups []okta.Group) {
	for _, group := range groups {
		printGroup(group)
	}

}

func printGroup(group okta.Group) {
	fmt.Printf("\tFound Group: ID: %v, Name: %v\n", group.ID, group.Profile.Name)
}

func printStart(fName string) string {
	fmt.Printf("** Begin %v ***\n", fName)
	return fName
}
func printEnd(fName string) {
	fmt.Printf("** End %v ***\n\n\n", fName)

}
