package server

import "github.com/labstack/echo/v4"

func loadConfig(e *echo.Echo) {
	e.Debug = false
	e.HideBanner = true
	e.IPExtractor = echo.ExtractIPFromXFFHeader()
}
