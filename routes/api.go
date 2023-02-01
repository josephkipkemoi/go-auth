package routes

import (
	"github.com/gin-gonic/gin"

	auth "go-auth/go-auth-api/handlers/auth"

	l "go-auth/go-auth-api/handlers"
	
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	public := r.Group("/api")

	public.GET("/", l.LandingHandler)
	public.POST("/v1/register", auth.Register)
	public.POST("/v1/login", auth.LoginHandler)

	return r
}
