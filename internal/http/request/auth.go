package request

import "github.com/ffajarpratama/gommerce-api/constant"

type Register struct {
	Name        string            `json:"name" validate:"required"`
	Email       string            `json:"email" validate:"required,email"`
	PhoneNumber string            `json:"phone_number" validate:"required"`
	Password    string            `json:"password" validate:"required"`
	Role        constant.UserRole `json:"-"`
}

type Login struct {
	Email    string            `json:"email" validate:"required,email"`
	Password string            `json:"password" validate:"required"`
	Role     constant.UserRole `json:"-"`
}
