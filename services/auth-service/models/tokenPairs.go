package models

import "gorm.io/gorm"

type TokenPairs struct {
	*gorm.Model
	Valid     bool
	SessionID uint `gorm:"not null"`
	Session   Sessions
}
