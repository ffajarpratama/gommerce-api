package usecase

import (
	"context"

	"github.com/ffajarpratama/gommerce-api/internal/http/request"
	"github.com/ffajarpratama/gommerce-api/internal/model"
	"github.com/google/uuid"
)

// CreateProduct implements IFaceUsecase.
func (u *Usecase) CreateProduct(ctx context.Context, req *request.CreateProduct) error {
	product := &model.Product{
		Name:         req.Name,
		Description:  req.Description,
		CategoryName: req.CategoryName,
		ThumbnailID:  req.ThumbnailID,
		Stock:        req.Stock,
		Price:        req.Price,
		SalePrice:    req.SalePrice,
	}

	err := u.repo.CreateProduct(ctx, product, u.db)
	if err != nil {
		return err
	}

	return nil
}

// FindAndCountProduct implements IFaceUsecase.
func (u *Usecase) FindAndCountProduct(ctx context.Context, params *request.ListProductQuery) ([]*model.Product, int64, error) {
	return u.repo.FindAndCountProduct(ctx, params)
}

// FindOneProduct implements IFaceUsecase.
func (u *Usecase) FindOneProduct(ctx context.Context, productID uuid.UUID) (*model.Product, error) {
	return u.repo.FindOneProduct(ctx, "product_id = ?", productID)
}

// UpdateProduct implements IFaceUsecase.
func (u *Usecase) UpdateProduct(ctx context.Context, req *request.UpdateProduct) error {
	data := map[string]interface{}{
		"name":          req.Name,
		"description":   req.Description,
		"category_name": req.CategoryName,
		"thumbnail_id":  req.ThumbnailID,
		"stock":         req.Stock,
		"price":         req.Price,
		"sale_price":    req.SalePrice,
	}

	err := u.repo.UpdateProduct(ctx, u.db, data, "product_id = ?", req.ProductID)
	if err != nil {
		return err
	}

	return nil
}

// DeleteProduct implements IFaceUsecase.
func (u *Usecase) DeleteProduct(ctx context.Context, productID uuid.UUID) error {
	return u.repo.DeleteProduct(ctx, u.db, "product_id = ?", productID)
}
