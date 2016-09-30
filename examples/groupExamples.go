package main

import (
	"fmt"
	"time"

	"github.com/chrismalek/oktasdk-go/okta"
)

func searchForGroupByName() {
	defer printEnd(printStart("SearchForGroupByName"))

	client := okta.NewClient(nil, orgName, apiToken, isProductionOKTAORG)

	fmt.Printf("Client Base URL: %v\n\n", client.BaseURL)

	groupFilterOptions := new(okta.GroupFilterOptions)
	groupFilterOptions.GetAllPages = true
	groupFilterOptions.NameStartsWith = "TEST"
	groupSearchResult, response, err := client.Groups.ListWithFilter(groupFilterOptions)

	if err != nil {
		fmt.Printf("Response Error %+v\n\t URL used:%v\n", err, response.Request.URL.String())
	} else {

		fmt.Printf("groupSearchResult Len : %v\n\tURL Used: %v\n", len(groupSearchResult), response.Request.URL)
		printGroupArray(groupSearchResult)
	}

}

func getFirst3PageOfOKTAGroupsUpdatedRecently() {
	defer printEnd(printStart("getFirst3PageOfOKTAGroupsUpdatedRecently"))

	client := okta.NewClient(nil, orgName, apiToken, isProductionOKTAORG)
	groupFilterOptions := new(okta.GroupFilterOptions)
	groupFilterOptions.GetAllPages = false
	groupFilterOptions.NumberOfPages = 3 // Limit to first 3 Pages
	groupFilterOptions.Limit = 5         // Override default Limit and pull 5 groups per page
	groupFilterOptions.GroupTypeEqual = okta.GroupTypeOKTA
	groupFilterOptions.LastUpdated.Value = time.Now().AddDate(0, -1, 0) // Last Month
	groupFilterOptions.LastUpdated.Operator = okta.FilterGreaterThanOperator

	groupSearchResult, response, err := client.Groups.ListWithFilter(groupFilterOptions)
	if err != nil {
		fmt.Printf("Response Error %+v\n\t URL used:%v\n", err, response.Request.URL.String())
	} else {

		fmt.Printf("groupSearchResult Len : %v\n\tURL Used: %v\n", len(groupSearchResult), response.Request.URL)
		printGroupArray(groupSearchResult)
	}
}

// Get a specific Group by the OKTA ID.
// This Example will not return a result in your org.
func getGroupByID() {
	defer printEnd(printStart("getGroupByID"))
	client := okta.NewClient(nil, orgName, apiToken, isProductionOKTAORG)

	group, response, err := client.Groups.GetByID("00g2nlq3l6DDWPIYBOTG")

	if err != nil {
		fmt.Printf("Response Error %+v\n\t URL used:%v\n", err, response.Request.URL.String())

	} else {
		printGroup(*group)
	}

}

func getRandomOKTAGroupUser() {
	defer printEnd(printStart("getRandomOKTAGroupUser"))

	client := okta.NewClient(nil, orgName, apiToken, isProductionOKTAORG)
	groupFilterOptions := new(okta.GroupFilterOptions)
	groupFilterOptions.GetAllPages = false
	groupFilterOptions.NumberOfPages = 1                   // only get 1 page - This is actuall the default behavior
	groupFilterOptions.Limit = 1                           // Override default Limit and pull 1 groups per page
	groupFilterOptions.GroupTypeEqual = okta.GroupTypeOKTA // OKTA mastered Group
	// Filter on groups that have had a membership update in the last month
	groupFilterOptions.LastMembershipUpdated.Value = time.Now().AddDate(0, -1, 0) // Last Month
	groupFilterOptions.LastMembershipUpdated.Operator = okta.FilterGreaterThanOperator

	groupSearchResult, response, err := client.Groups.ListWithFilter(groupFilterOptions)
	if err != nil {
		fmt.Printf("Response Error %+v\n\t URL used:%v\n", err, response.Request.URL.String())

	} else {
		if len(groupSearchResult) > 0 {
			printGroup(groupSearchResult[0])
			// Get All Users in Group - We passs a
			groupUserOptions := new(okta.GroupUserFilterOptions)
			groupUserOptions.GetAllPages = true

			userList, response, err := client.Groups.GetUsers(groupSearchResult[0].ID, groupUserOptions)
			if err != nil {
				fmt.Printf("Response Error %+v\n\t URL used:%v\n", err, response.Request.URL.String())

			} else {
				fmt.Printf("groupSearchResult Len : %v\n\tURL Used: %v\n", len(groupSearchResult), response.Request.URL)
				printUserArray(userList)
			}

		}

	}

}
func groupAddAndDelete() {
	defer printEnd(printStart("groupAdd"))
	client := okta.NewClient(nil, orgName, apiToken, isProductionOKTAORG)

	groupName := "OKTASDK-GO-" + time.Now().String()
	groupDescription := "Created by OKTASDK-GO @ " + time.Now().String()

	newGroup, response, err := client.Groups.Add(groupName, groupDescription)
	if err != nil {
		fmt.Printf("Response Error %+v\n\t URL used:%v\n", err, response.Request.URL.String())
		return
	}
	fmt.Printf("Created New Group\n")
	printGroup(*newGroup)

	fmt.Printf("Deleting New Group ID: %v\n", newGroup.ID)
	response, err = client.Groups.Delete(newGroup.ID)
	if err != nil {
		fmt.Printf("Response Error %+v\n\t URL used:%v\n", err, response.Request.URL.String())
		return
	}
}
