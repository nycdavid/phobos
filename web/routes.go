package web

import (
	"net/http"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

var GetRoutes = map[string]func(c *gin.Context){
	"/": func(c *gin.Context) {
		c.HTML(http.StatusOK, "layouts/application.tmpl", gin.H{
			"title": "Index",
		})
	},
}

func renderer() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("decks_new", "views/layouts/application.tmpl", "views/decks/new.tmpl")

	return r
}

func Preamble(engine *gin.Engine) {
	engine.LoadHTMLGlob("views/**/*")
	engine.HTMLRender = renderer()

	engine.Static("/assets", "./webpack/dist")
}

func DrawRoutes(engine *gin.Engine) {
	DecksController(engine)

	for path, fn := range GetRoutes {
		engine.GET(path, fn)
	}
}
