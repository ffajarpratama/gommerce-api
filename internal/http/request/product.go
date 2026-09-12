package request

import "github.com/google/uuid"

type ListProductQuery struct {
	BaseQuery
	CategoryName string
}

type CreateProduct struct {
	Name         string    `json:"name" validate:"required"`
	Description  string    `json:"description" validate:"required"`
	CategoryName string    `json:"category_name" validate:"required"`
	ThumbnailID  uuid.UUID `json:"thumbnail_id" validate:"required"`
	Stock        int       `json:"stock"`
	Price        float32   `json:"price"`
	SalePrice    float32   `json:"sale_price"`
}

type UpdateProduct struct {
	Name         string    `json:"name" validate:"required"`
	Description  string    `json:"description" validate:"required"`
	CategoryName string    `json:"category_name" validate:"required"`
	ThumbnailID  uuid.UUID `json:"thumbnail_id" validate:"required"`
	Stock        int       `json:"stock"`
	Price        float32   `json:"price"`
	SalePrice    float32   `json:"sale_price"`
	ProductID    uuid.UUID `json:"-"`
}
