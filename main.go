package main

import (
	"go-auth-api/go-auth/models" // load and connect to database
	"go-auth-api/go-auth/routes"

	// "github.com/gin-gonic/gin"
)

func init() {
	// gin.SetMode("debug")
}

func main() {
	var proxies []string
	proxies = append(proxies, "ipv4")
	
	// Connect to Database
	models.ConnectDB()
	// Setup and start server
	r := routes.SetupRouter()
	r.SetTrustedProxies(proxies)
	r.Run(":8080")	
}