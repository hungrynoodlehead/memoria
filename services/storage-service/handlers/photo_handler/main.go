package photo_handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/hungrynoodlehead/memoria/services/storage-service/utils"
)

type PhotoHandler struct {
	*chi.Mux
	DB      *utils.DB
	Storage *utils.Storage
	Logger  *utils.Logger
	Config  *utils.Config
}

func NewPhotoHandler(logger *utils.Logger, DB *utils.DB, storage *utils.Storage, config *utils.Config) *PhotoHandler {
	h := PhotoHandler{Mux: chi.NewMux(), Logger: logger, DB: DB, Storage: storage, Config: config}

	h.Get("/{user_id}/{photo_id}", h.get)
	h.Post("/upload", h.upload)

	return &h
}
