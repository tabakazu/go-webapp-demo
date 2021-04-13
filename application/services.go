package application

import (
	"github.com/tabakazu/golang-webapi-demo/application/usecase"
	"github.com/tabakazu/golang-webapi-demo/domain/repository"
)

type ItemServices struct {
	ListService   usecase.ItemList
	ShowService   usecase.ItemShow
	CreateService usecase.ItemCreate
	UpdateService usecase.ItemUpdate
	DeleteService usecase.ItemDelete
}

func NewItemServices(r repository.Item) ItemServices {
	return ItemServices{
		ListService:   NewItemListService(r),
		ShowService:   NewItemShowService(r),
		CreateService: NewItemCreateService(r),
		UpdateService: NewItemUpdateService(r),
		DeleteService: NewItemDeleteService(r),
	}
}
