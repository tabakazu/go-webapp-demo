package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tabakazu/golang-webapi-demo/domain/entity"
	"github.com/tabakazu/golang-webapi-demo/usecase/interactor"
)

type User struct {
	UseCase interactor.UserUseCases
}

func NewUser(uc interactor.UserUseCases) User {
	return User{uc}
}

func (ctrl User) Register(c echo.Context) error {
	var user entity.User
	if err := c.Bind(&user); err != nil {
		return err
	}

	result, err := ctrl.UseCase.Register.Execute(user)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, result)
}

func (ctrl User) Login(c echo.Context) error {
	var user entity.User
	if err := c.Bind(&user); err != nil {
		return err
	}

	result, err := ctrl.UseCase.Authenticate.Execute(user.Email, user.Password)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}
