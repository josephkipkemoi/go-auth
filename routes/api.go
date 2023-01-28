package routes

import (
	"github.com/gin-gonic/gin"

	auth "go-auth/go-auth-api/handlers/auth"
	l "go-auth/go-auth-api/handlers/landing"

	_ "go-auth/go-auth-api/handlers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", l.LandingHandler)
	r.POST("/api/v1/register", auth.RegistrationHandler)

	return r
}
