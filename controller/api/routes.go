package api

import (
	"github.com/labstack/echo/v4"
	"github.com/tabakazu/golang-webapi-demo/usecase"
)

type router struct {
	*echo.Echo
}

func NewRouter(e *echo.Echo) router {
	return router{e}
}

func (e router) SetupItemRoutes(
	itemsGetUseCase usecase.ItemsGet,
	itemGetUseCase usecase.ItemGet,
	itemCreateUseCase usecase.ItemCreate,
	itemUpdateUseCase usecase.ItemUpdate,
	itemDeleteUseCase usecase.ItemDelete,
) {
	ctrl := NewItemsController(
		itemsGetUseCase,
		itemGetUseCase,
		itemCreateUseCase,
		itemUpdateUseCase,
		itemDeleteUseCase,
	)
	e.GET("/items", ctrl.List)
	e.POST("/items", ctrl.Create)
	e.GET("/items/:id", ctrl.Show)
	e.PUT("/items/:id", ctrl.Update)
	e.DELETE("/items/:id", ctrl.Delete)
}
