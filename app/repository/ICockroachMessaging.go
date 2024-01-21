package repository

import "cockroachApp/app/model/dto"

type CockroachMessaging interface {
	PushNotification(m *dto.CockroachPushNotificationDto) error
}
