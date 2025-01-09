package photo_repository

import "errors"

var (
	ErrPhotoNotFound = errors.New("photo not found")
)
