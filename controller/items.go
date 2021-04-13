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
	ListService   usecase.ItemList
	ShowService   usecase.ItemShow
	CreateService usecase.ItemCreate
}

func NewItems(s application.ItemServices) Items {
	return Items{
		ListService:   s.ListService,
		ShowService:   s.ShowService,
		CreateService: s.CreateService,
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

func (ctrl Items) Create(c echo.Context) error {
	var item domain.Item
	if err := c.Bind(&item); err != nil {
		return err
	}

	result, err := ctrl.CreateService.Execute(item)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, result)
}
