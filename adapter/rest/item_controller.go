package rest

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/tabakazu/golang-webapi-demo/app/service"
	"github.com/tabakazu/golang-webapi-demo/domain/entity"
	"github.com/tabakazu/golang-webapi-demo/domain/value"
)

type itemController struct {
	UseCase *service.ItemUseCase
}

func NewItemController(uc *service.ItemUseCase) *itemController {
	return &itemController{UseCase: uc}
}

func (ctrl *itemController) List(c echo.Context) error {
	result, err := ctrl.UseCase.GetList.Execute()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

func (ctrl *itemController) Show(c echo.Context) error {
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

func (ctrl *itemController) Create(c echo.Context) error {
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

func (ctrl *itemController) Delete(c echo.Context) error {
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
