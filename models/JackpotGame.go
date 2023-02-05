package models

import (
	"errors"

	"gorm.io/gorm"
)

type JackpotGames struct {
	gorm.Model
	JackpotMarketID uint `json:"jackpotMarketId" gorm:"foreignKey:id;"`
	HomeTeam string `json:"homeTeam"`
	AwayTeam string `json:"awayTeam"`
	HomeOdds float32 `json:"homeOdds"`
	DrawOdds float32 `json:"drawOdds"`
	AwayOdds float32 `json:"awayOdds"`
}

// SaveJackpotGames methods takes the jackpotMarketID & saves in relation to JackpotMarket
func (j *JackpotGames) SaveJackpotGames() (*JackpotGames, error) {
	var err error

	err = DB.Create(&j).Error
	if err != nil {
		return &JackpotGames{}, err
	}

	return j, nil
}

func (j *JackpotGames) GetJackpotGames(id uint) (interface{}, error) {
	
	tx := DB.Find(j)

	if tx.Error != nil {
		return &JackpotGames{}, errors.New(tx.Error.Error())
	}

	return j, nil
}