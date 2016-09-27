package main

import (
	"fmt"
	"os"

	"github.com/chrismalek/oktasdk-go/okta"
)

var orgName = os.Getenv("OKTA_API_TEST_ORG")
var apiToken = os.Getenv("OKTA_API_TEST_TOKEN")

func main() {
	fmt.Printf("\n\n%%%%%% User Examples\n\n\n")
	nameSearchExample()
	getActiveUsersExampleOnePageAtATime()
	getActiveUsersExampleAllPages()
	getActiveUserUpdatedInLastMonthAllPages()
	fmt.Printf("\n\n%%%%%% Group Examples\n\n\n")

	searchForGroupByName()
	getFirst3PageOfOKTAGroupsUpdatedRecently()
}

func printUserArray(users []okta.User) {
	for _, user := range users {
		fmt.Printf("Found User: %v\n", user.Profile.Login)

	}
}

func printGroupArray(groups []okta.Group) {
	for _, group := range groups {
		fmt.Printf("Found Group: %v\n", group.Profile.Name)
	}

}

func printStart(fName string) string {
	fmt.Printf("** Begin %v ***\n", fName)
	return fName
}
func printEnd(fName string) {
	fmt.Printf("** End %v ***\n\n\n", fName)

}
