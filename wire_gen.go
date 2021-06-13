// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/tabakazu/go-webapp/application"
	"github.com/tabakazu/go-webapp/controller"
	"github.com/tabakazu/go-webapp/gateway"
)

// Injectors from wire.go:

func InitializeServer() *controller.Server {
	dbConfig := gateway.NewDBConfig()
	db := gateway.NewDB(dbConfig)
	userAccountRepository := gateway.NewUserAccountRepository(db)
	registerUserAccount := application.NewUserAccountRegisterService(userAccountRepository)
	userAccountController := controller.NewUserAccountController(registerUserAccount)
	server := controller.NewServer(userAccountController)
	return server
}
