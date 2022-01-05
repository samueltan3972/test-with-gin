package models

import "test/gin-test/database"

type Fruit struct {
	ID            int64   `db:"id, primarykey" json:"id"`
	Fruit_label   string  `db:"fruit_label" json:"email"`
	Fruit_name    string  `db:"fruit_name" json:"fruit_name"`
	Fruit_subtype string  `db:"fruit_subtype" json:"fruit_subtype"`
	Mass          float64 `db:"mass" json:"mass"`
	Width         float64 `db:"width" json:"width"`
	Height        float64 `db:"height" json:"height"`
	Color_score   float64 `db:"color_score" json:"color_score"`
}

type FruitModel struct{}

func (m FruitModel) GetAll() (fruits []Fruit, err error) {
	_, err = database.GetDB().Select(&fruits, "SELECT * FROM fruit")
	return fruits, err
}
