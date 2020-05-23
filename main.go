package main

import (
	"github.com/nycdavid/phobos/api"
	"github.com/nycdavid/phobos/dbconnector"
	"github.com/nycdavid/phobos/models"
	"github.com/nycdavid/phobos/web"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	dbo := dbconnector.NewDBO("development")
	models := models.Preamble(dbo)

	web.Preamble(engine)
	web.DrawRoutes(engine, models)

	api.DrawRoutes(engine, models)

	engine.Run()
}
