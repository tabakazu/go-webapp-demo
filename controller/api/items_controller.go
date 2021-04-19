package api

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/tabakazu/golang-webapi-demo/domain/entity"
	"github.com/tabakazu/golang-webapi-demo/domain/value"
	"github.com/tabakazu/golang-webapi-demo/usecase"
)

type ItemsController struct {
	GetListUseCase usecase.ItemsGet
	GetUseCase     usecase.ItemGet
	CreateUseCase  usecase.ItemCreate
	UpdateUseCase  usecase.ItemUpdate
	DeleteUseCase  usecase.ItemDelete
}

func NewItemsController(
	getListUseCase usecase.ItemsGet,
	getUseCase usecase.ItemGet,
	createUseCase usecase.ItemCreate,
	updateUseCase usecase.ItemUpdate,
	deleteUseCase usecase.ItemDelete,
) ItemsController {
	return ItemsController{
		GetListUseCase: getListUseCase,
		GetUseCase:     getUseCase,
		CreateUseCase:  createUseCase,
		UpdateUseCase:  updateUseCase,
		DeleteUseCase:  deleteUseCase,
	}
}

func (ctrl ItemsController) List(c echo.Context) error {
	result, err := ctrl.GetListUseCase.Execute()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

func (ctrl ItemsController) Show(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	result, err := ctrl.GetUseCase.Execute(value.ItemID(id))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

func (ctrl ItemsController) Create(c echo.Context) error {
	var item entity.Item
	if err := c.Bind(&item); err != nil {
		return err
	}

	result, err := ctrl.CreateUseCase.Execute(item)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, result)
}

func (ctrl ItemsController) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	var item entity.Item
	if err := c.Bind(&item); err != nil {
		return err
	}

	result, err := ctrl.UpdateUseCase.Execute(value.ItemID(id), item)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, result)
}

func (ctrl ItemsController) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	result, err := ctrl.DeleteUseCase.Execute(value.ItemID(id))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}
