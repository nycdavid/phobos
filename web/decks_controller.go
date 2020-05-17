package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DecksController(engine *gin.Engine) {
	routes := []map[string]interface{}{
		map[string]interface{}{
			"path":   "/decks/new",
			"method": "GET",
			"func":   DecksController_New(engine),
		},
		map[string]interface{}{
			"path":   "/decks",
			"method": "GET",
			"func":   DecksController_Index(engine),
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

func DecksController_Index(engine *gin.Engine) func(*gin.Context) {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "decks_index", gin.H{
			"decks": []map[string]interface{}{
				map[string]interface{}{
					"name": "Korean",
				},
			},
		})
	}
}

func DecksController_New(engine *gin.Engine) func(*gin.Context) {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "decks_new", nil)
	}
}
