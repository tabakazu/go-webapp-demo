package main

import (
	"log"

	"github.com/tabakazu/golang-webapi-demo/adapter/dbgateway"
	"github.com/tabakazu/golang-webapi-demo/adapter/rest"
	"github.com/tabakazu/golang-webapi-demo/app/service"
	"github.com/tabakazu/golang-webapi-demo/infra/db"
	"github.com/tabakazu/golang-webapi-demo/infra/web"
)

type (
	webServer interface {
		ListenAndServe()
	}
	sqlDB interface {
		Close() error
	}
)

func main() {
	s, c := InitializeWebServer()
	s.ListenAndServe()

	defer func() {
		if err := c.Close(); err != nil {
			log.Fatal(err)
		}
	}()
}

func InitializeWebServer() (webServer, sqlDB) {
	conn := db.NewConnection()
	srv := web.NewServer()

	healthCheckCtrl := rest.NewHealthCheckController()
	rest.SetupHealthCheckRoute(srv.Router, healthCheckCtrl)

	itemRepo := dbgateway.NewItemRepository(conn)
	itemUseCase := service.NewItemUseCase(itemRepo)
	itemCtrl := rest.NewItemController(itemUseCase)
	rest.SetupItemRoute(srv.Router, itemCtrl)

	db, err := conn.DB()
	if err != nil {
		log.Fatal(err)
	}
	return srv, db
}
