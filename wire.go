//+build wireinject

package main

import (
	"github.com/google/wire"
	appService "github.com/tabakazu/go-webapp/application/service"
	"github.com/tabakazu/go-webapp/external/datastore"
	datastoreRepo "github.com/tabakazu/go-webapp/external/datastore/repository"
	"github.com/tabakazu/go-webapp/interfaces/webapi"
	webapiCtrl "github.com/tabakazu/go-webapp/interfaces/webapi/controller"
)

func InitializeServer() *webapi.Server {
	wire.Build(
		datastore.NewDBConfig,
		datastore.NewConnection,
		datastoreRepo.NewUserAccountRepository,
		appService.NewUserAccountRegisterService,
		appService.NewUserAccountLoginService,
		appService.NewUserAccountShowService,
		webapiCtrl.NewUserAccountController,
		webapi.NewServer,
	)
	return &webapi.Server{}
}
