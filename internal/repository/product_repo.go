package repository

import (
	"context"

	"github.com/ffajarpratama/gommerce-api/internal/http/request"
	"github.com/ffajarpratama/gommerce-api/internal/model"
	"github.com/ffajarpratama/gommerce-api/util"
	"gorm.io/gorm"
)

// CreateProduct implements IFaceRepository.
func (r *Repository) CreateProduct(ctx context.Context, data *model.Product, db *gorm.DB) error {
	return r.BaseRepository.Create(db.WithContext(ctx), data)
}

// FindAndCountProduct implements IFaceRepository.
func (r *Repository) FindAndCountProduct(ctx context.Context, params *request.ListProductQuery) ([]*model.Product, int64, error) {
	var res = make([]*model.Product, 0)
	var cnt int64

	query := r.db.WithContext(ctx).Model(&model.Product{})

	if params.Keyword != "" {
		query = query.Where("name LIKE ?", "%"+params.Keyword+"%")
	}

	if params.CategoryName != "" {
		query = query.Where("category_name = ?", params.CategoryName)
	}

	err := query.Count(&cnt).Error
	if err != nil {
		return nil, 0, err
	}

	if err := query.
		Preload("Thumbnail").
		Limit(params.Limit).
		Offset(util.CalculateOffset(params.Page, params.Limit)).
		Find(&res).
		Error; err != nil {
		return nil, 0, err
	}

	return res, cnt, nil
}

// FindOneProduct implements IFaceRepository.
func (r *Repository) FindOneProduct(ctx context.Context, query ...interface{}) (*model.Product, error) {
	var res *model.Product

	if err := r.BaseRepository.FindOne(r.db.WithContext(ctx).Where(query[0], query[1:]...).Preload("Thumbnail"), &res); err != nil {
		return nil, err
	}

	return res, nil
}

// UpdateProduct implements IFaceRepository.
func (r *Repository) UpdateProduct(ctx context.Context, db *gorm.DB, data map[string]interface{}, query ...interface{}) error {
	return db.WithContext(ctx).Model(&model.Product{}).Where(query[0], query[1:]...).Updates(data).Error
}

// DeleteProduct implements IFaceRepository.
func (r *Repository) DeleteProduct(ctx context.Context, db *gorm.DB, query ...interface{}) error {
	return db.WithContext(ctx).Model(&model.Product{}).Where(query[0], query[1:]...).Delete(&model.Product{}).Error
}
