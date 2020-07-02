package models

import (
	"github.com/nycdavid/phobos/dbconnector"
)

type Cards struct {
	dbo *dbconnector.DBO
}

type Card struct {
	Id    int    `json:"id"`
	Front string `json:"front"`
	Back  string `json:"front"`
}

func (cs *Cards) All() []*Card {
	cards := make([]*Card, 0)

	return cards
}

func (cs *Cards) Create(c *Card) *Card {
}
