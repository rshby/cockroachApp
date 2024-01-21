package mock

import (
	"cockroachApp/app/model/dto"
	"github.com/stretchr/testify/mock"
)

type CockroachMessagingMock struct {
	Mock *mock.Mock
}

// function provider
func NewCockroachMessagingMock() *CockroachMessagingMock {
	return &CockroachMessagingMock{Mock: &mock.Mock{}}
}

func (c *CockroachMessagingMock) PushNotification(m *dto.CockroachPushNotificationDto) error {
	args := c.Mock.Called(m)
	return args.Error(0)
}
