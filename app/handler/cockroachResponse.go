package handler

import (
	"cockroachApp/app/model/dto"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func Response(ctx *fiber.Ctx, responseCode int, message string) error {
	ctx.Status(responseCode)
	return ctx.JSON(&dto.ApiResponse{
		StatusCode: responseCode,
		Status:     ConvertStatusFromCode(responseCode),
		Message:    message,
	})
}

// method convert status_code to status (text)
func ConvertStatusFromCode(statusCode int) string {
	switch statusCode {
	case http.StatusOK:
		return "ok"
	case http.StatusBadRequest:
		return "bad request"
	default:
		return "internal server error"
	}
}
