package test

import (
	"cockroachApp/app/handler"
	"cockroachApp/app/model/dto"
	mck "cockroachApp/app/test/mock"
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	// test success insert data
	t.Run("test success insert", func(t *testing.T) {
		usecase := mck.NewCockroachUseCaseMock()
		handler := handler.NewCockroachHttpHandler(usecase)

		// create fiber
		fiber := fiber.New()
		fiber.Post("/", handler.DetectCockroach)

		// create data
		request := dto.AddCockroachData{Amount: 1}
		reqJson, _ := json.Marshal(&request)

		// mock
		usecase.Mock.On("CockroachDataProcessing", &request).Return(nil).Times(1)

		// create http request
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(reqJson)))
		req.Header.Add("content-type", "application/json")

		// test
		response, err := fiber.Test(req)
		assert.NotNil(t, response)
		assert.Nil(t, err)

		// get response
		body, err := io.ReadAll(response.Body)
		assert.Nil(t, err)

		responseBody := map[string]any{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, http.StatusOK, int(responseBody["status_code"].(float64)))
		assert.Equal(t, "ok", responseBody["status"].(string))
	})

	// test error failed bad request
	t.Run("test error bad request", func(t *testing.T) {
		usecase := mck.NewCockroachUseCaseMock()
		handler := handler.NewCockroachHttpHandler(usecase)

		// add fiber
		app := fiber.New()
		app.Post("/", handler.DetectCockroach)

		// add http request
		request := httptest.NewRequest(http.MethodPost, "/", nil)
		request.Header.Add("Content-Type", "application/json")

		// hit and receive response
		response, err := app.Test(request)
		assert.Nil(t, err)

		// get response body
		body, _ := io.ReadAll(response.Body)
		responseBody := map[string]any{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, http.StatusBadRequest, int(responseBody["status_code"].(float64)))
		assert.Equal(t, "bad request", responseBody["status"].(string))
	})

	// test error failed internal server error
	t.Run("test error internal server error", func(t *testing.T) {
		usecase := mck.NewCockroachUseCaseMock()
		handler := handler.NewCockroachHttpHandler(usecase)

		// create fiber router
		app := fiber.New()
		app.Post("/", handler.DetectCockroach)

		// create request data
		req := dto.AddCockroachData{Amount: 1}
		reqJson, _ := json.Marshal(&req)
		requestBody := strings.NewReader(string(reqJson))

		// mock
		errorMessage := "failed cant process"
		usecase.Mock.On("CockroachDataProcessing", &req).Return(errors.New(errorMessage)).Times(1)

		// create http request
		request := httptest.NewRequest(http.MethodPost, "/", requestBody)
		request.Header.Add("Content-Type", "application/json")

		// hit and receive response
		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, http.StatusInternalServerError, response.StatusCode)

		// get response body
		body, _ := io.ReadAll(response.Body)
		responseBody := map[string]any{}
		json.Unmarshal(body, &responseBody)

		assert.Equal(t, http.StatusInternalServerError, int(responseBody["status_code"].(float64)))
		assert.Equal(t, "internal server error", responseBody["status"].(string))
		assert.Equal(t, errorMessage, responseBody["message"].(string))
	})
}
