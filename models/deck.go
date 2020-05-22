package models

import (
	"fmt"
	"log"

	"github.com/nycdavid/phobos/dbconnector"
)

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

func (d *Decks) Find(id string) (*Deck, error) {
	var deck Deck
	row := d.dbo.Conn.QueryRow(
		fmt.Sprintf("SELECT * FROM decks WHERE decks.id = %s", id),
	)

	e := row.Scan(&(deck.Id), &(deck.Name))
	if e != nil {
		return nil, e
	}

	return &deck, nil
}

func (d *Decks) Create(deck *Deck) (*Deck, error) {
	row := d.dbo.Conn.QueryRow(fmt.Sprintf(
		"INSERT INTO decks (name) VALUES ('%s') RETURNING id",
		deck.Name,
	))

	e := row.Scan(&(deck.Id))
	if e != nil {
		return nil, e
	}

	return deck, nil
}
