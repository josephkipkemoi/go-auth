package routes

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"go-auth-api/go-auth/controllers"
	"go-auth-api/go-auth/controllers/auth"

	"github.com/golang-jwt/jwt/v4"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(setHeaders())	
	r.Use(m())
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

func m() func(*gin.Context) {
	return func(c *gin.Context){
		authHeader := strings.Split(c.GetHeader("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Malformed Token",
			})
		} else {
			jwtToken := authHeader[1]
			t,err := jwt.Parse(jwtToken, func(t *jwt.Token) (interface{}, error) {
				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					// c.JSON(http.StatusBadRequest, gin.H{
					// 	"error": t.Header["alg"],
					// })
					return fmt.Printf("Error: %v", t.Header["alg"])
				}
				return []byte(SECRETKEY), nil
			})
			if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
				
				ctx := context.WithValue(c.Request.Context(), "props", claims)
				c.JSON(http.StatusOK, gin.H{
					"ctx": c.Request.WithContext(ctx),
				})
			} else {
				fmt.Println(err)
				c.JSON(http.StatusUnauthorized, gin.H{
					"error": "Unauthorized",
				})
			}
		}
	}	
}

var SECRETKEY = "maasai"

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Malformed Token"))
		} else {
			jwtToken := authHeader[1]
			t,err := jwt.Parse(jwtToken, func(t *jwt.Token) (interface{}, error) {
				if _,ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
				}
				return []byte(SECRETKEY), nil
			})
			if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
				ctx := context.WithValue(r.Context(), "props", claims)
				next.ServeHTTP(w, r.WithContext(ctx))
			} else {
				fmt.Println(err)
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized"))
			}

		}
		next.ServeHTTP(w,r)
	})
}


