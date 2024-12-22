package models

import "gorm.io/gorm"

type Album struct {
	gorm.Model
	Name        string   `gorm:"type:varchar(32);not null"`
	Description string   `gorm:"type:text"`
	OwnerID     uint64   `gorm:"not null"`
	Photos      []*Photo `gorm:"many2many:photo_albums;"`
}
