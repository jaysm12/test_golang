## Simple Golang API - Take Home Test
This repository contains a simple Golang API, developed as a take-home test for a recruitment process. The project demonstrates the creation and usage of a basic RESTful API using the Go programming language with Mysql.

### Requirements and Dependencies
```
Go 1.20 or later
Gorilla Mux 1.8.0
Gorm 1.9.16
```

### Installation
To install the required dependencies and set up the project, follow these steps:

+ Clone the repository:
```
git clone https://github.com/jaysm12/test_golang_otomo.git
```
+ Change to the project directory:
```
cd test_golang_otomo
```

+ Download the required Go modules:
```
go mod download
```
+ Change gorm database config in `/pkg/config/app.go` with your database

```
"<user>>:<password>@/otomo?charset=utf8&parseTime=true&loc=Local"
```
### Usage
To run the application, execute the following command:
```
cd cmd/main
go build && ./main
Server running on port: 3000
```

### API Endpoints
The following API endpoints are available:

+ `POST api/user/regiser` : Register new user

Body : 
```
{
  username (string)
  password (string)
}
```
+ `POST api/user/login` : Login existing user to obtain a token

Body:
```
{
  username (string)
  password (string)
}
```
Response:
```
{
  "success" : true
  "token" : <token>
}
```

+ `GET api/user` : Get List User

Header Auth : `Bearer <token>`

Response:
```
{
  "success" : true
  "users" : [{users}]
}
```

+ `POST api/user/delete`: Delete user by username and password
Header Auth : `Bearer <token>`
Body:
```
{
  username (string)
  password (string)
}
```
