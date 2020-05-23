package web

import (
	"net/http"

	"github.com/nycdavid/phobos/models"

	"github.com/gin-gonic/gin"
)

func DecksController(engine *gin.Engine, models *models.Models) {
	routes := []map[string]interface{}{
		map[string]interface{}{
			"path":   "/decks/new",
			"method": "GET",
			"func":   DecksController_New(engine),
		},
		map[string]interface{}{
			"path":   "/decks",
			"method": "GET",
			"func":   DecksController_Index(engine, models),
		},
	}

	for _, route := range routes {
		switch route["method"].(string) {
		case "GET":
			engine.GET(
				route["path"].(string),
				route["func"].(func(*gin.Context)),
			)
		}
	}
}

func DecksController_Index(engine *gin.Engine, m *models.Models) func(*gin.Context) {
	return func(c *gin.Context) {
		decks := m.Deck.All()
		c.HTML(http.StatusOK, "decks_index", gin.H{
			"decks":        decks,
			"showDeckPath": "/api/decks/",
		})
	}
}

func DecksController_New(engine *gin.Engine) func(*gin.Context) {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "decks_new", nil)
	}
}
