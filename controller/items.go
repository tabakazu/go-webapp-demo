package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/tabakazu/golang-webapi-demo/application"
	"github.com/tabakazu/golang-webapi-demo/application/usecase"
	"github.com/tabakazu/golang-webapi-demo/domain"
)

type Items struct {
	ListService usecase.ItemList
	ShowService usecase.ItemShow
}

func NewItems(s application.ItemServices) Items {
	return Items{
		ListService: s.ListService,
		ShowService: s.ShowService,
	}
}

func (ctrl Items) List(c echo.Context) error {
	result, err := ctrl.ListService.Execute()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

func (ctrl Items) Show(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	result, err := ctrl.ShowService.Execute(domain.ItemID(id))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}
