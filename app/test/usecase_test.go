package test

import (
	"cockroachApp/app/model/dto"
	mck "cockroachApp/app/test/mock"
	"cockroachApp/app/usecase"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestProcessingUsecase(t *testing.T) {
	cockroackRepo := mck.NewCockroachRepoMock()
	cockroachMessaging := mck.NewCockroachMessagingMock()
	cockroachUsecase := usecase.NewCockroachUsecaseImpl(cockroackRepo, cockroachMessaging)

	// test insert data success
	t.Run("test processing insert data", func(t *testing.T) {
		// create request data
		request := dto.AddCockroachData{Amount: 1}

		data := dto.InsertCockroachDto{
			Amount: request.Amount,
		}

		pushData := dto.CockroachPushNotificationDto{
			Title:        "Cockroach Detected!!!",
			Amount:       request.Amount,
			ReportedTime: time.Now().Local().Format("2006-01-02 15:04:05"),
		}

		// mock repo
		cockroackRepo.Mock.On("InsertCockroachData", &data).Return(nil).Times(1)

		// mock messaging
		cockroachMessaging.Mock.On("PushNotification", &pushData).Return(nil).Times(1)

		// test usercase
		err := cockroachUsecase.CockroachDataProcessing(&request)
		assert.Nil(t, err)
	})

	// test error failed to insert
	t.Run("test failed to insert", func(t *testing.T) {
		// create required data
		request := dto.AddCockroachData{Amount: 1}
		insertData := dto.InsertCockroachDto{
			Amount: request.Amount,
		}

		// mock
		errorMessage := "failed to insert data"
		cockroackRepo.Mock.On("InsertCockroachData", &insertData).Return(errors.New(errorMessage))

		// test
		err := cockroachUsecase.CockroachDataProcessing(&request)
		assert.NotNil(t, err)
		assert.Equal(t, errorMessage, err.Error())
	})

	// test error failed to push
	t.Run("test failed to push data", func(t *testing.T) {
		// create data
		request := dto.AddCockroachData{Amount: 1}

		insertData := dto.InsertCockroachDto{
			Amount: request.Amount,
		}

		// mock
		cockroackRepo.Mock.On("InsertCockroachData", &insertData).Return(nil).Times(1)

		errorMessage := "failed to push"
		cockroachMessaging.Mock.On("PushNotification", mock.Anything).Return(errors.New(errorMessage)).Times(1)

		// test
		err := cockroachUsecase.CockroachDataProcessing(&request)
		assert.NotNil(t, err)
		assert.Equal(t, errorMessage, err.Error())
		cockroackRepo.Mock.AssertExpectations(t)
		cockroachMessaging.Mock.AssertExpectations(t)
	})
}
