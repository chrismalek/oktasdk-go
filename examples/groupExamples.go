package main

import (
	"fmt"
	"time"

	"github.com/chrismalek/oktasdk-go/okta"
)

func searchForGroupByName() {
	defer printEnd(printStart("SearchForGroupByName"))

	client := okta.NewClient(nil, orgName, apiToken, false)

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

	client := okta.NewClient(nil, orgName, apiToken, false)
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
