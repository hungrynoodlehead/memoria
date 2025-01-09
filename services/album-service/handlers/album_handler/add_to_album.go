package album_handler

import (
	"errors"
	"github.com/google/uuid"
	"github.com/hungrynoodlehead/memoria/services/album-service/models"
	"github.com/hungrynoodlehead/memoria/services/album-service/repositories/album_repository"
	"github.com/hungrynoodlehead/memoria/services/album-service/repositories/photo_repository"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// @Description "Add photos to album"
// @Router /album/{id}/add [post]
// @Accept json
// @Produce json
// @Param id path uint64 true "Album ID"
// @Param form body album_handler.addToAlbum.addToAlbumForm true "JSON with array of new photos"
func (h *AlbumHandler) addToAlbum(c echo.Context) error {
	type addToAlbumForm struct {
		Photos []string `json:"photos"` // UUIDs of photos
	}

	albumIdStr := c.Param("id")
	if albumIdStr == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "missing album_id")
	}
	albumId, err := strconv.ParseUint(albumIdStr, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid album_id")
	}

	var form addToAlbumForm
	err = c.Bind(&form)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid form data")
	}
	if len(form.Photos) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid form data")
	}

	album, err := h.AlbumRepository.GetByID(albumId, "Photos")
	if err != nil {
		if errors.Is(err, album_repository.ErrAlbumNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "album not found")
		}
		return err
	}

	type errorsListType map[string]string
	errorsList := make(errorsListType)
	var newAlbum models.Album
	for _, photoId := range form.Photos {
		photoUuid, err := uuid.Parse(photoId)
		if err != nil {
			errorsList[photoId] = "invalid photo id"
			continue
		}
		photo, err := h.PhotoRepository.GetPhoto(photoUuid)
		if err != nil {
			if errors.Is(err, photo_repository.ErrPhotoNotFound) {
				errorsList[photoId] = "photo not found"
				continue
			}
			return err
		}

		newAlbum, err = h.AlbumRepository.AddToAlbum(*album, photo)
		if err != nil {
			return err
		}
	}

	var newPhotosIds []string
	for _, photo := range newAlbum.Photos {
		newPhotosIds = append(newPhotosIds, photo.UUID)
	}

	type addToAlbumResult struct {
		Photos []string       `json:"photos"`
		Errors errorsListType `json:"errors,omitempty"`
	}

	return c.JSON(http.StatusOK, addToAlbumResult{
		Photos: newPhotosIds,
		Errors: errorsList,
	})
}
