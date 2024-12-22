package album_repository

import "gorm.io/gorm"

type AlbumRepository struct {
	DB *gorm.DB
}

func NewAlbumRepository(db *gorm.DB) (*AlbumRepository, error) {
	return &AlbumRepository{
		DB: db,
	}, nil
}
