package main

import (
	"cockroachApp/config"
	"cockroachApp/database"
	"encoding/json"
	"fmt"
)

func main() {
	// declare config app
	cfg := config.NewConfigApp()
	cfgJson, _ := json.Marshal(&cfg)
	fmt.Println(string(cfgJson))

	var db database.Database = database.NewMySqlDatabase(cfg)
	dbConn := db.GetDB()
	fmt.Println(dbConn) // TODO : nanti dihapus
}
