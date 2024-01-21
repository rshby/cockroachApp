package database

import (
	"cockroachApp/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

type MySqlDB struct {
	DB *gorm.DB
}

// function provider
func NewMySqlDatabase(cfg *config.ConfigApp) Database {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.DatabaseName)

	log.Println(dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("cant connect db : %v", err)
	}
	return &MySqlDB{db}
}

func (m *MySqlDB) GetDB() *gorm.DB {
	return m.DB
}
