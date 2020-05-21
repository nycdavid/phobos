package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DecksController(engine *gin.Engine) {
	routes := []map[string]interface{}{
		map[string]interface{}{
			"path":   "/api/decks",
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
		c.JSON(http.StatusOK, map[string]string{"foo": "bar"})
	}
}
