package integration

import (
	"fmt"

	"github.com/jooaos/pismo/config"
	"gorm.io/gorm"
)

var BaseIntegrationTest *IntegrationTest

type IntegrationTest struct {
	DB *gorm.DB
}

func NewIntegraionTest() {
	config := config.NewConfig()
	db, _ := config.Database.OpenDatabaseConnection()
	BaseIntegrationTest = &IntegrationTest{
		DB: db,
	}
}

func SetUp() {
	if BaseIntegrationTest == nil {
		NewIntegraionTest()
	}
}

func SetDown() {
	databasesToDrop := []string{
		"transactions",
		"accounts",
	}
	for _, item := range databasesToDrop {
		BaseIntegrationTest.DB.Exec(fmt.Sprintf("DELETE FROM %s", item))
	}
}
