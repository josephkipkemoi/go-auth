package controllers

import (
	"github.com/gin-gonic/gin"
)

type Message struct {
	Message string `json:"message"`
}

var WelcomeMessage string = "Golang Auth API"

func LandingHandler(c *gin.Context) {
	message := Message{Message: WelcomeMessage}
	
	c.Header("Content-Type", "application/json")
	c.JSON(200, gin.H{
		"message": message.Message,
	})
}