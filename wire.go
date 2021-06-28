//+build wireinject

package main

import (
	"github.com/google/wire"
	appService "github.com/tabakazu/go-webapp/application/service"
	"github.com/tabakazu/go-webapp/external/datastore"
	datastoreRepo "github.com/tabakazu/go-webapp/external/datastore/repository"
	"github.com/tabakazu/go-webapp/interfaces/webapi"
	webapiCtrl "github.com/tabakazu/go-webapp/interfaces/webapi/controller"
	webapiGen "github.com/tabakazu/go-webapp/interfaces/webapi/generator"
	"gorm.io/gorm"
)

func InitializeDB() (*gorm.DB, func()) {
	dbConfig := datastore.NewDBConfig()
	db, dbClose := datastore.NewConnection(dbConfig)
	return db, dbClose
}

func InitializeServer(dbConn *gorm.DB) *webapi.Server {
	wire.Build(
		datastoreRepo.NewUserAccountRepository,
		webapiGen.NewUserTokenGenerator,
		appService.NewUserAccountRegisterService,
		appService.NewUserAccountLoginService,
		appService.NewUserAccountShowService,
		webapiCtrl.NewUserAccountController,
		webapi.NewServer,
	)
	return &webapi.Server{}
}
