package model

import (
	"fmt"
	"time"

	"github.com/ffajarpratama/gommerce-api/config"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Media struct {
	MediaID   uuid.UUID      `json:"media_id" gorm:"primaryKey"`
	Name      string         `json:"name"`
	Size      int            `json:"size"`
	Path      string         `json:"path"`
	Mimetype  string         `json:"mimetype"`
	MediaURL  string         `json:"media_url"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"column:deleted_at"`

	// json response
	CloudinaryURL string `json:"cloudinary_url" gorm:"-"`
}

func (Media) TableName() string {
	return "tr_media"
}

func (media *Media) BeforeCreate(tx *gorm.DB) (err error) {
	media.MediaID = uuid.New()
	return
}

func (media *Media) AfterFind(db *gorm.DB) (err error) {
	media.CloudinaryURL = fmt.Sprintf("https://res.cloudinary.com/%s/image/upload/w_500/%s", config.CLDCloudName, media.Path)
	return
}
