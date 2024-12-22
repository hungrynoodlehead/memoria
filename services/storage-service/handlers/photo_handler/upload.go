package photo_handler

import (
	"encoding/json"
	"fmt"
	"github.com/hungrynoodlehead/memoria/services/storage-service/models"
	"mime/multipart"
	"net/http"
)

// @router			/photo_repository/upload [post]
// @id				uploadPhoto
// @description		Upload a photo_repository
// @tags			photo_repository
// @accept			mpfd
// @param			data	formData	string	true	"photo_repository data"
// @param			photo_repository	formData	file	true	"photo_repository to be uploaded"
func (h *PhotoHandler) upload(w http.ResponseWriter, r *http.Request) {
	type uploadForm struct {
		Kind    models.PhotoKind `json:"kind" validate:"required" enums:"media,screenshot,meme"`
		AlbumID uint             `json:"album_id,omitempty" validate:"optional"`
		UserID  uint64           `json:"user_id" validate:"required"`
	}

	err := r.ParseMultipartForm(10 << 21)
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	var form uploadForm
	dataStr := r.FormValue("data")
	err = json.Unmarshal([]byte(dataStr), &form)
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	file, fileHeader, err := r.FormFile("photo_repository")
	if err != nil {
		http.Error(w, "No photo_repository provided", http.StatusBadRequest)
		return
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			// TODO: handler function
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	}(file)

	photo, err := h.PhotoRepository.CreatePhoto(form.UserID, file, fileHeader.Filename, fileHeader.Size, form.Kind, fileHeader.Header.Get("Content-Type"), form.AlbumID)
	if err != nil {
		http.Error(w, "Error creating photo_repository", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "ID: %s", photo.ID)
	return
}
