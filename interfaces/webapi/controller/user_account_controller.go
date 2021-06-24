package controller

import (
	"net/http"

	"github.com/tabakazu/go-webapp/application"
	"github.com/tabakazu/go-webapp/application/data"
)

type UserAccountController interface {
	RegisterHandler(Context) error
	LoginHandler(Context) error
	ShowHandler(Context) error
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
// @Param user body data.RegisterUserAccountParam true "RegisterUserAccountParam"
// @Success 200 {object} data.UserAccountResult
// @Failure 400,422
// @Router /user_account [post]
func (ctrl *userAccountController) RegisterHandler(c Context) error {
	var param data.RegisterUserAccountParam
	if err := c.Bind(&param); err != nil {
		return c.JSON(http.StatusBadRequest, &data.ErrorResult{Message: err.Error()})
	}

	if err := data.ValidateRegisterUserAccountParam(&param); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, data.NewValidationErrorResult(err))
	}

	r, err := ctrl.register.Execute(c.Request().Context(), &param)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &data.ErrorResult{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, r)
}

// LoginHandler godoc
// @Summary Login with user
// @Description login with user
// @Accept json
// @Produce json
// @Param user body data.LoginUserAccountParam true "LoginUserAccountParam"
// @Success 200 {object} data.LoginResult
// @Failure 400,401,422
// @Router /user_account/login [post]
func (ctrl *userAccountController) LoginHandler(c Context) error {
	var param data.LoginUserAccountParam
	if err := c.Bind(&param); err != nil {
		return c.JSON(http.StatusBadRequest, &data.ErrorResult{Message: err.Error()})
	}

	if err := data.ValidateLoginUserAccountParam(&param); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, data.NewValidationErrorResult(err))
	}

	r, err := ctrl.login.Execute(c.Request().Context(), &param)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, &data.ErrorResult{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, r)
}

// ShowHandler godoc
// @Summary Show a logged in user
// @Description Show a logged in user
// @Accept json
// @Produce json
// @Success 200 {object} data.UserAccountResult
// @Failure 401,404
// @Router /user_account [get]
func (ctrl *userAccountController) ShowHandler(c Context) error {
	if c.Session() == nil {
		return c.JSON(http.StatusUnauthorized, &data.ErrorResult{Message: "Your session is invalid."})
	}

	r, err := ctrl.show.Execute(c.Request().Context(), c.Session().UserID)
	if err != nil {
		return c.JSON(http.StatusNotFound, &data.ErrorResult{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, r)
}
