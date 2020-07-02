package api

import (
	"net/http"

	"github.com/nycdavid/phobos/models"

	"github.com/gin-gonic/gin"
)

func DecksCardsController(engine *gin.Engine, models *models.Models) {
	routes := []map[string]interface{}{
		map[string]interface{}{
			"path":   "api/decks/:id/cards",
			"method": "GET",
			"func":   DecksCardsController_Index(engine, models),
		},
	}

	// TODO : extract this to module, we'll need it in a few
	// places
	for _, route := range routes {
		switch route["method"].(string) {
		case "GET":
			engine.GET(
				route["path"].(string),
				route["func"].(func(*gin.Context)),
			)
		case "POST":
			engine.POST(
				route["path"].(string),
				route["func"].(func(*gin.Context)),
			)
		}
	}
}

func DecksCardsController_Index(engine *gin.Engine, m *models.Models) func(*gin.Context) {
	return func(c *gin.Context) {
		allCards := m.Card.All()

		print(allCards)
		c.JSON(http.StatusOK, allCards)
	}
}
