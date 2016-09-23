# OKTA SDK for GoLang (Unofficial)

An SDK to integrate with the [OKTA SDK](http://developer.okta.com/documentation/) for Golang.

Much credits has to be given to the [go-github](https://github.com/google/go-github) library for which this package mimics. I stole a huge amount of code and ideas from that project. 


*Warnings* (There are many)

* I have a far from mastered the Go language. 
* This is not a complete SDK. It has been developed based on my own needs.
* Use At your own risk!
* Not all features of the OKTA API are implemented.
* There are a long list of TODOS.
* The current version of the API is focused more around retrieving data from OKTA not actually updating. 


## Design

The design will mimic the Python and JAVA SDKs.



## Users API


Basic Example to to get "me" user which will be the user tied to the API Key. Replace "me" with any OKTA User ideas
```go

import (
	"fmt"
	"github.com/chrismalek/oktasdk-go/okta"
)

func main() {
	orgName := "your org name"
	apiToken := "your-api-key"
    isProductionOrg := false
	client := okta.NewClient(nil, orgName, apiToken, isProductionOrg)

	user, response, err := client.Users.GetByID("me")

	if err != nil {
		fmt.Printf("Response Error %+v, \n\t%+v", err)
	} else {
		fmt.Printf("user: %v\n", user)
	}

}

```

Find Users with a  OKTA user "filter" search, this is also a pagination example

```go

import (
	"fmt"
	"github.com/chrismalek/oktasdk-go/okta"
)

func main() {
	orgName := "your org name"
	apiToken := "your-api-key"
    isProductionOrg := false
	client := okta.NewClient(nil, orgName, apiToken, isProductionOrg)

    userFilter := &okta.UserListFilterOptions{}
	userFilter.FirstNameEqualTo = "Chris"
	userFilter.LastNameEqualTo = "Malek"


    var allUsers []*okta.User

	for {
		userList, response, err := client.Users.ListWithFilter(userFilter)
		if err != nil {
			break
		} else {

			allUsers = append(allUsers, userList...)
			if response.NextURL == nil {
				break
			}
			userFilter.NextURL = response.NextURL
		}

	}

	for _, user := range allUsers {
		fmt.Printf("Found User: %v\n", user.Profile.Login)

	}


}

```





