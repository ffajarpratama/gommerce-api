package repository

import (
	"context"

	"github.com/ffajarpratama/gommerce-api/internal/http/request"
	"github.com/ffajarpratama/gommerce-api/internal/model"
	"gorm.io/gorm"
)

type IFaceRepository interface {
	// user
	CreateUser(ctx context.Context, data *model.User, db *gorm.DB) error
	FindOneUser(ctx context.Context, query ...interface{}) (*model.User, error)

	// media
	CreateMedia(ctx context.Context, data *model.Media, db *gorm.DB) error

	// product
	CreateProduct(ctx context.Context, data *model.Product, db *gorm.DB) error
	FindAndCountProduct(ctx context.Context, params *request.ListProductQuery) ([]*model.Product, int64, error)
	FindOneProduct(ctx context.Context, query ...interface{}) (*model.Product, error)
	UpdateProduct(ctx context.Context, db *gorm.DB, data map[string]interface{}, query ...interface{}) error
	DeleteProduct(ctx context.Context, db *gorm.DB, query ...interface{}) error
}
