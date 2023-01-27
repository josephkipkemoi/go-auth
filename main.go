package main

import (
	db "go-auth/go-auth-api/database" // load database
	_ "go-auth/go-auth-api/env"       // load environment variables
	r "go-auth/go-auth-api/routes"    // setup router
)


func main() {
	// Connect to database
	db.Connect()
	// Setup and start server
	r := r.SetupRouter()
	r.Run(":8080")	
}