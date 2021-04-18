package main

import (
	"os"

	controller "github.com/tabakazu/golang-webapi-demo/controller/api"
	gateway "github.com/tabakazu/golang-webapi-demo/gateway/datastore"
	"github.com/tabakazu/golang-webapi-demo/infrastructure/db"
	"github.com/tabakazu/golang-webapi-demo/infrastructure/web"
	"github.com/tabakazu/golang-webapi-demo/usecase/interactor"
)

func main() {
	dbConn := db.New(os.Getenv("MYSQL_URL")).Connect()
	itemRepo := gateway.NewItemRepository(dbConn)
	itemsActions := controller.ItemsActions{
		List:   interactor.NewItemsGet(itemRepo),
		Show:   interactor.NewItemGet(itemRepo),
		Create: interactor.NewItemCreate(itemRepo),
		Update: interactor.NewItemUpdate(itemRepo),
		Delete: interactor.NewItemDelete(itemRepo),
	}
	itemsCtrl := controller.NewItems(itemsActions)

	s := web.NewServer()
	r := controller.NewRouter(s.Router)
	r.SetupRoutes(controller.RoutingSet{
		Items: itemsCtrl,
	})

	s.ListenAndServe()
}
