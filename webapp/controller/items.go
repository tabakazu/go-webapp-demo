package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/tabakazu/golang-webapi-demo/domain/entity"
	"github.com/tabakazu/golang-webapi-demo/domain/value"
	"github.com/tabakazu/golang-webapi-demo/usecase"
)

type Items struct {
	ListAction   usecase.ItemsGet
	ShowAction   usecase.ItemGet
	CreateAction usecase.ItemCreate
	UpdateAction usecase.ItemUpdate
	DeleteAction usecase.ItemDelete
}

type ItemsActions struct {
	List   usecase.ItemsGet
	Show   usecase.ItemGet
	Create usecase.ItemCreate
	Update usecase.ItemUpdate
	Delete usecase.ItemDelete
}

func NewItems(a ItemsActions) Items {
	return Items{
		ListAction:   a.List,
		ShowAction:   a.Show,
		CreateAction: a.Create,
		UpdateAction: a.Update,
		DeleteAction: a.Delete,
	}
}

func (ctrl Items) List(c echo.Context) error {
	result, err := ctrl.ListAction.Execute()
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

	result, err := ctrl.ShowAction.Execute(value.ItemID(id))
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

	result, err := ctrl.CreateAction.Execute(item)
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

	result, err := ctrl.UpdateAction.Execute(value.ItemID(id), item)
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

	result, err := ctrl.DeleteAction.Execute(value.ItemID(id))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}
