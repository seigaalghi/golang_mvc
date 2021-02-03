package models

import "gorm.io/gorm"

// Author adalah models atau schema untuk author
type Author struct {
	gorm.Model
	Name     string `json:"name" gorm:"type:varchar(100)" binding:"required,min=4"`
	Email    string `json:"email" gorm:"type:varchar(100)" binding:"required,email"`
	Password string `json:"password" gorm:"type:varchar(200)" binding:"required,min=6,alphanum"`
}

// LoginForm is ..
type LoginForm struct {
	Email    string `json:"email" gorm:"type:varchar(100)" binding:"required,email"`
	Password string `json:"password" gorm:"type:varchar(200)" binding:"required,min=6,alphanum"`
}
