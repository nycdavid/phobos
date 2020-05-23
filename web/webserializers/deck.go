package webserializers

import (
	"fmt"

	"github.com/nycdavid/phobos/models"
)

const ApiDeckShowRoot = "/api/decks"

type Deck struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	ApiShowPath string `json:"showPath"`
}

func NewDeck(deck *models.Deck) *Deck {
	return &Deck{
		Id:          deck.Id,
		Name:        deck.Name,
		ApiShowPath: fmt.Sprintf("%s/%d", ApiDeckShowRoot, deck.Id),
	}
}
