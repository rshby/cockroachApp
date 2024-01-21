package repository

import "cockroachApp/app/model/dto"

type CockroachRepository interface {
	InsertCockroachData(in *dto.InsertCockroachDto) error
}
