package tests

import (
	"bytes"
	"encoding/json"
	"fmt"

	"net/http"
	"net/http/httptest"
	"testing"

	"go-auth/go-auth-api/env"
	_ "go-auth/go-auth-api/env"           // load environment variables
	h "go-auth/go-auth-api/handlers/auth" // load auth handlers
	r "go-auth/go-auth-api/routes"        // load routes

	"github.com/stretchr/testify/assert"
)

// Test can register new user
func TestNewUserRegistration(t *testing.T) {
	url := 	env.GetDevAppUrl()
	router := r.SetupRouter()

	w := httptest.NewRecorder()
	// Create user from User Struct
	u := h.User{
		PhoneNumber: 700545785,
		Password: "j",
	}
	// User created json message
	message := h.Message{
		ValidationErrMsg: "Validation Error: Unproccessable Content",
		SuccessMessage: "201 Created",
	}
	// json encode
	e, err := json.Marshal(&u)
	if err != nil {
		fmt.Println(err)
	}
	// Return io.Reader to post to route
	bodyReader := bytes.NewReader(e)
	// Send body
	req, _ := http.NewRequest("POST", url + "api/v1/register", bodyReader)
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code, "Should return 201 status code after user is created")
	assert.Equal(t, message.SuccessMessage, w.Result().Status, "Should have user created success message" )
}