package routes

import "github.com/gin-gonic/gin"

import L "go-auth/go-auth-api/handlers/landing"

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", L.LandingHandler)

	return r
}
