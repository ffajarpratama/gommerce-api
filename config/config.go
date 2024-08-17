package config

import (
	"log"

	"github.com/spf13/viper"
)

var GlobalConfig Config

type Config struct {
	App   App
	JWT   JWT
	MySQL MySQL
}

type App struct {
	Environment string
	Port        int
	URL         string
}

type JWT struct {
	Secret string
}

type MySQL struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	DSN      string
}

func New() *Config {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("[config-file-fail-load] \n", err.Error())
	}

	v := viper.GetViper()
	viper.AutomaticEnv()

	return &Config{
		App: App{
			Environment: v.GetString("APP_ENV"),
			Port:        v.GetInt("APP_PORT"),
			URL:         v.GetString("APP_URL"),
		},
		JWT: JWT{
			Secret: v.GetString("JWT_SECRET"),
		},
		MySQL: MySQL{
			Host:     v.GetString("DB_HOST"),
			Port:     v.GetString("DB_PORT"),
			User:     v.GetString("DB_USER"),
			Password: v.GetString("DB_PASSWORD"),
			Database: v.GetString("DB_DATABASE"),
			DSN:      v.GetString("DB_DSN"),
		},
	}
}
