package main

import (
	"fmt"
	"time"

	"github.com/chrismalek/oktasdk-go/okta"
)

func nameSearchExample() {
	defer printEnd(printStart("nameSearchExample"))

	client := okta.NewClient(nil, orgName, apiToken, isProductionOKTAORG)

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
	client := okta.NewClient(nil, orgName, apiToken, isProductionOKTAORG)
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

			fmt.Printf("\tPage return %d users\n", len(userPage))

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

	client := okta.NewClient(nil, orgName, apiToken, isProductionOKTAORG)

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

// In this example, we will get all active users who are updated in the last month
func getActiveUserUpdatedInLastMonthAllPages() {
	defer printEnd(printStart("getActiveUserUpdatedInLastMonthAllPages"))

	client := okta.NewClient(nil, orgName, apiToken, isProductionOKTAORG)

	fmt.Printf("Client Base URL: %v\n\n", client.BaseURL)

	userFilter := &okta.UserListFilterOptions{}
	userFilter.GetAllPages = true
	userFilter.StatusEqualTo = okta.UserStatusActive
	userFilter.LastUpdated.Value = time.Now().AddDate(0, -1, 0)
	userFilter.LastUpdated.Operator = okta.FilterGreaterThanOperator

	allUsers, response, err := client.Users.ListWithFilter(userFilter)

	if err != nil {
		fmt.Printf("Response Error %+v\n\t URL used:%v\n", err, response.Request.URL.String())
	}

	fmt.Printf("len(all_users) = %v\n", len(allUsers))
	printUserArray(allUsers)

}

func getFirstActiveUserRoles() {
	defer printEnd(printStart("getFirstActiveUserRoles"))

	client := okta.NewClient(nil, orgName, apiToken, isProductionOKTAORG)

	fmt.Printf("Client Base URL: %v\n\n", client.BaseURL)

	userFilter := &okta.UserListFilterOptions{}
	userFilter.Limit = 1
	userFilter.StatusEqualTo = okta.UserStatusActive

	randomActiveUsers, response, err := client.Users.ListWithFilter(userFilter)

	if err != nil {
		fmt.Printf("Response Error %+v\n\t URL used:%v\n", err, response.Request.URL.String())
	}

	fmt.Printf("len(all_users) = %v\n", len(randomActiveUsers))
	printUserArray(randomActiveUsers)
	if len(randomActiveUsers) > 0 {
		response, err := client.Users.PopulateGroups(&randomActiveUsers[0])
		if err != nil {
			fmt.Printf("Response Error %+v\n\t URL used:%v\n", err, response.Request.URL.String())
		}
		printGroupArray(randomActiveUsers[0].Groups)
	}

}
