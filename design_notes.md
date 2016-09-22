- TODO: Rename to go-OKTASDK
- TODO: LOOK AT https://github.com/google/go-github for inspiration
- TODO: LOOK AT https://github.com/dghubble/sling as an API client

# Design Notes

http://developer.okta.com/docs/api/getting_started/design_principles.html


Need to respond to rate limits, When getting close to rate limit need to actually sleep or pass that information to the client. If one client hits API rate limit, all other clients timeout.

* the rate limiter could be implemented in the go routine that is pushing work to the workers. If it pushes too much per second the work is being consumed to fast. 

https://support.okta.com/help/articles/Knowledge_Article/API-Rate-Limiting?retURL=%2Fhelp%2Fapex%2FKnowledgeArticleJson%3Fc%3DOkta_Documentation%253ACustomizing_Okta%26p%3D101%26inline%3D1&popup=true

import github.com/chrismalek/oktasdk-go/okta

orgName := "yourORGNAME"
client := okta.NewClient(nil,orgName )


clientConfig := oktaClient.NewClient("https://nu.oktapreview.com", "shh...don't..tell..anyone")

userClient := NewUserClient(clientConfig)

user, err := userClient.getUserByID("alfkdjsl3dfasd")




user := user.New()

User - 
- do we want user to have methods? instead of user client?
 - With a UserClient object, you would have to pass a user object to the client and execute methods from the user client.
 - With methods on an user object like user.Deactivate, user.ResetPassword it may be cleaner in the code.
 - What about examples where you want to create a new user? 
     + user := user.New() user.FirstName
     + user.create()
     + ? Searching for users - May return many different users so the userClient needs differnt methods as well more around searching

- UserClient
    + SearchByName - Returns array of users
    + GetUserByID - Returns user


API Parameters & Design

* limit
* Pagination

## Models

* User
  * profile
  * credentials
      * provider
      * password
      * recovery_question
  * links
* Role Model
    - used to list types of admin user is
* Groups
* Factors
* Apps
* Events



How do we handle custom attributes? 

## Operations

* User
  * Create User
      * Various ways to create users
  * Get user
      * Me
      * By ID
      * By Login
      * Find Operation
      * List User by various filters
          * status, lastupdated, id, profile.login, profile.email, profile.firstName, profile.lastName
  * update User
      - password
      - user object
  * Groups
  * AppLinks
  * activate
  * deactivate
  * suspend
  * unsuspend
  * unlock
  * reset_password
  * expire_password
  * reset_factors
  * forgotpasswords
  * Changepassword
  * change_recovery_question

* Roles (Admin Roles)
* Groups
    - Get Group
    - List Groups (filters)
    - Add Group
    - UPdate Group
    - Remove Group
    - Group Members
    - Add User To Group
    - Remove User From Group
    - List Apps
* Factors
    - Get user FActor(s)
    - Get enrolled factors
    - Eligible factors
    - Enroll in factor
    - reset factor
    - verify factors
