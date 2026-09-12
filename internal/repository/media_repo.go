package repository

import (
	"context"

	"github.com/ffajarpratama/gommerce-api/internal/model"
	"gorm.io/gorm"
)

// CreateMedia implements IFaceRepository.
func (r *Repository) CreateMedia(ctx context.Context, data *model.Media, db *gorm.DB) error {
	return r.BaseRepository.Create(db.WithContext(ctx), data)
}
