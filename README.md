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

# Sample CURL To Run the API
```- Create User
  curl --location --request POST 'http://localhost:8000/v1/aquafarm/user' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "your-name"
  }'

- Create Farm:
curl --location --request POST 'http://localhost:8000/v1/aquafarm/farm' \
--header 'Content-Type: application/json' \
--data-raw '{
    "farm_name": "Farm Surabaya",
    "user_id": "7335983b-b7a4-4473-bf21-401600c12f5b"
  }'

- Create Pond:
curl --location --request POST 'http://localhost:8000/v1/aquafarm/pond' \
--header 'Content-Type: application/json' \
--data-raw '{
  "pond_name": "Pond Bandung 1",
  "farm_id": "dfe1b9f8-4919-4353-a576-08c4664c7e4c",
  "user_id": "804056c5-3bd2-47a7-a8de-aa9d713c09c1"
}'

- Update Farm:
curl --location --request PUT 'http://localhost:8000/v1/aquafarm/farm/7b58e567-3b00-4e1f-b504-d8790aa65c32' \
--header 'Content-Type: application/json' \
--data-raw '{
  "farm_name": "Farm Bandung 2",
  "user_id": "7335983b-b7a4-4473-bf21-401600c12f5b"
}'

- Update Pond:
curl --location --request PUT 'http://localhost:8000/v1/aquafarm/pond/1ca67bb8-8d9d-49c2-bd37-0eb28fe508a3' \
--header 'Content-Type: application/json' \
--data-raw '{
  "pond_name": "Pond Bandung Updated",
  "farm_id": "dfe1b9f8-4919-4353-a576-08c4664c7e4c",
  "user_id": "7335983b-b7a4-4473-bf21-401600c12f5b"
}'

- Delete Farm:
curl --location --request DELETE 'http://localhost:8000/v1/aquafarm/farm/dfe1b9f8-4919-4353-a576-08c4664c7e4c' \
--header 'Content-Type: application/json' \
--data-raw '{
  "user_id": "7335983b-b7a4-4473-bf21-401600c12f5b"
}'

- Delete Pond:
curl --location --request DELETE 'http://localhost:8000/v1/aquafarm/pond/1ca67bb8-8d9d-49c2-bd37-0eb28fe508a3' \
--header 'Content-Type: application/json' \
--data-raw '{
  "user_id": "7335983b-b7a4-4473-bf21-401600c12f5b"
}'

- Get All Farms:
curl --location --request GET 'http://localhost:8000/v1/aquafarm/farms?user_id=7335983b-b7a4-4473-bf21-401600c12f5b'

- Get All Ponds:
curl --location --request GET 'http://localhost:8000/v1/aquafarm/ponds?user_id=7335983b-b7a4-4473-bf21-401600c12f5b'

- Get Farm By ID:
curl --location --request GET 'http://localhost:8000/v1/aquafarm/farm/7b58e567-3b00-4e1f-b504-d8790aa65c32?user_id=7335983b-b7a4-4473-bf21-401600c12f5b'

- Get Pond By ID:
curl --location --request GET 'http://localhost:8000/v1/aquafarm/pond/1ca67bb8-8d9d-49c2-bd37-0eb28fe508a3?user_id=7335983b-b7a4-4473-bf21-401600c12f5b'

- Get All Statistics:
curl --location --request GET 'http://localhost:8000/v1/aquafarm/statistics'
