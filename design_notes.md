- TODO: Rename to go-OKTASDK
- TODO: LOOK AT https://github.com/google/go-github for inspiration
- TODO: LOOK AT https://github.com/dghubble/sling as an API client

# Design Notes

http://developer.okta.com/docs/api/getting_started/design_principles.html



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
