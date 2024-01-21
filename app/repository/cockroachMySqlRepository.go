package repository

import (
	"cockroachApp/app/model/dto"
	"cockroachApp/app/model/entity"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

type CockroachMySqlRepository struct {
	db *gorm.DB
}

// create function provider
func NewCockroachMySqlRepository(db *gorm.DB) CockroachRepository {
	return &CockroachMySqlRepository{db: db}
}

// method to insert data
func (c *CockroachMySqlRepository) InsertCockroachData(in *dto.InsertCockroachDto) error {
	data := entity.Cockroach{
		Amount: in.Amount,
	}

	if err := c.db.Model(&entity.Cockroach{}).Save(&data).Error; err != nil {
		log.Errorf("inser error data cockroach : %v", err.Error())
		return err
	}

	log.Debugf("success insert data cockroach")
	return nil
}
