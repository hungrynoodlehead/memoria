package photo_handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

//	 @Router		/photos/get [get]
//		@Id			getUserPhotos
//		@Summary	Returns ID of all user photos
//		@Tags		photo_repository
//		@Param		userId	query	uint	true	"User ID"
func (h *PhotoHandler) getUserPhotos(w http.ResponseWriter, r *http.Request) {
	userIdStr := r.URL.Query().Get("userId")
	userId, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	photos, err := h.PhotoRepository.GetUserPhotos(userId)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if len(photos) == 0 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	var photosIds []string
	for _, photo := range photos {
		photosIds = append(photosIds, photo.ID)
	}

	w.Header().Set("Content-Type", "application/json")
	result, err := json.Marshal(photosIds)
	fmt.Fprintf(w, string(result))
}
