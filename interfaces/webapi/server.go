package webapi

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tabakazu/go-webapp/interfaces/webapi/controller"
)

type Server struct {
	router *echo.Echo
}

// NewServer returns initialized Server
// @title WebApp API Doc
// @version 1.0
// @host localhost:8080
func NewServer(
	userAccountCtrl controller.UserAccountController,
) *Server {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.HTTPErrorHandler = CustomHTTPErrorHandler

	SetupUserAccountRoutes(e, userAccountCtrl)

	return &Server{router: e}
}

func (srv *Server) ListenAndServe() {
	srv.router.Logger.Fatal(srv.router.Start(":8080"))
}
