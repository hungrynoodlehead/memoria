package userhandler

import (
	"encoding/json"
	"errors"
	"net/http"
	"reflect"

	"github.com/hungrynoodlehead/photos/services/auth-service/helpers"
	"github.com/hungrynoodlehead/photos/services/auth-service/models"
	"gorm.io/gorm"
)

func (h *UserHandler) changePassword(w http.ResponseWriter, r *http.Request) error {
	type changePasswordForm struct {
		OldPassword string `json:"old_password" validate:"required"`
		NewPassword string `json:"new_password" validate:"required"`
	}

	var form changePasswordForm
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	dec.Decode(&form)

	claims, err := h.JWTUtils.RetrieveClaimsFromContext(r.Context())
	if err != nil {
		return err
	}

	var user models.User

	err = h.DB.Preload("Credentials").Find(&models.User{}, claims.UserID).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "User not found", http.StatusNotFound)
			return nil
		}
		return err
	}

	creds := &user.Credentials
	oldHash, _ := helpers.DeriveKey(form.OldPassword, creds.PasswordSalt)

	if !reflect.DeepEqual(creds.PasswordHash, oldHash) {
		http.Error(w, "Wrong password", http.StatusForbidden)
		return nil
	}

	newHash, newSalt := helpers.DeriveKey(form.NewPassword, nil)

	creds.PasswordHash = newHash
	creds.PasswordSalt = newSalt

	h.DB.Save(&user)
	return nil
}
