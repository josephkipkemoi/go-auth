package routes

import (
	"github.com/gin-gonic/gin"

	auth "go-auth/go-auth-api/controllers/auth"

	"go-auth/go-auth-api/controllers"
	
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	public := r.Group("/api")

	public.GET("/", controllers.LandingHandler)
	public.POST("/v1/register", auth.Register)
	public.POST("/v1/login", auth.LoginHandler)
	public.POST("v1/jackpots", controllers.StoreMarket)
	public.POST("v1/jackpots/games", controllers.StoreGames)

	return r
}
