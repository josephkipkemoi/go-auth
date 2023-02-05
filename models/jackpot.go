package models

import (
	"gorm.io/gorm"
)

type Jackpot struct {
	gorm.Model
	JackpotMarketID uint `json:"jackpotMarketId"`
	FixtureID uint `json:"fixtureId"`
	HomeTeam string `json:"homeTeam"`
	AwayTeam string `json:"awayTeam"`
	HomeOdds float32 `json:"homeOdds"`
	DrawOdds float32 `json:"drawOdds"`
	AwayOdds float32 `json:"awayOdds"`
}

func (j *Jackpot) Save() (*Jackpot, error) {
	var err error

	err = DB.Create(&j).Error
	if err != nil {
		return &Jackpot{}, err
	}

	return j, nil
}

func (j *Jackpot) Show(id int) ([]Jackpot, error) {
	var jpGames []Jackpot
	var query = "SELECT * FROM \"jackpots\" WHERE jackpot_market_id=?"

	if err := DB.Raw(query,id).Scan(&jpGames).Error; err != nil {
		return []Jackpot{}, err
	}

	return jpGames, nil
}

func (j *Jackpot) Update(id int) (*Jackpot, error) {
	tx := DB.Model(j).Where("id = ?", id).Updates(j)
	if tx.Error != nil {
		return &Jackpot{}, nil
	}

	return j, nil
}