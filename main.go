package main

import (
	"github.com/tabakazu/golang-webapi-demo/controller"
	"github.com/tabakazu/golang-webapi-demo/gateway"
	"github.com/tabakazu/golang-webapi-demo/infrastructure/db"
	"github.com/tabakazu/golang-webapi-demo/infrastructure/web"
	"github.com/tabakazu/golang-webapi-demo/usecase/interactor"
)

func main() {
	conn := db.NewConnection()
	itemRepo := gateway.NewItemRepository(conn)
	itemUseCases := interactor.NewItemsUseCases(itemRepo)
	itemsCtrl := controller.NewItems(itemUseCases)
	userRepo := gateway.NewUserRepository(conn)
	userUseCases := interactor.NewUserUseCases(userRepo)
	userCtrl := controller.NewUser(userUseCases)

	s := web.NewServer()
	s.SetupItemRoutes(itemsCtrl)
	s.SetupUserRoutes(userCtrl)
	s.ListenAndServe()
}
