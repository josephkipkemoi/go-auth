package models

import "gorm.io/gorm"

// JackpotMarket hasMany Jackpot games
type JackpotMarket struct {
	gorm.Model
	Market string `gorm:"not null;" json:"market"`
	MarketID uint 
	JackpotGames []JackpotGames `gorm:"constraint:OnUpdate:CASCADE, onDelete: SET NULL;"`
}

func (j *JackpotMarket) SaveJackpotMarket() (*JackpotMarket, error){
	var err error
	err = DB.Create(&j).Error

	if err != nil {
		return &JackpotMarket{}, err
	}

	return j, nil
}