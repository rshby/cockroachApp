package mock

import (
	"cockroachApp/app/model/dto"
	"github.com/stretchr/testify/mock"
)

type CockroachRepoMock struct {
	Mock *mock.Mock
}

// function provider
func NewCockroachRepoMock() *CockroachRepoMock {
	return &CockroachRepoMock{Mock: &mock.Mock{}}
}

func (c *CockroachRepoMock) InsertCockroachData(in *dto.InsertCockroachDto) error {
	args := c.Mock.Called(in)
	return args.Error(0)
}
