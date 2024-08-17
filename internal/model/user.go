package model

import (
	"time"

	"github.com/ffajarpratama/gommerce-api/constant"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	UserID      uuid.UUID         `json:"user_id" gorm:"primaryKey"`
	Name        string            `json:"name"`
	Email       string            `json:"email"`
	PhoneNumber string            `json:"phone_number"`
	Password    string            `json:"-" gorm:"column:password"`
	Role        constant.UserRole `json:"role"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
	DeletedAt   gorm.DeletedAt    `json:"-" gorm:"column:deleted_at"`

	// json fields
	AccessToken  string `json:"access_token,omitempty" gorm:"-"`
	RefreshToken string `json:"refresh_token,omitempty" gorm:"-"`
}

func (User) TableName() string {
	return "tr_user"
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.UserID = uuid.New()
	return
}
