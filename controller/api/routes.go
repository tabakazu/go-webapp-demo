package api

import (
	"github.com/labstack/echo/v4"
)

type RoutingSet struct {
	Items Items
}

type router struct {
	*echo.Echo
}

func NewRouter(e *echo.Echo) router {
	return router{e}
}

func (e router) SetupRoutes(r RoutingSet) {
	e.GET("/items", r.Items.List)
	e.POST("/items", r.Items.Create)
	e.GET("/items/:id", r.Items.Show)
	e.PUT("/items/:id", r.Items.Update)
	e.DELETE("/items/:id", r.Items.Delete)
}
