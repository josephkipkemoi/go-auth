package auth

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	// db "go-auth/go-auth-api/database"
	"go-auth/go-auth-api/models"
)

type LoginInput struct {
	PhoneNumber int64 `json:"phoneNumber"`
	Password string `json:"password"`
}

func LoginHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	
	u := &models.User{}
	i := &LoginInput{}

	d := json.NewDecoder(c.Request.Body)
	err := d.Decode(i)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{	
			"error": err.Error(),
		})
		return
	}

	validate = validator.New()
	e := validate.Struct(i)
	if e != nil {
		 errs, ok := validationErrors(e)
		 if !ok {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": errs,
			})
			return
		 }
	}
	u.PhoneNumber = i.PhoneNumber
	u.Password = i.Password
	u.AuthUser()

	ok := validateUser(u,i)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"Invalid Credentials",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"user": gin.H{
			"id": u.ID,
			"phoneNumber":u.PhoneNumber,
			// "token":r,
		},
	})
}	

func validateUser(u *models.User, i *LoginInput) bool {
	if u.Password == i.Password {
		return true
	}
	return false
}