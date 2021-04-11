package main

import (
	"os"

	"github.com/tabakazu/golang-webapi-demo/application"
	"github.com/tabakazu/golang-webapi-demo/controller"
	"github.com/tabakazu/golang-webapi-demo/db"
	"github.com/tabakazu/golang-webapi-demo/gateway"
	"github.com/tabakazu/golang-webapi-demo/web"
)

func main() {
	dbConn := db.New(os.Getenv("MYSQL_URL")).Connect()
	itemRepo := gateway.NewItemRepository(dbConn)
	itemServices := application.NewItemServices(itemRepo)
	itemsCtrl := controller.NewItems(itemServices)

	s := web.NewServer(web.RoutingSet{
		Items: itemsCtrl,
	})
	s.ListenAndServe()
}
