package landing

import "github.com/gin-gonic/gin"

func LandingHandler(c *gin.Context) {
	c.String(200, "Golang Auth API")
}