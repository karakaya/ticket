
## Ticket Support System in Go
Users easily can create ticket for their issues. 
IT workers can track problems and put the priority order. Also, you can get the monthly report, like how much time you spent on problems, what kind of problems you faced with it.
This project will end the phone chaos for basic problems.

### Technologies
- Go
    - Mux
    - GORM


#### Example endpoints
Endpoint|Method|Function
--------    |-------|------------
/user       | POST  | Create User
/user       | GET   | Get All Users
/user/{id}  | GET   | Get Specific User
/user/{id}  | PATCH | Update User
/user/{id}  | DELETE| Delete User

###
#### To start
- If you wanna migrate and initialize first user, just uncommit migration lines from `main.go`
- ``go run main.go`` to start, or `docker-compose up`




### Project Task-List
- [ ] Monthly Reports
- [ ] Integrating departments to follow which department has chronic problems
- [ ] (may) Integrating inventory tracking system


### Technical Task-List 
 

- [ ] Recover from soft-delete
- [ ] .env file
- [ ] Easy insert with validation rules

