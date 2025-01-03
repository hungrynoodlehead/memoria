package photo_handler

import (
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"io"
	"net/http"
	"strconv"
)

//	 @Router		/photos/file/{userId}/{photoId} [get]
//		@Id			getPhotoFile
//		@Summary	Returns file from storage
//		@Tags		photo_repository
//		@Param		photoId	path	string	true	"Photo UUID"
//		@Param		userId	path	uint	true	"User ID"
func (h *PhotoHandler) getPhotoFile(w http.ResponseWriter, r *http.Request) {
	photoId := chi.URLParam(r, "photoId")
	if photoId == "" {
		http.Error(w, "photo_repository id is required", http.StatusBadRequest)
		return
	}

	//TODO: auth through jwt middleware
	userIdString := chi.URLParam(r, "userId")
	if userIdString == "" {
		http.Error(w, "userId is required", http.StatusBadRequest)
		return
	}
	userId, err := strconv.ParseUint(userIdString, 10, 64)
	if err != nil {
		http.Error(w, "userId is invalid", http.StatusBadRequest)
		return
	}

	photoUUID, err := uuid.Parse(photoId)
	if err != nil {
		http.Error(w, "photo_repository uuid is invalid", http.StatusBadRequest)
	}

	data, err := h.PhotoRepository.GetPhotoData(photoUUID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			http.Error(w, "photo_repository not found", http.StatusNotFound)
			return
		}
		panic(err)
		return
	}

	if data.UserID != userId {
		http.Error(w, "you have not access to that photo_repository", http.StatusForbidden)
		return
	}

	file, err := h.PhotoRepository.GetPhotoFile(photoUUID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			http.Error(w, "photo_repository not found", http.StatusNotFound)
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	r.Header.Set("Content-Type", data.ContentType)
	r.Header.Set("Content-Disposition", fmt.Sprintf("inline; filename='%s'", data.FileName))

	defer file.Close()
	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
