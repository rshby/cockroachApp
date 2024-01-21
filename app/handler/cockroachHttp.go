package handler

import (
	"cockroachApp/app/model/dto"
	"cockroachApp/app/usecase"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type CockroachHttpHandler struct {
	cockroachUsecase usecase.CockroachUsecase
}

// create function provider
func NewCockroachHttpHandler(cockroachUsercase usecase.CockroachUsecase) CockroachHandler {
	return &CockroachHttpHandler{
		cockroachUsecase: cockroachUsercase,
	}
}

// method handler
func (c *CockroachHttpHandler) DetectCockroach(ctx *fiber.Ctx) error {
	request := dto.AddCockroachData{}
	err := ctx.BodyParser(&request)
	if err != nil {
		return Response(ctx, http.StatusBadRequest, err.Error())
	}

	// process data
	if err := c.cockroachUsecase.CockroachDataProcessing(&request); err != nil {
		return Response(ctx, http.StatusInternalServerError, err.Error())
	}

	// success
	return Response(ctx, http.StatusOK, "success send data")
}
