package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	Routes []map[string]interface{}
}

func NewDecksController() *Controller {
	routes := []map[string]interface{}{
		map[string]interface{}{
			"path":   "/decks/new",
			"method": "GET",
			"func":   DecksController_New,
		},
	}

	&Controller{Routes: routes}
}

func DecksController_New(c *gin.Context) {
	c.HTML(http.StatusOK, "decks/new.tmpl")
}
