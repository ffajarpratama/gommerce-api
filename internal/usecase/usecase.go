package usecase

import (
	"github.com/ffajarpratama/gommerce-api/config"
	"github.com/ffajarpratama/gommerce-api/internal/repository"
	"github.com/ffajarpratama/gommerce-api/lib/cloudinary"
	"gorm.io/gorm"
)

type Usecase struct {
	cnf  *config.Config
	repo repository.IFaceRepository
	db   *gorm.DB
	cld  *cloudinary.Cloudinary
}

func New(cnf *config.Config, repo repository.IFaceRepository, db *gorm.DB, cld *cloudinary.Cloudinary) IFaceUsecase {
	return &Usecase{
		cnf:  cnf,
		repo: repo,
		db:   db,
		cld:  cld,
	}
}
