package usecase

import (
	"context"

	"github.com/ffajarpratama/gommerce-api/internal/http/request"
	"github.com/ffajarpratama/gommerce-api/internal/model"
	"github.com/google/uuid"
)

type IFaceUsecase interface {
	// auth
	Register(ctx context.Context, req *request.Register) (*model.User, error)
	Login(ctx context.Context, req *request.Login) (*model.User, error)
	GetProfile(ctx context.Context, userID uuid.UUID) (*model.User, error)

	// media
	CreateMedia(ctx context.Context, req *request.CreateMedia) (*model.Media, error)

	// product
	CreateProduct(ctx context.Context, req *request.CreateProduct) error
	FindAndCountProduct(ctx context.Context, params *request.ListProductQuery) ([]*model.Product, int64, error)
	FindOneProduct(ctx context.Context, productID uuid.UUID) (*model.Product, error)
	UpdateProduct(ctx context.Context, req *request.UpdateProduct) error
	DeleteProduct(ctx context.Context, productID uuid.UUID) error
}
