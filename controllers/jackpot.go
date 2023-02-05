package controllers

import (
	"encoding/json"
	"go-auth/go-auth-api/models"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type JackpotMarketInput struct {
	Market string `json:"market" binding:"required" validate:"required"`
	JackpotMarketID uint `json:"marketId" validate:"required"`
}

var validate *validator.Validate

func StoreMarket(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	i := &JackpotMarketInput{}

	d := json.NewDecoder(c.Request.Body)
	err := d.Decode(i)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"errors": "JSON Parse Error",
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
	m.JackpotMarketID = i.JackpotMarketID

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

type JackpotGamesInput struct {
	JackpotMarketID uint `validate:"required"`
	HomeTeam string `validate:"required"`
	AwayTeam string `validate:"required"`
	HomeOdds float32 `validate:"required"`
	DrawOdds float32 `validate:"required"`
	AwayOdds float32 `validate:"required"`
}

func StoreGames(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	j := &models.Jackpot{}
	i := &JackpotGamesInput{}

	d := json.NewDecoder(c.Request.Body)
	e := d.Decode(i)
	if e != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"errors": e.Error(),
		})
		return
	}

	validate = validator.New()
	err := validate.Struct(i)
	errs, ok := ValidationErrors(err)
	if !ok {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": errs,
		})
		return
	}

	j.JackpotMarketID = i.JackpotMarketID
	j.HomeTeam = i.HomeTeam
	j.AwayTeam = i.AwayTeam
	j.HomeOdds = i.HomeOdds
	j.DrawOdds = i.DrawOdds
	j.AwayOdds = i.AwayOdds

	data, er := j.SaveJpGames()
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": er,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": data,
	})
}

func ShowJackpotGames(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	
	j := models.Jackpot{}
	q := c.Request.FormValue("jp_id")
	id, err := strconv.Atoi(q)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	data, e := j.GetJpGames(id)
	if e != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": e,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}