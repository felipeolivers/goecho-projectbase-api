package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.elastic.co/apm/module/apmechov4"
)

func loadMiddleware(e *echo.Echo) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.Secure())
	e.Use(apmechov4.Middleware())
}
