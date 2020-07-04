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

func (d *Decks) Find(id int) (*Deck, error) {
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

func (d *Decks) Create(data map[string]interface{}) (*Deck, error) {
	row := d.dbo.Conn.QueryRow(fmt.Sprintf(`
		INSERT INTO decks (name)
		VALUES ('%s') RETURNING id, name`,
		data["name"].(string),
	))

	var deck Deck
	e := row.Scan(&(deck.Id), &(deck.Name))
	if e != nil {
		return nil, e
	}

	return &deck, nil
}

func (d *Decks) DeleteAll() error {
	_, err := d.dbo.Conn.Query("TRUNCATE decks;")
	if err != nil {
		return err
	}

	return nil
}
