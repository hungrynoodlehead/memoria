package photo_handler

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"net/http"
	"strconv"
)

//	 @Router		/photo_repository/{userId}/{photoId} [delete]
//		@Id			deletePhoto
//		@Summary	Deletes a photo_repository
//		@Tags		photo_repository
//		@Param		photoId	path	string	true	"Photo UUID"
//		@Param		userId	path	uint	true	"User ID"
func (h *PhotoHandler) delete(w http.ResponseWriter, r *http.Request) {
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

	if userId != data.UserID {
		http.Error(w, "you have not access to that photo_repository", http.StatusForbidden)
		return
	}

	err = h.PhotoRepository.DeletePhoto(photoUUID)
	if err != nil {
		panic(err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	return
}
