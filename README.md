CRUD-Clean Architecture in Go
This project demonstrates a simple implementation of clean architecture using Go, the Echo framework, and GORM with a PostgreSQL database. It provides two basic APIs for user management: registration and login.

Project Overview
Language: Go
Framework: Echo
Database: PostgreSQL
ORM: GORM
API Endpoints
1. Register User
Endpoint: /register
Method: POST
Description: Registers a new user.
Request Body:
json
{
  "username": "string",
  "name": "string",
  "email": "string",
  "phone": "string",
  "password": "string"
}
Response:
json
{
  "message": "User registered successfully."
}
2. Login User
Endpoint: /login
Method: POST

Request Body:
json

{
  "username": "string",
  "password": "string"
}

Installation
Clone the repository:

bash

git clone https://github.com/adi-dev-x/Basic-Crud-for-clean-architecture-in-GO

Change to the project directory:

bash

cd Basic-Crud-for-clean-architecture-in-GO
Install dependencies:

bash

go mod download
Run the application:

bash
Copy code
go run cmd/usermanagement/main.go
