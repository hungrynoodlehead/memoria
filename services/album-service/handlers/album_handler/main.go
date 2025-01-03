package album_handler

import (
	"github.com/hungrynoodlehead/memoria/services/album-service/repositories/album_repository"
	"github.com/hungrynoodlehead/memoria/services/album-service/repositories/photo_repository"
	"github.com/labstack/echo/v4"
)

type AlbumHandler struct {
	AlbumRepository *album_repository.AlbumRepository
	PhotoRepository *photo_repository.PhotoRepository
}

func BindAlbumHandler(g *echo.Group, albumRepository *album_repository.AlbumRepository, photoRepository *photo_repository.PhotoRepository) error {
	h := AlbumHandler{
		AlbumRepository: albumRepository,
		PhotoRepository: photoRepository,
	}

	g.POST("/create", h.createAlbum)
	g.GET("/:id", h.getAlbum)
	g.POST("/:id/add", h.addToAlbum)
	return nil
}
