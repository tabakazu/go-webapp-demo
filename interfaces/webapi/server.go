package webapi

import (
	"os"

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

	authMiddleware := middleware.JWT([]byte(os.Getenv("SECRET_KEY")))
	e.POST("/user_account", func(c echo.Context) error { return userAccountCtrl.RegisterHandler(context{c}) })
	e.POST("/user_account/login", func(c echo.Context) error { return userAccountCtrl.LoginHandler(context{c}) })
	e.GET("/user_account", func(c echo.Context) error { return userAccountCtrl.ShowHandler(context{c}) }, authMiddleware)

	return &Server{router: e}
}

func (srv *Server) ListenAndServe() {
	srv.router.Logger.Fatal(srv.router.Start(":8080"))
}
