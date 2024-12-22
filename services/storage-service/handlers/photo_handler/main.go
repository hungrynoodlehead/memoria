package photo_handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/hungrynoodlehead/memoria/services/storage-service/repositories/photo_repository"
	"github.com/hungrynoodlehead/memoria/services/storage-service/utils"
)

type PhotoHandler struct {
	*chi.Mux
	DB              *utils.DB
	Storage         *utils.Storage
	Logger          *utils.Logger
	Config          *utils.Config
	PhotoRepository *photo_repository.PhotoRepository
}

func NewPhotoHandler(logger *utils.Logger, DB *utils.DB, storage *utils.Storage, config *utils.Config, photoRepository *photo_repository.PhotoRepository) *PhotoHandler {
	h := PhotoHandler{Mux: chi.NewMux(), Logger: logger, DB: DB, Storage: storage, Config: config, PhotoRepository: photoRepository}

	h.Get("/file/{userId}/{photoId}", h.getPhotoFile)
	h.Post("/upload", h.upload)
	h.Delete("/{userId}/{photoId}", h.delete)

	return &h
}
