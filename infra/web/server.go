package web

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server interface {
	ListenAndServe()
}

type server struct {
	Router *echo.Echo
}

func NewServer() *server {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	return &server{Router: e}
}

func (s *server) ListenAndServe() {
	e := s.Router
	e.Logger.Fatal(e.Start(":8080"))
}

func (s *server) Routing() *echo.Echo {
	return s.Router
}
