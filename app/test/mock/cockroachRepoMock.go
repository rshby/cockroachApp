package mock

import (
	"cockroachApp/app/model/dto"
	"github.com/stretchr/testify/mock"
)

type CockroachRepoMock struct {
	Mock *mock.Mock
}

func (c *CockroachRepoMock) InsertCockroachData(in *dto.InsertCockroachDto) error {
	//TODO implement me
	panic("implement me")
}
