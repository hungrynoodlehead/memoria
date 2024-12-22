package models

import "gorm.io/gorm"

type Photo struct {
	gorm.Model
	UUID   string
	Albums []*Album `gorm:"many2many:photo_albums;"`
	Kind   PhotoKind
}

type PhotoKind string

const (
	PhotoKindMedia       PhotoKind = "media"
	PhotoKindScreenshots PhotoKind = "screenshots"
	PhotoKindMemes       PhotoKind = "memes"
)
