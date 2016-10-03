package main

import (
	"fmt"
	"os"

	"github.com/chrismalek/oktasdk-go/okta"
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
