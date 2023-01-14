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
| METHOD | ENDPOINT | HEADERS | PARAMS | STATUS | RESPONSE (JSON) |
| ----------- | -------- | ------- | ------ | ----------- | -------- |
| GET    | /        | Content-Type: application/json    | N/A | Success: 200 | {"message": "Golang Auth API"}

