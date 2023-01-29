# Golang JWT Auth with Go Gin and Postgres SQL
## Dependencies
+ Go 1.19 > 
## Getting Started
### How do I install the application?
```
$ git clone https://github.com/josephkipkemoi/go-auth.git
$ cd go-auth
install go modules
```
### Running the application?
`$ go run main.go`
### Running the tests?
`$ go test ./tests`
## API REFERENCE
### Landing Route
```
`/`
+ GET - Get landing message, returned as JSON
```
| METHOD | ENDPOINT | HEADERS | PARAMS | STATUS | RESPONSE (JSON) |
| ----------- | -------- | ------- | ------ | ----------- | -------- |
| GET    | /        | Content-Type: application/json    | N/A | Success: 200 | {"message": "Golang Auth API"}
| POST   | /api/v1/register | Content-Type: application/json , Authorization: bearer jwt_token | {"phoneNumber": int, "password": string} | Success: 201,Unproccessable Entity: 422, Bad Request: 400| { "http_status": 201, "status":"User Created","user: {"id": int, "phoneNumber": int "isVerified": bool, "createdAt": string}, "token": "jwt_token"}  / {"http_status": 422, "errors": array} / {"http_status:400": "errors"}
| POST   | /api/v1/login  | Content-Type: application/json, Authorization: bearer jwt_token | {"phoneNumber": int, "password": string} | Success: 200, Not Found: 404, Bad Request: 400 | {"http_status": 200, "status": "success", "user": {"id": int, "phoneNumber": int, "isVerified": bool, "createdAt": string}, "token": "jwt_token"}, {"http_status": 404, "error": array}, {"http_status": 400, "error": array}
