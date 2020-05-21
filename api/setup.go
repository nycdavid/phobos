package api

import (
	"github.com/gin-gonic/gin"
)

func DrawRoutes(engine *gin.Engine) {
	DecksController(engine)
}
