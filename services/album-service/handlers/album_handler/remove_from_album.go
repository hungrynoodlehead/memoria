package album_handler

import (
	"errors"
	"github.com/google/uuid"
	"github.com/hungrynoodlehead/memoria/services/album-service/models"
	"github.com/hungrynoodlehead/memoria/services/album-service/repositories/album_repository"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

func (h *AlbumHandler) RemoveFromAlbum(c echo.Context) error {
	type removeFromAlbumForm struct {
		AlbumID   uint64   `param:"id"`
		PhotosIDs []string `json:"photos"`
		Purge     bool     `json:"purge"`
	}

	var form removeFromAlbumForm
	if err := c.Bind(&form); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()}) // TODO: replace with bad request text?
	}
	if form.AlbumID == 0 {
		return c.String(http.StatusBadRequest, "album id is required")
	}
	if len(form.PhotosIDs) == 0 {
		return c.String(http.StatusBadRequest, "photos id is required")
	}

	album, err := h.AlbumRepository.GetByID(form.AlbumID, "Photos")
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return c.String(http.StatusNotFound, "album not found")
		}
		return err
	}

	var errorList map[string]string
	var newAlbum models.Album
	for _, id := range form.PhotosIDs {
		photoUuid, err := uuid.Parse(id)
		if err != nil {
			errorList[id] = "invalid photo id"
			continue
		}

		photo, err := h.PhotoRepository.GetPhoto(photoUuid)
		if err != nil {
			if errors.Is(err, album_repository.ErrAlbumNotFound) {
				errorList[id] = "album not found"
				continue
			}
			return err
		}

		newAlbum, err = h.AlbumRepository.DeleteFromAlbum(album, &photo, form.Purge)
		if err != nil {
			return err
		}
	}

	type removeFromAlbumResponse struct {
		NewAlbum models.Album      `json:"new_album"`
		Errors   map[string]string `json:"errors,omitempty"`
	}
	return c.JSON(http.StatusOK, removeFromAlbumResponse{
		NewAlbum: newAlbum,
		Errors:   errorList,
	})
}
