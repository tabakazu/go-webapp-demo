package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type healthCheckController struct{}

func NewHealthCheckController() *healthCheckController {
	return &healthCheckController{}
}

func (ctrl *healthCheckController) Show(c echo.Context) error {
	return c.String(http.StatusOK, "OK!")
}
