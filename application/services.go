package application

import (
	"github.com/tabakazu/golang-webapi-demo/application/usecase"
	"github.com/tabakazu/golang-webapi-demo/domain/repository"
)

type ItemServices struct {
	ShowService usecase.ItemShow
}

func NewItemServices(r repository.Item) ItemServices {
	return ItemServices{
		ShowService: NewItemShowService(r),
	}
}
