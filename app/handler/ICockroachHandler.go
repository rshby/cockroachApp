package handler

import "github.com/gofiber/fiber/v2"

type CockroachHandler interface {
	DetectCockroach(ctx *fiber.Ctx) error
}
