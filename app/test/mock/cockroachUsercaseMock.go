package mock

import (
	"cockroachApp/app/model/dto"
	"github.com/stretchr/testify/mock"
)

type CockroachUseCaseMock struct {
	Mock *mock.Mock
}

// function provider
func NewCockroachUseCaseMock() *CockroachUseCaseMock {
	return &CockroachUseCaseMock{Mock: &mock.Mock{}}
}

// method implement mock
func (c *CockroachUseCaseMock) CockroachDataProcessing(in *dto.AddCockroachData) error {
	args := c.Mock.Called(in)
	return args.Error(0)
}
