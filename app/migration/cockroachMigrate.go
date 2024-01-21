package main

import (
	"cockroachApp/app/model/entity"
	"cockroachApp/config"
	"cockroachApp/database"
)

func main() {
	// generate config
	cfg := config.NewConfigApp()

	// load database
	db := database.NewMySqlDatabase(cfg)

	cockroachMigrate(db)
}

func cockroachMigrate(db database.Database) {
	db.GetDB().Migrator().CreateTable(&entity.Cockroach{})
	data := []entity.Cockroach{
		{Amount: 1},
		{Amount: 2},
		{Amount: 2},
		{Amount: 5},
		{Amount: 3},
	}

	db.GetDB().CreateInBatches(&data, 10)
}
