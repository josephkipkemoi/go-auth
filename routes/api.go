package routes

import(
	 "github.com/gin-gonic/gin"

	 l "go-auth/go-auth-api/handlers/landing"
	 auth "go-auth/go-auth-api/handlers/auth"

)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// Landing Handler
	r.GET("/", l.LandingHandler)
	// Auth Handler
	r.POST("/api/v1/register", auth.RegistrationHandler)

	return r
}
