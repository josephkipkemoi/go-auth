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
| HTTP METHOD | ENDPOINT | HEADERS | PARAMS | STATUS CODE | RESPONSE |
| ----------- | -------- | ------- | ------ | ----------- | -------- |
| GET    | /        | Content-Type: application/json    | ------ | Success: 200 | json {"message": "Golang Auth API"}

