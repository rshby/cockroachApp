package usecase

import "cockroachApp/app/model/dto"

type CockroachUsecase interface {
	CockroachDataProcessing(in *dto.AddCockroachData) error
}
