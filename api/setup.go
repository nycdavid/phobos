package api

import (
	"github.com/nycdavid/phobos/models"

	"github.com/gin-gonic/gin"
)

func DrawRoutes(engine *gin.Engine, models *models.Models) {
	DecksController(engine, models)
	DecksCardsController(engine, models)
}
