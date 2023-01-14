package landing

import (
	"github.com/gin-gonic/gin"
)

func LandingHandler(c *gin.Context) {
	message := "Golang Auth API"
	
	c.Header("Content-Type", "application/json")
	c.JSON(200, gin.H{
		"message": message,
	})
}