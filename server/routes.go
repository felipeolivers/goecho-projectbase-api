package server

import (
	"github.com/felipeolivers/goecho-projectbase-api/server/routes"
	"github.com/labstack/echo/v4"
)

func loadRoutes(e *echo.Echo) {
	routes.Healthz(e)
	routes.Swagger(e)
}
