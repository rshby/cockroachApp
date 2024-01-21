package entity

import "time"

type Cockroach struct {
	Id        uint32    `gorm:"column:id;not null;primaryKey;autoIncrement" json:"id,omitempty"`
	Amount    uint32    `gorm:"column:amount;not null" json:"amount,omitempty"`
	CreatedAt time.Time `gorm:"column:created_at;not null;autoCreateTime" json:"created_at,omitempty"`
}

func (c *Cockroach) TableName() string {
	return "cockroach"
}
