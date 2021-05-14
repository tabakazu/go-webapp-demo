package main

import (
	"github.com/tabakazu/golang-webapi-demo/adapter/dbgateway"
	"github.com/tabakazu/golang-webapi-demo/adapter/rest"
	"github.com/tabakazu/golang-webapi-demo/app/service"
	"github.com/tabakazu/golang-webapi-demo/infra/db"
	"github.com/tabakazu/golang-webapi-demo/infra/web"
)

func main() {
	s := InitializeWebServer()
	s.ListenAndServe()
}

func InitializeWebServer() web.Server {
	conn := db.NewConnection()
	srv := web.NewServer()

	healthCheckCtrl := rest.NewHealthCheckController()
	rest.SetupHealthCheckRoute(srv.Router, healthCheckCtrl)

	itemRepo := dbgateway.NewItemRepository(conn)
	itemUseCase := service.NewItemUseCase(itemRepo)
	itemCtrl := rest.NewItemController(itemUseCase)
	rest.SetupItemRoute(srv.Router, itemCtrl)

	return srv
}
