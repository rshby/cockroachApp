package server

import (
	"cockroachApp/app/handler"
	"cockroachApp/config"
	"cockroachApp/database"
	"cockroachApp/router"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"
)

type ServerApp struct {
	App    *fiber.App
	DB     *gorm.DB
	Config *config.ConfigApp
}

func NewServerApp(cfg *config.ConfigApp, db database.Database) Server {
	app := fiber.New(fiber.Config{
		Prefork: false,
	})

	app.Use(logger.New())

	return &ServerApp{
		App:    app,
		DB:     db.GetDB(),
		Config: cfg,
	}
}

func (s *ServerApp) RunServer() error {
	addr := fmt.Sprintf(":%v", s.Config.App.Port)
	return s.App.Listen(addr)
}

func (s *ServerApp) AddRouter(handler handler.CockroachHandler) {
	v1 := s.App.Group("/v1")
	router.NewRouter(v1, handler)
}
