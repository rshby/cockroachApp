package main

import (
	"cockroachApp/config"
	"cockroachApp/database"
	"fmt"
)

func main() {
	// declare config app
	cfg := config.NewConfigApp()

	// declare database
	db := database.NewMySqlDatabase(cfg)
	dbConn := db.GetDB()
	fmt.Println(dbConn) // TODO : nanti dihapus
}
