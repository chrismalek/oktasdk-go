# OKTA SDK for GoLang (Unofficial)

An SDK to integrate with the [OKTA SDK](http://developer.okta.com/documentation/) for Golang.

Much credits has to be given to the [go-github](https://github.com/google/go-github) library for which this package mimics. I stole a huge amount of code and ideas from that project. 




*Warnings* (There are many)

* I have far from mastered the Go language. 
* This is not a complete SDK. It has been developed based on my own needs.
* Use At your own risk!
* Not all features of the OKTA API are implemented.
* There are a long list of TODOS.
* The current version of the API is focused more around retrieving data from OKTA not actually updating. 


## Current State

The current state of the this SDK is in development. There is still much work to be done. Currently, we have basic user and group functionality working. We just focus on doing reads from OKTA and have not focused on doing updates.

## Operations

* User
  * Create User (NOT Implemented) &#9785;
      * Various ways to create users
  * Get user
      * Me (Implemented using Users.GetByID pass "me" as parameter)
      * By ID (Implemented Users.GetByID)
      * By Login (Implemented via Users.ListWithFilter passing in Login as filter parameter)
      * List User by various filters (Implemented via Users.ListWithFilter)
          * status, lastupdated, id, profile.login, profile.email, profile.firstName, profile.lastName
	  * [List User with Search (Early Access)](http://developer.okta.com/docs/api/resources/users.html#list-users-with-search)   (NOT Implemented)
  * update User (NOT Implemented) &#9785;
      - password &#9785;
      - user object &#9785;
  * Groups - get user groups (Implemented with Users.PopulateGroups)
  * activate (NOT Implemented) &#9785;
  * deactivate (NOT Implemented) &#9785;
  * suspend (NOT Implemented) &#9785;
  * unsuspend (NOT Implemented) &#9785;
  * unlock (NOT Implemented) &#9785;
  * reset_password (NOT Implemented) &#9785;
  * expire_password (NOT Implemented) &#9785;
  * reset_factors (NOT Implemented) &#9785;
  * forgotpasswords (NOT Implemented) &#9785;
  * Changepassword (NOT Implemented) &#9785;
  * change_recovery_question (NOT Implemented) &#9785;
* Roles (Admin Roles) (NOT Implemented) &#9785;
* Groups (okta.Groups)
    - Get Group (Implemented with Groups.GetByID) &#9745;
    - List Groups (Implemented with Groups.ListWithFilter) &#9745;
    - Add Group (Implemented Groups.Add) &#9745;
    - Update Group (NOT Implemented) &#9785;
    - Delete Group (Implemented Groups.Delete) &#9745;
    - Group Members (Implemented with Groups.GetUsers)
    - Add User To Group (NOT Implemented) &#9785;
    - Remove User From Group (NOT Implemented) &#9785;
    - List Apps (NOT Implemented) &#9785;
* Factors (NOT Implemented)
    - Get user FActor(s) (NOT Implemented) &#9785;
    - Get enrolled factors (NOT Implemented) &#9785;
    - Eligible factors (NOT Implemented) &#9785;
    - Enroll in factor (NOT Implemented) &#9785;
    - reset factor (NOT Implemented) &#9785;
    - verify factors (NOT Implemented) &#9785;



# OKTA Links

Important OKTA Links

http://developer.okta.com/docs/api/getting_started/design_principles.html



## Example code


There are several runnable examples in the `/examples` directory. You can run them by using the following:

First you need two environment variables configured:

* `OKTA_API_TEST_ORG` - The Orgname of your "Preview OKTA Environment"
* `OKTA_API_TEST_TOKEN` - An API Token for your preview ORG

```
cd $GOPATH/src/github.com/chrismalek/oktasdk-go
go run examples/*.go
```

You can see some examples in that code in the usage:

* `/examples/userExamples.go` - Examples Using the User Client
* `/examples/groupExamples.go` - Exaples using the Group Client






