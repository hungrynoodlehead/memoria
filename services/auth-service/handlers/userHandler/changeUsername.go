package userhandler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/hungrynoodlehead/photos/services/auth-service/helpers"
	"github.com/hungrynoodlehead/photos/services/auth-service/models"
	"gorm.io/gorm"
)

func (h *UserHandler) changeUsername(w http.ResponseWriter, r *http.Request) error {
	type changeUsernameForm struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	claims, err := h.JWTUtils.RetrieveClaimsFromContext(r.Context())
	if err != nil {
		return err
	}

	var form changeUsernameForm
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	dec.Decode(&form)

	var user models.User
	err = h.DB.Preload("Credentials").Find(&models.User{}, claims.UserID).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "User not found", http.StatusForbidden)
			return nil
		}
		return err
	}

	hash, _ := helpers.DeriveKey(form.Password, user.Credentials.PasswordSalt)
	if !reflect.DeepEqual(hash, user.Credentials.PasswordHash) {
		http.Error(w, "Wrong password", http.StatusForbidden)
		return nil
	}

	user.Username = form.Username
	h.DB.Save(&user)

	fmt.Fprint(w, "Successfully changed username")
	return nil
}
