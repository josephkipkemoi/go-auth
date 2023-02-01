package auth

import (
	"encoding/json"
	// "log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	db "go-auth/go-auth-api/database"
)

type LoginInput struct {
	PhoneNumber int `json:"phoneNumber"`
	Password string `json:"password"`
}

func LoginHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	
	u := &User{}
	i := &LoginInput{}

	d := json.NewDecoder(c.Request.Body)
	err := d.Decode(i)
	if err != nil {
		e, ok := customError(err.Error())
		if !ok {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": e,
			})
			return
		}
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
	
	// Passed validation
	// Check record in DB
	res := db.Connect().Where("phone_number", i.PhoneNumber).First(u)
	if res.Error != nil {
		errs, ok := customError(res.Error.Error())
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errs,
			})
			return
		}
	}

	ok := validateUser(u,i)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"Invalid Credentials",
		})
		return
	}

	c.Header("Authorization", "bearer " + token)

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"user": gin.H{
			"id": u.ID,
			"phoneNumber": u.PhoneNumber,
		},
	})
}	

func validateUser(u *User, i *LoginInput) bool {
	// i.Password = HashPassword(i.Password)
	if u.Password == i.Password {
		return true
	}
	return false
}