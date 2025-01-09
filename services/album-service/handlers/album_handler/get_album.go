package album_handler

import (
	"errors"
	"github.com/hungrynoodlehead/memoria/services/album-service/repositories/album_repository"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// @Description Get album data
// @Router /album/{id} [get]
// @Param id path uint64 true "Album ID"
// @Produce json
func (h *AlbumHandler) getAlbum(e echo.Context) error {
	albumIdStr := e.Param("id")
	if albumIdStr == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "missing albumId")
	}
	albumId, err := strconv.ParseUint(albumIdStr, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid albumId")
	}

	album, err := h.AlbumRepository.GetByID(albumId, "Photos")
	if err != nil {
		if errors.Is(err, album_repository.ErrAlbumNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "album not found")
		}
		return err
	}

	type response struct {
		ID          uint     `json:"id"`
		Name        string   `json:"name"`
		Description string   `json:"description"`
		OwnerID     uint64   `json:"owner_id"`
		Photos      []string `json:"photos"`
	}

	var ph []string
	for _, p := range album.Photos {
		ph = append(ph, p.UUID)
	}

	return e.JSON(http.StatusOK, response{
		ID:          album.ID,
		Name:        album.Name,
		Description: album.Description,
		OwnerID:     album.OwnerID,
		Photos:      ph,
	})
}
