package usecase

import (
	"github.com/ffajarpratama/gommerce-api/config"
	"github.com/ffajarpratama/gommerce-api/internal/repository"
	"gorm.io/gorm"
)

type Usecase struct {
	cnf  *config.Config
	repo repository.IFaceRepository
	db   *gorm.DB
}

func New(cnf *config.Config, repo repository.IFaceRepository, db *gorm.DB) IFaceUsecase {
	return &Usecase{
		cnf:  cnf,
		repo: repo,
		db:   db,
	}
}
