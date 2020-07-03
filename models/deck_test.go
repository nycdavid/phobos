package models

import (
	"fmt"
	"testing"

	"github.com/nycdavid/phobos/dbconnector"
	"github.com/nycdavid/phobos/models"
)

func TestDecks_Create(t *testing.T) {
	dbo := dbconnector.NewDBO("test")
	mods := models.Preamble(dbo)
	defer mods.Deck.DeleteAll()

	createdDeck, e := mods.Deck.Create(map[string]interface{}{
		"name": "Foo",
	})
	if e != nil {
		t.Error(e)
	}

	fmt.Println(mods.Deck.Find(createdDeck.Id))
}
