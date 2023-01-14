package main

import r "go-auth/go-auth-api/routes"

func main() {
	r := r.SetupRouter()

	r.Run(":8080")	
}