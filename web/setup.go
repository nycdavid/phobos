package web

import (
	"net/http"

	"github.com/nycdavid/phobos/models"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

var GetRoutes = map[string]func(c *gin.Context){
	"/": func(c *gin.Context) {
		c.HTML(http.StatusOK, "layouts/application.html", gin.H{
			"title": "Index",
		})
	},
}

func renderer() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("layout", "views/layouts/application.html")
	r.AddFromFiles(
		"decks_index",
		"views/layouts/application.html",
		"views/decks/index.html",
	)
	r.AddFromFiles(
		"decks_new",
		"views/layouts/application.html",
		"views/decks/new.html",
	)

	return r
}

func Preamble(engine *gin.Engine) {
	engine.LoadHTMLGlob("views/**/*")
	engine.HTMLRender = renderer()
	engine.StaticFile("/favicon.ico", "./images/favicon.ico")
}

func DrawRoutes(engine *gin.Engine, models *models.Models) {
	DecksController(engine, models)

	for path, fn := range GetRoutes {
		engine.GET(path, fn)
	}
}
