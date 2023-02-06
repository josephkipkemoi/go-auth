package routes

import (

	"github.com/gin-gonic/gin"

	"go-auth-api/go-auth/controllers"
	"go-auth-api/go-auth/controllers/auth"
	// "go-auth/go-auth-api/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(setHeaders())	

	public := r.Group("/api/v1/")

	public.GET("/", controllers.LandingHandler)
	public.POST("/register",auth.Register)
	public.POST("/login", auth.LoginHandler)
	public.POST("/jackpots", controllers.StoreMarket)
	public.POST("/jackpots/games", controllers.Store)
	public.GET("/jackpots/games", controllers.Show)
	public.PATCH("/jackpots/games/patch", (controllers.Update))

	return r
}

func setHeaders() func(*gin.Context) {
	return func(c *gin.Context){
		c.Header("Content-Type","application/json:charset=utf-8")
		c.Header("Host", c.Request.Host)
		c.Header("X-Powered-By", "go/1.19")
		c.Header("Access-Control-Allow-Origin", "http://localhost:3000/")
		c.Header("Access-Control-Allow-Credentials", "true")
	}	
}

