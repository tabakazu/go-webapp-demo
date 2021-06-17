package controller

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

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
	e.POST("/user_account/login", userAccountCtrl.LoginHandler)
	e.GET("/user_account", userAccountCtrl.ShowHandler, middleware.JWT([]byte(os.Getenv("SECRET_KEY"))))

	return &Server{router: e}
}

func (srv *Server) ListenAndServe() {
	srv.router.Logger.Fatal(srv.router.Start(":8080"))
}
