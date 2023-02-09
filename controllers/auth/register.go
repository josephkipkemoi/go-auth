package auth

import (
	"go-auth-api/go-auth/controllers"
	"go-auth-api/go-auth/models"
	"go-auth-api/go-auth/utils/tokens"
	"log"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type RegistrationInput struct {
	PhoneNumber int64 `json:"phoneNumber" binding:"required" validate:"required"`
	Password string `json:"password" binding:"required" validate:"required"`
}

var validate *validator.Validate

func Register(c *gin.Context) {
	i := &RegistrationInput{}
	if err := c.ShouldBindJSON(i); err != nil {
		errs, ok := controllers.ValidationErrors(err)
		if !ok {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": errs,
			})
			return
		}
	}

	u := models.User{}
	u.PhoneNumber = i.PhoneNumber
	u.Password = i.Password
	u.HashPassword()
	_, err := u.SaveUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	
	token,err := tokens.GenerateToken()
	log.Fatal(err)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err,
		})
		return
	}
	c.Header("Authorization", "Bearer " + token)
	
	c.JSON(201, gin.H{
		"user": gin.H{
			"id": u.ID,
			"phoneNumber": u.PhoneNumber,
			"createdAt": u.CreatedAt,
		},
	})
}

