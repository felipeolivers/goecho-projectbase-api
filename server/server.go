package server

import (
	"github.com/felipeolivers/goecho-projectbase-api/config"
	"github.com/labstack/echo/v4"
)

func Start() {
	e := echo.New()

	loadConfig(e)
	loadMiddleware(e)
	loadRoutes(e)

	e.Logger.Fatal(e.Start(":" + config.Env.AppPortServer))

}
