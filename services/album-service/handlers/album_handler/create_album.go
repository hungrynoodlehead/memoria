package album_handler

import (
	"errors"
	"github.com/google/uuid"
	"github.com/hungrynoodlehead/memoria/services/album-service/repositories/photo_repository"
	"github.com/labstack/echo/v4"
)

// @Description Create new album
// @Router /album/create [post]
// @Param form body album_handler.createAlbum.createAlbumForm true "form to be sent"
// @Accept json
// @Produce json
func (h *AlbumHandler) createAlbum(c echo.Context) error {
	type createAlbumForm struct {
		Name        string   `json:"name"`
		Description string   `json:"description"`
		PhotosID    []string `json:"photos_id"`
		OwnerID     uint64   `json:"owner_id"` //TODO: REPLACE WITH JWT!<!<!<!<!
	}

	var errorList map[string]string

	var form createAlbumForm

	err := c.Bind(&form)
	if err != nil {
		return err
	}

	album, err := h.AlbumRepository.Create(form.Name, form.Description, form.OwnerID)
	if err != nil {
		return err
	}

	if len(form.PhotosID) > 0 {
		for _, photoUUID := range form.PhotosID {
			id, err := uuid.Parse(photoUUID)
			if err != nil {
				errorList[photoUUID] = "invalid photo UUID"
				break
			}

			photo, err := h.PhotoRepository.GetPhoto(id)
			if err != nil {
				if errors.Is(err, photo_repository.ErrPhotoNotFound) {
					errorList[photoUUID] = "not found"
					break
				} else {
					return err
				}
			}

			_, err = h.AlbumRepository.AddToAlbum(album, photo)
			if err != nil {
				return err
			}
		}
	}

	c.JSON(200, errorList)
	return nil
}
