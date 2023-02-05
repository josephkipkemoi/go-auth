package models

import "gorm.io/gorm"

const (
	MegaJackpotMarketId uint = uint(201)
	JackpotFiveMarketId
)
// JackpotMarket hasMany Jackpot games
type JackpotMarket struct {
	gorm.Model
	Market string `gorm:"not null;" json:"market"`
	JackpotMarketID uint  `json:"jackpotMarketId" gorm:"foreignKey:id"`
	JackpotGames []JackpotGames `json:"jackpotGames" gorm:"foreignKey:id"`
}

func (j *JackpotMarket) SaveJackpotMarket() (*JackpotMarket, error){
	var err error
	err = DB.Create(&j).Error

	if err != nil {
		return &JackpotMarket{}, err
	}

	return j, nil
}