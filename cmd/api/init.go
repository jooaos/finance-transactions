package main

import (
	"log"

	"github.com/jooaos/pismo/config"
	"github.com/jooaos/pismo/internal/controller"
	"github.com/jooaos/pismo/internal/repository/adapter"
	"github.com/jooaos/pismo/internal/service"
)

type Api struct {
	Config      config.Config
	Controllers Controllers
}

type Controllers struct {
	AccountController     controller.IAccountController
	TransactionController controller.ITransactionController
}

func InitDependenciesApi() Api {
	config := config.NewConfig()

	databaseConnection, err := config.Database.OpenDatabaseConnection()
	if err != nil {
		log.Fatal("Error while open database connection", err.Error())
	}

	accountRepository := adapter.NewAccountRepositoryMariaDB(databaseConnection)
	transactionRepository := adapter.NewTransactionRepositoryMariaDB(databaseConnection)

	accountService := service.NewAccountService(accountRepository)
	transactionService := service.NewTransactionService(transactionRepository, accountRepository)

	accountController := controller.NewAccountController(accountService)
	transactionController := controller.NewTransactionController(transactionService)

	return Api{
		Config: config,
		Controllers: Controllers{
			AccountController:     accountController,
			TransactionController: transactionController,
		},
	}
}
