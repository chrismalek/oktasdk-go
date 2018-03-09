package main

import (
	"fmt"
	"os"

	"github.com/nzoschke/oktasdk-go/okta"
)

var orgName = os.Getenv("OKTA_API_TEST_ORG")
var apiToken = os.Getenv("OKTA_API_TEST_TOKEN")
var isProductionOKTAORG = false

func main() {
	fmt.Printf("\n\n%%%%%% User Examples\n\n\n")
	nameSearchExample()
	getActiveUsersExampleOnePageAtATime()
	getActiveUsersExampleAllPages()
	getActiveUserUpdatedInLastMonthAllPages()
	getFirstActiveUserRoles()
	createUserNoPassword()
	createUserWithPassword()
	createUserWithRecoveryAndPassword()
	CreateUserThenActivate()
	SetUserPassword()
	deactivateUser()
	getUser("00u5wb3ybyXqakBDa0h7")
	getUserMFAFactor("00u5wb3ybyXqakBDa0h7")

	fmt.Printf("\n\n%%%%%% Group Examples\n\n\n")

	searchForGroupByName()
	getFirst3PageOfOKTAGroupsUpdatedRecently()
	getGroupByID()
	getRandomOKTAGroupUser()
	groupAddAndDelete()

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
