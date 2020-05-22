package api

import (
	"net/http"

	"github.com/nycdavid/phobos/models"

	"github.com/gin-gonic/gin"
)

func DecksController(engine *gin.Engine, models *models.Models) {
	routes := []map[string]interface{}{
		map[string]interface{}{
			"path":   "/api/decks",
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

func DecksController_Index(engine *gin.Engine, models *models.Models) func(*gin.Context) {
	return func(c *gin.Context) {
		decks := models.Deck.All()
		c.JSON(http.StatusOK, decks)
	}
}
