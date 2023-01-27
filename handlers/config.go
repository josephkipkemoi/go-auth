package handlers

import (
	"net/http"
	"time"
)

// Initialize handlers package with pre-configured headers
var Client *http.Client

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