package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nycdavid/phobos/api"
	"github.com/nycdavid/phobos/dbconnector"
	"github.com/nycdavid/phobos/models"

	"github.com/gin-gonic/gin"
)

func TestDecksCardsController_Create(t *testing.T) {
	engine := gin.Default()

	dbo := dbconnector.NewDBO("test")
	mdls := models.Preamble(dbo)
	mdls.Deck.Create(map[string]interface{}{
		"front": "foo",
		"back":  "bar",
	})

	ts := httptest.NewServer((func() *gin.Engine {
		api.DrawRoutes(engine, mdls)

		return engine
	})())
	defer ts.Close()

	http.Get(fmt.Sprintf("%s/api/decks/1/cards", ts.URL))
}
