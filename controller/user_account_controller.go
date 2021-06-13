package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/tabakazu/go-webapp/application"
)

type UserAccountController interface {
	RegisterHandler(echo.Context) error
}

type userAccountController struct {
	register application.RegisterUserAccount
}

func NewUserAccountController(
	r application.RegisterUserAccount,
) UserAccountController {
	return &userAccountController{
		register: r,
	}
}

func (ctrl *userAccountController) RegisterHandler(c echo.Context) error {
	var param application.RegisterUserAccountParam
	if err := c.Bind(&param); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, 1*time.Second)

	r, err := ctrl.register.Execute(ctx, &param)
	if err != nil {
		return c.String(http.StatusUnauthorized, err.Error())
	}
	return c.JSON(http.StatusOK, r)
}
