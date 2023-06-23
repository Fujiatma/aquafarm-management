# aquafarm-management
* This project build with Golang 1.19, prototype for managing a simple aquafarm
* Using MySQL for the Database
* Using gorm

# Project Structure
- main.go
- src
    - config
        - database.go
    - controllers
        - controller.go
    - models
        - common.go
        - farm.go
        - pond.go
        - statistic.go
        - user.go
    - repositories
        - farm_repository.go
        - pond_repository.go
        - statistic_repository.go
        - user_repository.go
    - request
        - request.go
    - response
        - response.go
    - routes
        - routes.go
    - services
        - farm.go
        - farm_.go
        - pond.go
        - pond_.go
        - statistic.go
        - statistic_.go
        - user.go
        - user_.go
- test
    - service_test
        - farm_service_test.go
        - pond_service_test.go
        - statistic_service_test.go
        - user_service_test.go
    - mocks
        - farm_repository_mock.go
        - pond_repository_mock.go
        - statistic_repository_mock.go
        - user_repository_mock.go

# Explanation
- main.go:The main file to run the application.
- src: The main directory of the application.
    - config: Contains configuration-related files, such as database settings.
    - controllers: Contains controller logic to handle HTTP requests.
    - models: Contains the data model definitions used in the application.
    - repositories: Contains the database access logic.
    - routes: Contains the definitions of HTTP routes to be handled by the application.
    - services: Contains the business logic for each application feature.



# How to run/install the Project
- go mod tidy
- config the database (MySQL) on your local. You can look to .env file the configuration of the DB
- start the project
    - "go run main.go"  and will run in http://localhost:8000


# Project Scope
This Project is handle the all requirements below:
- The application should be able to handle a POST request to create a farm or pond. Request to create a duplicate entry should be denied.
- The application should be able to handle a PUT request to update an existing farm or pond or create a new one if the entity specified in
  the payload doesn't exist yet.
- The application should be able to handle a DELETE request to soft delete an existing farm or pond. Throw an error if the specified farm
  or pond doesn't exist yet.
- The application should be able to handle a GET request to get a list of all currently existing farms or ponds in the system. If no entity is
  found then the API should return the HTTP status 404 Not Found.
- The application should be able to handle a GET request specifically for a farm or a pond by using its ID in the path parameter. If the ID
  specified doesn’t exist in the storage, the API should return the HTTP status 404 Not Found.
- The application should also be able to count how many times each endpoint is called. The statistics would then be able to be viewed with
  a GET request in a JSON format.


# Database Design and DDL
- You can see the DB Design and the DDL from Folder /documentation