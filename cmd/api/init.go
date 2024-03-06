package main

import (
	"log"

	"github.com/jooaos/pismo/config"
	"gorm.io/gorm"
)

type Api struct {
	Config config.Config
	// TODO(joao.soares) Remove after tests
	DbConnection *gorm.DB
}

func InitDependenciesApi() Api {
	config := config.NewConfig()

	databaseConnection, err := config.Database.OpenDatabaseConnection()
	if err != nil {
		log.Fatal("Error while open database connection", err.Error())
	}

	return Api{
		Config:       config,
		DbConnection: databaseConnection,
	}
}
