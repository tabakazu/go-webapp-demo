package web

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tabakazu/golang-webapi-demo/controller"
)

type server struct {
	*echo.Echo
}

func NewServer() *server {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	return &server{e}
}

func (s server) SetupItemRoutes(ctrl controller.Items) {
	s.GET("/items", ctrl.List)
	s.POST("/items", ctrl.Create)
	s.GET("/items/:id", ctrl.Show)
	s.PUT("/items/:id", ctrl.Update)
	s.DELETE("/items/:id", ctrl.Delete)
}

func (s server) ListenAndServe() {
	s.Logger.Fatal(s.Start(":8080"))
}
