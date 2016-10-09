package main

import (
	"encoding/json"
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

func createUserNoPassword() {
	defer printEnd(printStart("createUserNoPassword"))

	client := okta.NewClient(nil, orgName, apiToken, isProductionOKTAORG)

	fmt.Printf("Client Base URL: %v\n\n", client.BaseURL)

	newUserTemplate := client.Users.NewUser()
	newUserTemplate.Profile.FirstName = "Test SDK First"
	newUserTemplate.Profile.LastName = "Test SDK Last" + time.Now().Format("2006-01-02")
	newUserTemplate.Profile.Login = "testuser2@localhost.com"
	newUserTemplate.Profile.Email = newUserTemplate.Profile.Login

	jsonTest, _ := json.Marshal(newUserTemplate)

	fmt.Printf("User Json\n\t%v\n\n", string(jsonTest))
	createNewUserAsActive := false

	newUser, _, err := client.Users.Create(newUserTemplate, createNewUserAsActive)

	if err != nil {

		fmt.Printf("Error Creating User:\n \t%v", err)
		return
	}
	fmt.Printf("NewUser Created\n")
	printUser(*newUser)

}
func createUserWithPassword() {
	defer printEnd(printStart("createUserWithPassword"))

	client := okta.NewClient(nil, orgName, apiToken, isProductionOKTAORG)

	fmt.Printf("Client Base URL: %v\n\n", client.BaseURL)

	newUserTemplate := client.Users.NewUser()
	newUserTemplate.Profile.FirstName = "Test SDK First"
	newUserTemplate.Profile.LastName = "Test SDK Last" + time.Now().Format("2006-01-02")
	newUserTemplate.Profile.Login = "testuserwpassword3@localhost.com"
	newUserTemplate.Profile.Email = newUserTemplate.Profile.Login
	newUserTemplate.Profile.DisplayName = "OKTA SDK Test User"
	newUserTemplate.Profile.Division = "IT"
	newUserTemplate.SetPassword("cottoN.hothousE.adoptivE.ivE.87")

	jsonTest, _ := json.Marshal(newUserTemplate)

	fmt.Printf("User Json\n\t%v\n\n", string(jsonTest))

	createNewUserAsActive := true
	newUser, _, err := client.Users.Create(newUserTemplate, createNewUserAsActive)

	if err != nil {

		fmt.Printf("Error Creating User:\n \t%v", err)
		return
	}
	fmt.Printf("NewUser Created\n")
	printUser(*newUser)

}

func createUserWithRecoveryAndPassword() {

	defer printEnd(printStart("createUserWithRecoveryAndPassword"))

	client := okta.NewClient(nil, orgName, apiToken, isProductionOKTAORG)

	fmt.Printf("Client Base URL: %v\n\n", client.BaseURL)

	newUserTemplate := client.Users.NewUser()
	newUserTemplate.Profile.FirstName = "Test SDK First"
	newUserTemplate.Profile.LastName = "Test SDK Last" + time.Now().Format("2006-01-02")
	newUserTemplate.Profile.Login = "testuserwrecovery@localhost.com"
	newUserTemplate.Profile.Email = newUserTemplate.Profile.Login
	newUserTemplate.Profile.DisplayName = "OKTA SDK createUserWithRecoveryAndPassword"
	newUserTemplate.Profile.Division = "IT"

	newUserTemplate.SetPassword("cottoN.hothousE.adoptivE.ivE.87")

	newUserTemplate.SetRecoveryQuestion("What is your car?", "Tesla")

	jsonTest, _ := json.Marshal(newUserTemplate)

	fmt.Printf("User Json\n\t%v\n\n", string(jsonTest))

	createNewUserAsActive := true
	newUser, _, err := client.Users.Create(newUserTemplate, createNewUserAsActive)

	if err != nil {

		fmt.Printf("Error Creating User:\n \t%v", err)
		return
	}
	fmt.Printf("NewUser Created\n")
	printUser(*newUser)

}

func CreateUserThenActivate() {

	defer printEnd(printStart("CreateUserThenActivate"))

	client := okta.NewClient(nil, orgName, apiToken, isProductionOKTAORG)

	fmt.Printf("Client Base URL: %v\n\n", client.BaseURL)

	newUserTemplate := client.Users.NewUser()
	newUserTemplate.Profile.FirstName = "Test SDK First"
	newUserTemplate.Profile.LastName = "Test SDK Last" + time.Now().Format("2006-01-02")
	newUserTemplate.Profile.Login = "CreateUserThenActivate2@localhost.com"
	newUserTemplate.Profile.Email = newUserTemplate.Profile.Login
	newUserTemplate.Profile.DisplayName = "OKTA SDK CreateUserThenActivate"

	jsonTest, _ := json.Marshal(newUserTemplate)

	fmt.Printf("User Json\n\t%v\n\n", string(jsonTest))

	createNewUserAsActive := false
	newUser, _, err := client.Users.Create(newUserTemplate, createNewUserAsActive)

	if err != nil {

		fmt.Printf("Error Creating User:\n \t%v", err)
		return
	}

	fmt.Printf("NewUser Created\n")
	printUser(*newUser)
	sendEmail := false
	activationInfo, _, err := client.Users.Activate(newUser.ID, sendEmail)

	if err != nil {
		fmt.Printf("Error Activating User:\n \t%v\n", err)
		return
	}
	fmt.Printf("User activated :  send this URL to user: %v\n", activationInfo.ActivationURL)

}

func SetUserPassword() {

	defer printEnd(printStart("CreateUserThenActivate"))

	client := okta.NewClient(nil, orgName, apiToken, isProductionOKTAORG)

	newUser, _, err := client.Users.SetPassword("00u8cmwszdE5qztLT0h7", "heRo.laKe.ransOm.23.dogfoOd")

	if err != nil {
		fmt.Printf("Error Setting Password On User:\n \t%v\n", err)
		return
	}
	fmt.Printf("Password Set\n")
	printUser(*newUser)

}

func deactivateUser() {
	defer printEnd(printStart("deactivateUser"))

	client := okta.NewClient(nil, orgName, apiToken, isProductionOKTAORG)
	_, err := client.Users.Deactivate("00u8cmwszdE5qztLT0h7")
	if err != nil {
		fmt.Printf("Could Not deactivate user:\n%v\n", err)
		return
	}

	fmt.Printf("User deactivated\n")
}

func getUserMFAFactor(oktaid string) {
	defer printEnd(printStart("getUserMFAFactor"))

	client := okta.NewClient(nil, orgName, apiToken, isProductionOKTAORG)
	user, _, err := client.Users.GetByID(oktaid)

	if err != nil {
		fmt.Printf("Errr Getting Users:\n \t%v\n", err)
		return
	}

	_, err = client.Users.PopulateMFAFactors(user)

	if err != nil {
		fmt.Printf("Errr Getting MFA Factors:\n \t%v\n", err)
		return
	}
	_, err = client.Users.PopulateGroups(user)
	if err != nil {
		fmt.Printf("Errr Getting Groups:\n \t%v\n", err)
		return
	}
	printUser(*user)

}
