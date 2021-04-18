package web

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type server struct {
	Router *echo.Echo
}

func NewServer() *server {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	return &server{Router: e}
}

func (s server) ListenAndServe() {
	s.Router.Logger.Fatal(s.Router.Start(":8080"))
}
