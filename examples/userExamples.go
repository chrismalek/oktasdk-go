package main

import (
	"fmt"
	"os"

	"github.com/chrismalek/oktasdk-go/okta"
)

var orgName = os.Getenv("OKTA_API_TEST_ORG")
var apiToken = os.Getenv("OKTA_API_TEST_TOKEN")

func main() {

	nameSearchExample()
	getActiveUsersExampleOnePageAtATime()
	getActiveUsersExampleAllPages()
}

func printUserArray(users []okta.User) {
	for _, user := range users {
		fmt.Printf("Found User: %v\n", user.Profile.Login)

	}

}

func nameSearchExample() {
	defer printEnd(printStart("nameSearchExample"))

	client := okta.NewClient(nil, orgName, apiToken, false)

	fmt.Printf("Client Base URL: %v\n\n", client.BaseURL)

	fmt.Printf("****\nChris Malek Search\n")
	userFilter := &okta.UserListFilterOptions{}
	userFilter.FirstNameEqualTo = "Chris"
	userFilter.LastNameEqualTo = "Malek"
	userList, response, err := client.Users.ListWithFilter(userFilter)

	if err != nil {
		fmt.Printf("Response Error %+v\n\t URL used:%v\n", err, response.Request.URL.String())

	} else {
		printUserArray(userList)
	}
}

// In this example we will find all active users and our code will page through them
// One page at a time.
func getActiveUsersExampleOnePageAtATime() {
	defer printEnd(printStart("getActiveUsersExampleOnePageAtATime"))
	client := okta.NewClient(nil, orgName, apiToken, false)
	fmt.Printf("Client Base URL: %v\n\n", client.BaseURL)

	userFilter := &okta.UserListFilterOptions{}
	userFilter.Limit = 10 // You want 10 users per page
	userFilter.StatusEqualTo = okta.UserStatusActive

	var allUsers []okta.User

	i := 0
	for {

		i++
		userPage, response, err := client.Users.ListWithFilter(userFilter)

		if err != nil {
			fmt.Printf("Response Error %+v\n\t URL used:%v\n", err, response.Request.URL.String())
		} else {

			allUsers = append(allUsers, userPage...)

			fmt.Printf("Page return %d users\n", len(userPage))

			if response.NextURL == nil {
				break
			}
			userFilter.NextURL = response.NextURL
		}
	}
	fmt.Printf("len(all_users) = %v\n", len(allUsers))
	printUserArray(allUsers)

}

// In this example we will find all active users but we have the SDK do the paging and return all users.
func getActiveUsersExampleAllPages() {
	defer printEnd(printStart("getActiveUsersExampleAllPages"))

	client := okta.NewClient(nil, orgName, apiToken, false)

	fmt.Printf("Client Base URL: %v\n\n", client.BaseURL)

	userFilter := &okta.UserListFilterOptions{}
	userFilter.GetAllPages = true
	userFilter.StatusEqualTo = okta.UserStatusActive

	allUsers, response, err := client.Users.ListWithFilter(userFilter)

	if err != nil {
		fmt.Printf("Response Error %+v\n\t URL used:%v\n", err, response.Request.URL.String())
	}

	fmt.Printf("len(all_users) = %v\n", len(allUsers))
	printUserArray(allUsers)

}

func printStart(fName string) string {
	fmt.Printf("** Begin %v ***\n", fName)
	return fName
}
func printEnd(fName string) {
	fmt.Printf("** End %v ***\n", fName)

}
