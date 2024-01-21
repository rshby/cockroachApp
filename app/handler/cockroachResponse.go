package handler

import (
	"cockroachApp/app/model/dto"
	"github.com/gofiber/fiber/v2"
)

func Response(ctx *fiber.Ctx, responseCode int, message string) error {
	ctx.Status(responseCode)
	return ctx.JSON(&dto.ApiResponse{
		StatusCode: responseCode,
		Status:     "ok",
		Message:    message,
	})
}
