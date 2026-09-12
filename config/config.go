package config

import (
	"log"

	"github.com/spf13/viper"
)

var (
	GlobalConfig Config
	CLDCloudName string
)

type Config struct {
	App        App
	JWT        JWT
	MySQL      MySQL
	Cloudinary Cloudinary
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

type Cloudinary struct {
	CloudName string
	APIKey    string
	APISecret string
	RootDir   string
	BaseURL   string
}

func New() *Config {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("[config-file-fail-load] \n", err.Error())
	}

	v := viper.GetViper()
	viper.AutomaticEnv()

	CLDCloudName = v.GetString("CLOUDINARY_NAME")

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
		Cloudinary: Cloudinary{
			CloudName: CLDCloudName,
			APIKey:    v.GetString("CLOUDINARY_API_KEY"),
			APISecret: v.GetString("CLOUDINARY_API_SECRET"),
			RootDir:   v.GetString("CLOUDINARY_ROOT_DIR"),
			BaseURL:   v.GetString("CLOUDINARY_BASE_URL"),
		},
	}
}
