package auth

import (
	"go-auth/go-auth-api/models"
	"go-auth/go-auth-api/utils/tokens"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type RegistrationInput struct {
	PhoneNumber int64 `json:"phoneNumber" binding:"required" validate:"required" gorm:"unique"`
	Password string `json:"password" binding:"required" validate:"required"`
}

var validate *validator.Validate

func Register(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	i := &RegistrationInput{}
	if err := c.ShouldBindJSON(i); err != nil {
		errs, ok := validationErrors(err)
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
	
	token := tokens.GenerateToken()

	c.Header("Authorization", "bearer " + token)
	
	c.JSON(201, gin.H{
		"user": gin.H{
			"id": u.ID,
			"phoneNumber": u.PhoneNumber,
			"createdAt": u.CreatedAt,
		},
	})
}

// validationErrors returns found errors stored in a slice and true if errors are found empty slice and false otherwise
func validationErrors(err error) ([]string, bool) {
	errs := []string{}

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return errs, false
		}
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println()
			fmt.Println(err)
			errs = append(errs, err.Field() + " Field is required")
		}
		return errs, false
	}
	return errs,true
}

