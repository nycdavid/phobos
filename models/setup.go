package models

import (
	"log"

	"github.com/nycdavid/phobos/dbconnector"

	_ "github.com/lib/pq"
)

type Models struct {
	Deck *Decks
}

type Decks struct {
	dbo *dbconnector.DBO
}

type Deck struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (d *Decks) All() []*Deck {
	var decks []*Deck

	rows, e := d.dbo.Conn.Query("SELECT * FROM decks;")
	if e != nil {
		log.Println(e)
		return make([]*Deck, 0)
	}

	for rows.Next() {
		deck := Deck{}
		if e = rows.Scan(&(deck.Id), &(deck.Name)); e != nil {
			log.Println(e)
		}

		decks = append(decks, &deck)
	}

	return decks
}

func Preamble(dbo *dbconnector.DBO) *Models {
	return &Models{
		Deck: &Decks{dbo: dbo},
	}
}
