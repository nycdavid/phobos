package api

import (
	"encoding/json"
	"log"
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
			"path":   "/api/decks",
			"method": "POST",
			"func":   DecksController_Create(engine, models),
		},
		map[string]interface{}{
			"path":   "/api/decks/:id",
			"method": "GET",
			"func":   DecksController_Show(engine, models),
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

func DecksController_Create(engine *gin.Engine, m *models.Models) func(*gin.Context) {
	return func(c *gin.Context) {
		dec := json.NewDecoder(c.Request.Body)
		var deck models.Deck
		e := dec.Decode(&deck)
		if e != nil {
			c.JSON(http.StatusInternalServerError, nil)
		} else {
			createdDeck, e := m.Deck.Create(&deck)

			if e != nil {
				log.Println(e)
				c.JSON(http.StatusInternalServerError, nil)
			} else {
				c.JSON(http.StatusOK, createdDeck)
			}
		}
	}
}
