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
  * Create User (Implemented)  &#9745;
      * Create User without password (Implemented in user.Create )  &#9745;
        * See Examples `examples/userExamples.go createUserNoPassword()`
      * Create User with password (Implemented in user.Create )  &#9745;
        * See Examples `examples/userExamples.go createUserWithPassword()`
  * Get user
      * Me (Implemented using Users.GetByID pass "me" as parameter)  &#9745;
      * By ID (Implemented Users.GetByID) &#9745;
      * By Login (Implemented via Users.ListWithFilter passing in Login as filter parameter)  &#9745;
      * List User by various filters (Implemented via Users.ListWithFilter)  &#9745;
          * status, lastupdated, id, profile.login, profile.email, profile.firstName, profile.lastName
	  * [List User with Search (Early Access)](http://developer.okta.com/docs/api/resources/users.html#list-users-with-search)   (NOT Implemented)  &#9785;
  * update User (NOT Implemented) &#9785;
      - password &#9785;
      - user object &#9785;
  * Groups - get user groups (Implemented with Users.PopulateGroups) &#9745;
  * activate (implemented in Users.Activate) &#9745;
  * deactivate (implemented in Users.Deactivate) &#9745;
  * suspend (implemented in Users.Suspend) &#9745;
  * unsuspend (implemented in Users.Unsuspend) &#9745;
  * unlock (implemented in Users.Unlock) &#9745;
  * reset_password (implemented in Users.ResetPassword) &#9745;
  * SetPassword (Implemented in Users.SetPassword) &#9745;
  * expire_password (NOT Implemented) &#9785;
  * reset_factors (NOT Implemented) &#9785;
  * forgotpassword (NOT Implemented) &#9785;
  * change_recovery_question (NOT Implemented) &#9785;
  * List Enrolled Factors (implemented in Users.PopulateEnrolledFactors)  &#9745;
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
    - (implemented in Users.PopulateEnrolledFactors)  &#9745;
    - Eligible factors (NOT Implemented) &#9785;
    - Enroll in factor (NOT Implemented) &#9785;
    - reset factor (NOT Implemented) &#9785;
    - verify factors (NOT Implemented) &#9785;
* Apps (Barely Implemented)
    - get App (Apps.GetByID) &#9745;
    - get App Users (Apps.GetUsers)  &#9745;
    - Many more API Interactions to go &#9785;


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






