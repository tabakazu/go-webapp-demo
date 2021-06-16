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
	LoginHandler(echo.Context) error
}

type userAccountController struct {
	register application.RegisterUserAccount
	login    application.LoginUserAccount
}

func NewUserAccountController(
	r application.RegisterUserAccount,
	l application.LoginUserAccount,
) UserAccountController {
	return &userAccountController{
		register: r,
		login:    l,
	}
}

func (ctrl *userAccountController) RegisterHandler(c echo.Context) error {
	var param application.RegisterUserAccountParam
	if err := c.Bind(&param); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := application.ValidateRegisterUserAccountParam(&param); err != nil {
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

func (ctrl *userAccountController) LoginHandler(c echo.Context) error {
	var param application.LoginUserAccountParam
	if err := c.Bind(&param); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := application.ValidateLoginUserAccountParam(&param); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, 1*time.Second)

	r, err := ctrl.login.Execute(ctx, &param)
	if err != nil {
		return c.String(http.StatusUnauthorized, err.Error())
	}

	return c.JSON(http.StatusOK, r)
}
