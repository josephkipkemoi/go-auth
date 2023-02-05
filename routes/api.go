package routes

import (
	"github.com/gin-gonic/gin"

	auth "go-auth/go-auth-api/controllers/auth"

	"go-auth/go-auth-api/controllers"
	
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	
	public := r.Group("/api/v1/")

	public.GET("/", controllers.LandingHandler)
	public.POST("/register", auth.Register)
	public.POST("/login", auth.LoginHandler)
	public.POST("/jackpots", controllers.StoreMarket)
	public.POST("/jackpots/games", controllers.StoreGames)
	public.GET("/jackpots/games?jp_id=1", controllers.ShowJackpotGames)

	return r
}
