package webapi

import (
	"github.com/labstack/echo/v4"
	"github.com/tabakazu/go-webapp/interfaces/webapi/controller"
)

type customContext struct {
	echo.Context
	session *controller.Session
}

func (c *customContext) Session() *controller.Session {
	return c.session
}
