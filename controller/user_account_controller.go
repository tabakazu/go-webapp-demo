package controller

import (
	"context"
	"net/http"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/tabakazu/go-webapp/application"
)

type UserAccountController interface {
	RegisterHandler(echo.Context) error
	LoginHandler(echo.Context) error
	ShowHandler(echo.Context) error
}

type userAccountController struct {
	register application.RegisterUserAccount
	login    application.LoginUserAccount
	show     application.ShowUserAccount
}

func NewUserAccountController(
	r application.RegisterUserAccount,
	l application.LoginUserAccount,
	s application.ShowUserAccount,
) UserAccountController {
	return &userAccountController{
		register: r,
		login:    l,
		show:     s,
	}
}

// RegisterHandler godoc
// @Summary Register a user
// @Description create user
// @Accept json
// @Produce json
// @Param user body application.RegisterUserAccountParam true "RegisterUserAccountParam"
// @Success 200 {object} application.RegisterUserAccountResult
// @Failure 400
// @Router /user_account [post]
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

// LoginHandler godoc
// @Summary Login with user
// @Description login with user
// @Accept json
// @Produce json
// @Param user body application.LoginUserAccountParam true "LoginUserAccountParam"
// @Success 200 {object} application.LoginUserAccountResult
// @Failure 400,401
// @Router /user_account/login [post]
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

// ShowHandler godoc
// @Summary Show a logged in user
// @Description Show a logged in user
// @Accept json
// @Produce json
// @Success 200 {object} application.ShowUserAccountResult
// @Failure 401
// @Router /user_account [get]
func (ctrl *userAccountController) ShowHandler(c echo.Context) error {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, 1*time.Second)

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID, err := strconv.Atoi(claims["user_id"].(string))
	if err != nil {
		return err
	}

	r, err := ctrl.show.Execute(ctx, userID)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, r)
}
