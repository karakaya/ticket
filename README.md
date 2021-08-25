
##Ticket Support System in Go
Users easily can create ticket for their issues.
IT workers can track problems and put the priority order.
Also, you can get the monthly report. Like how much time you spent for problems, what kind of problems you faced with it.
The project will end the phone chaos for basic problems.

###Technologies
- Go
    - Mux
    - GORM


####Example endpoints
Endpoint|Method|Function
--------    |-------|------------
/user       | POST  | Create User
/user       | GET   | Get All Users
/user/{id}  | GET   | Get Specific User
/user/{id}  | PATCH | Update User
/user/{id}  | DELETE| Delete User

###
####To start
- ``go run main.go`` start with migration




###Project Task-List
- [ ] Monthly Reports
- [ ] Integrating departments to follow which department has chronic problems
- [ ] (may) Integrating inventory tracking system


###Technical Task-List 
- [ ] Tests
- [ ] Authentication
- [ ] Middleware
- [ ] Recover from soft-delete
- [ ] .env file
- [ ] Easy migration
- [ ] Easy insert with validation rules 
- [ ] CSRF Protection
- [ ] Dockerfile
- [ ] Basic front-end for demo