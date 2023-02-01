package main

import (
	"go-auth/go-auth-api/models" // load and connect to database
	r "go-auth/go-auth-api/routes"    // setup router
)


func main() {
	// Connect to Database
	models.ConnectDB()
	// Setup and start server
	r := r.SetupRouter()
	
	r.Run(":8080")	
}