package webapi

import (
	"github.com/labstack/echo/v4"
	"github.com/tabakazu/go-webapp/interfaces/webapi/controller"
)

func SetupUserAccountRoutes(e *echo.Echo, userAccountCtrl controller.UserAccountController) {
	e.POST("/user_account", func(c echo.Context) error {
		cc := c.(*customContext)
		return userAccountCtrl.RegisterHandler(cc)
	}, CustomContext)

	e.POST("/user_account/login", func(c echo.Context) error {
		cc := c.(*customContext)
		return userAccountCtrl.LoginHandler(cc)
	}, CustomContext)

	e.GET("/user_account", func(c echo.Context) error {
		cc := c.(*customContext)
		return userAccountCtrl.ShowHandler(cc)
	}, CustomContext, UserTokenAuth)
}
