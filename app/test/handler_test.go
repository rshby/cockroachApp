package test

import (
	"cockroachApp/app/handler"
	"cockroachApp/app/model/dto"
	mck "cockroachApp/app/test/mock"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
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
}
