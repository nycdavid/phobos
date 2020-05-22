package models

import (
	"github.com/nycdavid/phobos/dbconnector"

	_ "github.com/lib/pq"
)

type Models struct {
	Deck *Decks
}

func Preamble(dbo *dbconnector.DBO) *Models {
	return &Models{
		Deck: &Decks{dbo: dbo},
	}
}
