package photo_handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/hungrynoodlehead/memoria/services/storage-service/models"
	"github.com/minio/minio-go/v7"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"
)

// @router			/photo/upload [post]
// @id				uploadPhoto
// @description	Upload a photo
// @tags			photo
// @accept			mpfd
// @param			data	formData	photo_handler.upload.uploadForm		true	"photo data"
// @param			photo	formData	file								true	"photo to be uploaded"
func (h *PhotoHandler) upload(w http.ResponseWriter, r *http.Request) {
	type uploadForm struct {
		Kind    models.PhotoKind `json:"kind" validate:"required" enums:"media,screenshot,meme"`
		AlbumID int              `json:"album_id,omitempty" validate:"optional"`
	}

	err := r.ParseMultipartForm(10 << 21)
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	formString := r.FormValue("metadata")
	if formString == "" {
		http.Error(w, "No photo data provided", http.StatusBadRequest)
		return
	}

	var form uploadForm
	err = json.Unmarshal([]byte(formString), &form)
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
	}

	file, fileHeader, err := r.FormFile("photo")
	if err != nil {
		http.Error(w, "No photo provided", http.StatusBadRequest)
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

	collection := h.DB.Collection("photos")

	photoUuid, err := uuid.NewV7()
	if err != nil {
		//TODO: unified 500 error handler
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	id := photoUuid.String()
	uuidBinary := primitive.Binary{
		Subtype: 0x04, // UUID subtype in MongoDB
		Data:    photoUuid[:],
	}
	objectName := fmt.Sprintf("%d/%s", userId, id)
	contentType := fileHeader.Header.Get("Content-Type")

	_, err = h.Storage.PutObject(
		context.Background(),
		"photos",
		objectName,
		file,
		fileHeader.Size,
		minio.PutObjectOptions{ContentType: contentType},
	)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	photoData := models.Photo{
		ID:          uuidBinary,
		UserID:      userId,
		FileName:    fileHeader.Filename,
		FileSize:    fileHeader.Size,
		ContentType: contentType,
		UploadedAt:  time.Now(),
		Metadata:    nil,
	}

	_, err = collection.InsertOne(context.Background(), photoData)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "ID: %s", id)
	return
}
