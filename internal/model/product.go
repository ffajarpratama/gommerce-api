package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ProductID    uuid.UUID      `json:"product_id" gorm:"primaryKey"`
	Name         string         `json:"name"`
	Description  string         `json:"description"`
	CategoryName string         `json:"category_name"`
	ThumbnailID  uuid.UUID      `json:"thumbnail_id"`
	Stock        int            `json:"stock"`
	Price        float32        `json:"price"`
	SalePrice    float32        `json:"sale_price"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"column:deleted_at"`

	Thumbnail *Media `json:"thumbnail" gorm:"foreignKey:ThumbnailID; references:MediaID"`
}

func (Product) TableName() string {
	return "tr_product"
}

func (product *Product) BeforeCreate(tx *gorm.DB) (err error) {
	product.ProductID = uuid.New()
	return
}
