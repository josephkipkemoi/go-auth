package controllers

import (
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
)

// Initialize handlers package with pre-configured headers
var Client *http.Client
// Initiliaze validation package
var Validate *validator.Validate


// REST Methods Predefined
func init() {
	// Control over:	*TLS Config
	// 					*Keep alive
	// 					*Compression
	tr := &http.Transport{
		MaxIdleConns: 10,
		IdleConnTimeout: 30 * time.Second,
		DisableCompression: true,
	}
	// Control over:	*Client headers
	// 					*Redirect policies
	// 					*
	Client = &http.Client{
		Transport: tr,
	}

}

func main(){
	Validate =	validator.New()

}