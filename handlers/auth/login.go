package auth

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	db "go-auth/go-auth-api/database"
)


func LoginHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	u := &User{}
	d := json.NewDecoder(c.Request.Body)
	err := d.Decode(u)
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
	er := validate.Struct(u)
	if er != nil {
		 errs, ok := validationErrors(er)
		 if !ok {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": errs,
			})
			return
		 }
	}
	// Passed validation
	// Check record in DB
	res := db.Connect().Find(u, u.Password)
	if res.Error != nil {
		errs, ok := customError(res.Error.Error())
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errs,
			})
			return
		}
	}

	c.Header("Authorization", "bearer " + token)

	c.JSON(http.StatusNoContent, gin.H{
		"status": "success",
		"user": gin.H{
			"id": u.ID,
			"phoneNumber": u.PhoneNumber,
			"isVerified": u.IsVerified,
			"createdAt": u.CreatedAt,
		},
	})
}	