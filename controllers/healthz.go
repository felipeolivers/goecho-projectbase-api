package controllers

import (
	"net/http"

	"github.com/felipeolivers/goecho-projectbase-api/config"
	"github.com/felipeolivers/goecho-projectbase-api/models"
	"github.com/labstack/echo/v4"
	"go.elastic.co/apm"
)

// @Tags healthz
// @Summary Get Api Info
// @Description Get Api Info
// @Accept json
// @Produce json
// @Success 200 {object} models.Healthz
// @Router /healthz [get]
func Healthz(c echo.Context) error {
	span, _ := apm.StartSpan(c.Request().Context(), "controllers.Healthz", "custom")
	defer span.End()
	info := models.Healthz{
		Name:        config.Env.AppName,
		Version:     config.Env.AppVersion,
		Environment: config.Env.AppEnvironment,
	}
	return c.JSON(http.StatusOK, info)
}
