package rest

import (
	"github.com/labstack/echo/v4"
)

func SetupHealthCheckRoute(r *echo.Echo, ctrl *healthCheckController) {
	r.GET("/health_check", ctrl.Show)
}

func SetupItemRoute(r *echo.Echo, ctrl *itemController) {
	r.GET("/items", ctrl.List)
	r.GET("/items/:id", ctrl.Show)
	r.POST("/items", ctrl.Create)
	r.DELETE("/items/:id", ctrl.Delete)
}
