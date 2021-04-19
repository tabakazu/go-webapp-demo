package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/tabakazu/golang-webapi-demo/domain/entity"
	"github.com/tabakazu/golang-webapi-demo/domain/value"
	"github.com/tabakazu/golang-webapi-demo/usecase/interactor"
)

type Items struct {
	UseCase interactor.ItemsUseCases
}

func NewItems(uc interactor.ItemsUseCases) Items {
	return Items{uc}
}

func (ctrl Items) List(c echo.Context) error {
	result, err := ctrl.UseCase.GetList.Execute()
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

	result, err := ctrl.UseCase.Get.Execute(value.ItemID(id))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

func (ctrl Items) Create(c echo.Context) error {
	var item entity.Item
	if err := c.Bind(&item); err != nil {
		return err
	}

	result, err := ctrl.UseCase.Create.Execute(item)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, result)
}

func (ctrl Items) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	var item entity.Item
	if err := c.Bind(&item); err != nil {
		return err
	}

	result, err := ctrl.UseCase.Update.Execute(value.ItemID(id), item)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, result)
}

func (ctrl Items) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	result, err := ctrl.UseCase.Delete.Execute(value.ItemID(id))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}
