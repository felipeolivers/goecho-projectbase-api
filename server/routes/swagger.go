package routes

import (
	"github.com/felipeolivers/goecho-projectbase-api/config"
	"github.com/felipeolivers/goecho-projectbase-api/docs"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @description Copyright © 2021, Luiz Filipe Miranda de Oliveira.
// @description Todos os direitos reservados.
func Swagger(e *echo.Echo) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Swagger Project Base API
	docs.SwaggerInfo.Title = config.Env.AppName
	docs.SwaggerInfo.Version = config.Env.AppVersion
}
