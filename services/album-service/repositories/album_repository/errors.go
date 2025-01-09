package album_repository

import "errors"

var (
	ErrPhotoAlreadyInAlbum = errors.New("photo already in album")
	ErrPhotoNotInAlbum     = errors.New("photo not in album")
	ErrAlbumNotFound       = errors.New("album not found")
)
