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

func DecksController_New(engine *gin.Engine) func(*gin.Context) {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "decks_new", nil)
	}
}
