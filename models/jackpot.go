package models

import "gorm.io/gorm"

type Jackpot struct {
	gorm.Model
	JackpotMarketID uint `json:"jackpotMarketId"`
	HomeTeam string `json:"homeTeam"`
	AwayTeam string `json:"awayTeam"`
	HomeOdds float32 `json:"homeOdds"`
	DrawOdds float32 `json:"drawOdds"`
	AwayOdds float32 `json:"awayOdds"`
}

func (j *Jackpot) SaveJpGames() (*Jackpot, error) {
	var err error

	err = DB.Create(&j).Error
	if err != nil {
		return &Jackpot{}, err
	}

	return j, nil
}

//GetJpGames gets jackpot games for the given jackpot market ID 
func (j *Jackpot) GetJpGames(id int) ([]Jackpot, error) {
	var jpGames []Jackpot
	var query = "SELECT * FROM \"jackpots\" WHERE jackpot_market_id=?"

	tx := DB.Raw(query,id).Scan(&jpGames)
	if tx.Error != nil {
		return jpGames, tx.Error
	}

	return jpGames, nil
}

func (j *Jackpot) UpdateJpGames(id int) (*Jackpot, error) {
	tx := DB.Update("jackpots", j)
	if tx.Error != nil {
		return &Jackpot{}, nil
	}
	return j, nil
}