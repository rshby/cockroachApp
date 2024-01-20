package config

import (
	"github.com/spf13/viper"
	"log"
)

type ConfigApp struct {
	App      *App      `json:"app,omitempty"`
	Database *Database `json:"database,omitempty"`
}

type App struct {
	Name   string `json:"name,omitempty"`
	Author string `json:"author,omitempty"`
	Port   int    `json:"port,omitempty"`
}

type Database struct {
	Host         string `json:"host,omitempty"`
	User         string `json:"user,omitempty"`
	Password     string `json:"password,omitempty"`
	Port         int    `json:"port,omitempty"`
	DatabaseName string `json:"database_name,omitempty"`
}

func NewConfigApp() *ConfigApp {
	config := LoadConfig()

	return &ConfigApp{
		App: &App{
			Name:   config.GetString("app.name"),
			Author: config.GetString("app.author"),
			Port:   config.GetInt("app.port"),
		},
		Database: &Database{
			Host:         config.GetString("database.host"),
			User:         config.GetString("database.user"),
			Password:     config.GetString("database.password"),
			Port:         config.GetInt("database.port"),
			DatabaseName: config.GetString("database.database_name"),
		},
	}
}

func LoadConfig() *viper.Viper {
	config := viper.New()
	config.SetConfigFile("config.json")
	config.SetConfigType("json")
	config.AddConfigPath("./")

	// load config
	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("cant load config : %v", err.Error())
	}

	// success load config
	return config
}
