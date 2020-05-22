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
		map[string]interface{}{
			"path":   "/api/decks/:id",
			"method": "GET",
			"func":   DecksController_Show(engine, models),
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

func DecksController_Show(engine *gin.Engine, models *models.Models) func(*gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")

		deck, e := models.Deck.Find(id)
		if e != nil {
			c.JSON(http.StatusNotFound, nil)
		} else {
			c.JSON(http.StatusOK, deck)
		}
	}
}
