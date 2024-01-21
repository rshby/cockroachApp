package usecase

import (
	"cockroachApp/app/model/dto"
	"cockroachApp/app/repository"
	"time"
)

type CockroachUsecaseImpl struct {
	cockroachRepository repository.CockroachRepository
	cockroachMessaging  repository.CockroachMessaging
}

// create function provider
func NewCockroachUsecaseImpl(cockroachRepository repository.CockroachRepository, cockroachMessaging repository.CockroachMessaging) CockroachUsecase {
	return &CockroachUsecaseImpl{
		cockroachRepository: cockroachRepository,
		cockroachMessaging:  cockroachMessaging,
	}
}

// method process insert
func (c *CockroachUsecaseImpl) CockroachDataProcessing(in *dto.AddCockroachData) error {
	// create data request
	request := dto.InsertCockroachDto{
		Amount: in.Amount,
	}

	// insert to database
	err := c.cockroachRepository.InsertCockroachData(&request)
	if err != nil {
		return err
	}

	// push notif
	pushCockroachData := dto.CockroachPushNotificationDto{
		Title:        "Cockroach Detected!!!",
		Amount:       in.Amount,
		ReportedTime: time.Now().Local().Format("2006-01-02 15:04:05"),
	}

	if err := c.cockroachMessaging.PushNotification(&pushCockroachData); err != nil {
		return err
	}

	// success
	return nil
}
