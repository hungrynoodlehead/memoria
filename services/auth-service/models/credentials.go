package models

import "gorm.io/gorm"

type Credentials struct {
	gorm.Model
	UserID       uint   `gorm:"not null"`
	PasswordHash []byte `gorm:"not null"`
	PasswordSalt []byte `gorm:"not null"`
}
