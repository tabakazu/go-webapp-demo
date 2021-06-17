//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/tabakazu/go-webapp/application"
	"github.com/tabakazu/go-webapp/controller"
	"github.com/tabakazu/go-webapp/gateway"
)

func InitializeServer() *controller.Server {
	wire.Build(
		gateway.NewDBConfig,
		gateway.NewDB,
		gateway.NewUserAccountRepository,
		application.NewUserAccountRegisterService,
		application.NewUserAccountLoginService,
		application.NewUserAccountShowService,
		controller.NewUserAccountController,
		controller.NewServer,
	)
	return &controller.Server{}
}
