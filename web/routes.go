package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var GetRoutes = map[string]func(c *gin.Context){
	"/": func(c *gin.Context) {
		c.HTML(http.StatusOK, "layouts/application.tmpl", gin.H{
			"title": "Index",
		})
	},
}

func Preamble(engine *gin.Engine) {
	engine.LoadHTMLGlob("views/**/*")

	engine.Static("/assets", "./webpack/dist")
}

func DrawRoutes(engine *gin.Engine) {
	DecksController(engine)

	for path, fn := range GetRoutes {
		engine.GET(path, fn)
	}
}
