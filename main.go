package main

import (
	"cockroachApp/app/handler"
	"cockroachApp/app/repository"
	"cockroachApp/app/usecase"
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
	cockroachRepo := repository.NewCockroachMySqlRepository(db.GetDB())
	messagingRepo := repository.NewCockroachFCMMessaging()

	// register usecase
	cockroachUsecase := usecase.NewCockroachUsecaseImpl(cockroachRepo, messagingRepo)

	// register handler
	cockroachHandler := handler.NewCockroachHttpHandler(cockroachUsecase)

	// intiate server
	server := server.NewServerApp(cfg, db)
	server.AddRouter(cockroachHandler)
	if err := server.RunServer(); err != nil {
		log.Fatalf("cant run app : %v", err)
	}
}
