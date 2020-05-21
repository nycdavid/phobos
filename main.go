package main

import (
	"github.com/nycdavid/phobos/api"
	"github.com/nycdavid/phobos/web"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	web.Preamble(engine)

	web.DrawRoutes(engine)
	api.DrawRoutes(engine)

	engine.Run()
}
