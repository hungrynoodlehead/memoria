package album_handler

import (
	"errors"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

func (h *AlbumHandler) DeleteAlbum(c echo.Context) error {
	// TODO: AUTH

	type deleteAlbumForm struct {
		AlbumID uint64 `param:"id"`
		Purge   bool   `query:"purge"`
	}

	var form deleteAlbumForm
	if err := c.Bind(&form); err != nil {
		return c.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}
	if form.AlbumID == 0 {
		return c.JSON(http.StatusBadRequest, "album ID must be specified")
	}

	album, err := h.AlbumRepository.GetByID(form.AlbumID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, "album not found")
		}
		return err
	}

	err = h.AlbumRepository.DeleteAlbum(album, form.Purge)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
