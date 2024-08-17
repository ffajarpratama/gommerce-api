package repository

import (
	"context"

	"github.com/ffajarpratama/gommerce-api/internal/model"
	"gorm.io/gorm"
)

type IFaceRepository interface {
	// user
	CreateUser(ctx context.Context, data *model.User, db *gorm.DB) error
	FindOneUser(ctx context.Context, query ...interface{}) (*model.User, error)
}
