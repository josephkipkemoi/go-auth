package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	// "log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	db "go-auth/go-auth-api/database"
)

type User struct {
	ID int64 `json:"id"`
	PhoneNumber int `json:"phoneNumber" validate:"required"`
	Password string `json:"password" validate:"required"`
	CreatedAt time.Time `json:"createdAt"`
}

type Message struct {
	SuccessMessage string `json:"successMessage"`
	ValidationErrMsg string `json:"validationErrMsg"`
}

var validate *validator.Validate
var token string = "jwt_token"

func RegistrationHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	u := &User{}
	vM := &Message{
		ValidationErrMsg: "Validation Error: Unproccessable Content",
		SuccessMessage: "201 Created",
	}
	
	d := json.NewDecoder(c.Request.Body)
	er := d.Decode(&u)
	if er != nil {
		e, ok := customError(er.Error())
		if !ok {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": e,
			})
			return
		}
	}
	
	validate =	validator.New()
	err := validate.Struct(u)

	errs, ok := validationErrors(err)
	
	if !ok {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": vM.ValidationErrMsg,
			"error": errs,
		})	
		return
	}
	// Data Passed validation
	// Store to db
	res := db.Connect().Create(&u)
	if res.Error != nil {
		e, ok := customError(res.Error.Error())
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": e,
			})
			return
		}
	}

	c.Header("Authorization", "bearer " + token)
	
	c.JSON(201, gin.H{
		"status": vM.SuccessMessage,
		"user": gin.H{
			"id": u.ID,
			"phoneNumber": u.PhoneNumber,
			"createdAt": u.CreatedAt,
		},
		"token": token,
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

func customError(err string) ([]string, bool) {
	errs := []string{}

	if err != "" {
		errs = append(errs, err)
		return errs, false
	}
	return errs, true
}