package models

import "gorm.io/gorm"

type JackpotGames struct {
	gorm.Model
	JackpotMarketID uint 
	HomeTeam string
	AwayTeam string
	HomeOdds float32
	DrawOdds float32
	AwayOdds float32
}

// SaveJackpotGames methods takes the jackpotMarketID & saves in relation to JackpotMarket
func (j *JackpotGames) SaveJackpotGames() (*JackpotGames, error) {
	var err error

	err = DB.Create(&j).Error
	if err != nil {
		return &JackpotGames{}, err
	}

	return &JackpotGames{}, nil
}