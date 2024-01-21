package dto

import "time"

type InsertCockroachDto struct {
	Id        uint32    `json:"id,omitempty"`
	Amount    uint32    `json:"amount,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
