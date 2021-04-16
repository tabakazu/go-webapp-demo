package main

import (
	"os"

	"github.com/tabakazu/golang-webapi-demo/db"
	"github.com/tabakazu/golang-webapi-demo/db/gateway"
	"github.com/tabakazu/golang-webapi-demo/usecase/interactor"
	"github.com/tabakazu/golang-webapi-demo/webapp"
	"github.com/tabakazu/golang-webapi-demo/webapp/controller"
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

	s := webapp.NewServer(webapp.RoutingSet{
		Items: itemsCtrl,
	})
	s.ListenAndServe()
}
