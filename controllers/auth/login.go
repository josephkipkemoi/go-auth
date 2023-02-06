package auth

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"go-auth-api/go-auth/models"
	"go-auth-api/go-auth/utils/tokens"
	"go-auth-api/go-auth/controllers"
)

type LoginInput struct {
	PhoneNumber int64 `json:"phoneNumber"`
	Password string `json:"password"`
}

func LoginHandler(c *gin.Context) {	
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
		 errs, ok := controllers.ValidationErrors(e)
		 if !ok {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": errs,
			})
			return
		 }
	}
	u.PhoneNumber = i.PhoneNumber
	u.Password = i.Password
	// u.AuthUser()

	ok := validateUser(u,i)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"Invalid Credentials",
		})
		return
	}

	token := tokens.GenerateToken()

	c.Header("Authorization", "bearer " + token)

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"user": gin.H{
			"id": u.ID,
			"phoneNumber":u.PhoneNumber,
		},
	})
}	

func validateUser(u *models.User, i *LoginInput) bool {
	if u.Password == i.Password {
		return true
	}
	return false
}