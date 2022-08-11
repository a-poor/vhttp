# vhttp

__created by Austin Poor_

A library for testing HTTP requests and responses from the `net/http` package.

## TO-DO 

**Requests:**
- [ ] Method
  - [x] Is x
  - [x] IsNot x
  - [x] Is (GET/POST/PUT/...)
- [ ] URL
  - [X] Has...
    - [X] ...Scheme
    - [X] ...User
    - [X] ...Scheme
    - [X] ...Host
    - [X] ...Path
  - [ ] Path variables
  - [X] Query variables
  - [X] Path matches glob
  - [X] Is x
- [ ] Headers
  - [X] Has header x
  - [ ] Doesn't have header x
- [ ] Body (as bytes)
- [ ] TCP
- [ ] Form Data

**Responses:**
- [ ] Status Code
  - [ ] Is x
  - [ ] IsNot x
  - [ ] IsIn [x]
  - [ ] IsNotIn [x]
  - [ ] InRange x-y
  - [ ] NotInRange x-y
  - [ ] IsSuccess
  - [ ] IsError
  - [ ] IsRedirect
  - [ ] IsClientError
  - [ ] IsServerError
- [ ] Headers
  - [ ] Has header x
  - [ ] Doesn't have header x
- [ ] Body (as bytes)
- [ ] TCP
- [ ] Form Data
