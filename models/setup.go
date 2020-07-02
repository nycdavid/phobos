package models

import (
	"github.com/nycdavid/phobos/dbconnector"

	_ "github.com/lib/pq"
)

type Models struct {
	Deck *Decks
	Card *Cards
}

func Preamble(dbo *dbconnector.DBO) *Models {
	return &Models{
		Deck: &Decks{dbo: dbo},
		Card: &Cards{dbo: dbo},
	}
}
