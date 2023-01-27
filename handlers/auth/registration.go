package auth

import (
	"time"

	// h "go-auth/go-auth-api/handlers"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
}

func RegistrationHandler(c *gin.Context) {

	c.Header("Content-Type", "application/json")
	
	c.JSON(201, gin.H{
		"status": "successs",
		"data": "res",
	})
}