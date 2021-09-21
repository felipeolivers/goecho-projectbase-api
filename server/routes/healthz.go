package routes

import (
	"github.com/felipeolivers/goecho-projectbase-api/controllers"
	"github.com/labstack/echo/v4"
)

func Healthz(e *echo.Echo) {
	g := e.Group("/healthz")
	g.GET("", controllers.Healthz)
}
