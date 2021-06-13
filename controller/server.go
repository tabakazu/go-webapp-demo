package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// type Server interface {
// 	ListenAndServe()
// }

// type Server struct {
// 	*echo.Echo
// }

// func NewServer() *Server {
// 	e := echo.New()
// 	e.Use(middleware.Logger())
// 	e.Use(middleware.Recover())

// 	return &Server{e}
// }

// func (s *Server) ListenAndServe() {
// 	s.Logger.Fatal(s.Start(":8080"))
// }

type Server struct {
	router *echo.Echo
}

func NewServer(
	userAccountCtrl UserAccountController,
) *Server {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/user_account", userAccountCtrl.RegisterHandler)

	return &Server{router: e}
}

func (srv *Server) ListenAndServe() {
	srv.router.Logger.Fatal(srv.router.Start(":8080"))
}
