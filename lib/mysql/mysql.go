package mysql

import (
	"log"

	"github.com/ffajarpratama/gommerce-api/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewMySQLClient(cnf *config.Config) (*gorm.DB, error) {
	logLevel := logger.Error
	switch cnf.App.Environment {
	case "production":
		logLevel = logger.Error
	case "development", "staging":
		logLevel = logger.Warn
	default:
		logLevel = logger.Info
	}

	gormConfig := gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	}

	conn, err := gorm.Open(mysql.Open(cnf.MySQL.DSN), &gormConfig)
	if err != nil {
		log.Fatal("[mysql-connection-error] \n", err.Error())
		return nil, err
	}

	db, err := conn.DB()
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("[error: db.Ping()] \n", err.Error())
		return nil, err
	}

	log.Println("[mysql-connected]")

	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(100)

	return conn, nil
}
