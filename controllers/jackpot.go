package controllers

import (
	"encoding/json"
	"go-auth/go-auth-api/models"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type JackpotMarketInput struct {
	Market string `json:"market" binding:"required" validate:"required"`
	MarketID uint `json:"marketId" validate:"required"`
}

var validate *validator.Validate

const (
	MegaJackpotMarketId uint = uint(1)
	JackpotFiveMarketId
)

func StoreMarket(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	i := &JackpotMarketInput{}

	// encoder(inputStruct,c.Request.Body) e
	d := json.NewDecoder(c.Request.Body)
	err := d.Decode(i)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"errors": err.Error(),
		})
		return
	}

	validate = validator.New()
	e := validate.Struct(i)

	errs, ok := ValidationErrors(e)
	if !ok {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": errs,
		})
		return
	}

	m := models.JackpotMarket{}
	m.Market = i.Market

	data, e := m.SaveJackpotMarket()
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": e,
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "201 Created",
		"market": data.Market,
	})
}

func StoreGames(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	
}