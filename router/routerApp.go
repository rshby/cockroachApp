package router

import (
	"cockroachApp/app/handler"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func NewRouter(app fiber.Router, handler handler.CockroachHandler) {
	app.Get("/test", func(ctx *fiber.Ctx) error {
		return ctx.JSON(map[string]any{
			"status_code": http.StatusOK,
			"message":     "success test",
		})
	})

	app.Post("/cockroach", handler.DetectCockroach)
}
