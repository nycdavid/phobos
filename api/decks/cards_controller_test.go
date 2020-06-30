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
	ts := httptest.NewServer((func() *gin.Engine {
		engine := gin.Default()

		dbo := dbconnector.NewDBO("test")
		models := models.Preamble(dbo)

		api.DrawRoutes(engine, models)

		return engine
	})())
	defer ts.Close()

	resp, e := http.Get("/decks")

	if e != nil {
		t.Errorf("Error: %s", e)
	}

	fmt.Println(resp)
}
