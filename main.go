package main

import (
	"cockroachApp/config"
	"cockroachApp/database"
	server "cockroachApp/server"
	"log"
)

func main() {
	// declare config app
	cfg := config.NewConfigApp()

	// declare database
	db := database.NewMySqlDatabase(cfg)

	// register repository

	// register usecase

	// register handler

	// intiate server
	server := server.NewServerApp(cfg, db)
	server.AddRouter()
	if err := server.RunServer(); err != nil {
		log.Fatalf("cant run app : %v", err)
	}
}
