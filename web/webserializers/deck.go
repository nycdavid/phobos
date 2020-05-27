package webserializers

import (
	"fmt"

	"github.com/nycdavid/phobos/models"
)

const ApiDeckShowRoot = "/api/decks"

type Deck struct {
	Id               int    `json:"id"`
	Name             string `json:"name"`
	ApiDeckPath      string `json:"deckPath"`
	ApiDeckCardsPath string `json:"deckCardsPath"`
}

func NewDeck(deck *models.Deck) *Deck {
	return &Deck{
		Id:               deck.Id,
		Name:             deck.Name,
		ApiDeckPath:      fmt.Sprintf("%s/%d", ApiDeckShowRoot, deck.Id),
		ApiDeckCardsPath: fmt.Sprintf("%s/%d/cards", ApiDeckShowRoot, deck.Id),
	}
}
