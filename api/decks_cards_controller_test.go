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
	models.Cards.Create(&models.Card{Front: "foo"})

	ts := httptest.NewServer((func() *gin.Engine {
		engine := gin.Default()

		dbo := dbconnector.NewDBO("test")
		models := models.Preamble(dbo)

		api.DrawRoutes(engine, models)

		return engine
	})())
	defer ts.Close()

	http.Get(fmt.Sprintf("%s/api/decks/1/cards", ts.URL))
}
