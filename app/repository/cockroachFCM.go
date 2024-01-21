package repository

import (
	"cockroachApp/app/model/dto"
	"encoding/json"
	"github.com/gofiber/fiber/v2/log"
)

type CockroachFCMMessaging struct {
}

// create function provider
func NewCockroachFCMMessaging() CockroachMessaging {
	return &CockroachFCMMessaging{}
}

func (c *CockroachFCMMessaging) PushNotification(m *dto.CockroachPushNotificationDto) error {
	dataJson, _ := json.Marshal(m)
	log.Debugf("push fcm notification with data : %v", string(dataJson))
	return nil
}
