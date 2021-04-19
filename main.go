package main

import (
	"github.com/tabakazu/golang-webapi-demo/controller/api"
	"github.com/tabakazu/golang-webapi-demo/gateway/datastore"
	"github.com/tabakazu/golang-webapi-demo/infrastructure/db"
	"github.com/tabakazu/golang-webapi-demo/infrastructure/web"
	"github.com/tabakazu/golang-webapi-demo/usecase/interactor"
)

func main() {
	conn := db.NewConnection()
	itemRepo := datastore.NewItemRepository(conn)
	itemsGetUseCase := interactor.NewItemsGet(itemRepo)
	itemGetUseCase := interactor.NewItemGet(itemRepo)
	itemCreateUseCase := interactor.NewItemCreate(itemRepo)
	itemUpdateUseCase := interactor.NewItemUpdate(itemRepo)
	itemDeleteUseCase := interactor.NewItemDelete(itemRepo)

	s := web.NewServer()
	r := api.NewRouter(s.Router)
	r.SetupItemRoutes(
		itemsGetUseCase,
		itemGetUseCase,
		itemCreateUseCase,
		itemUpdateUseCase,
		itemDeleteUseCase,
	)

	s.ListenAndServe()
}
