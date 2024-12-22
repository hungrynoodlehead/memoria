package authhandler

import (
	"encoding/json"
	"errors"
	"net/http"
	"reflect"

	"github.com/hungrynoodlehead/photos/services/auth-service/helpers"
	"github.com/hungrynoodlehead/photos/services/auth-service/models"
	"gorm.io/gorm"
)

// @Description	User login
// @Tags			authentication
// @Router			/auth/login [post]
// @Accept			json
// @Produce		json
// @Param		body	body		authhandler.login.loginForm	true	"User login form"
// @Success		200		{object}	authhandler.login.loginResponse
// @Failure		400		{string}	string	"Bad request"
// @Failure		401		{string}	string	"User with this Username and Password not found"
// @Failure		500		{string}	string	"Cannot generate tokens"
func (h *AuthHandler) login(w http.ResponseWriter, r *http.Request) {
	type loginForm struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	type loginResponse struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}

	var form loginForm
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(&form)

	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if form.Username == "" || form.Password == "" {
		http.Error(w, "Username and Password fields cannot be empty", http.StatusBadRequest)
		return
	}

	var user models.User

	err = h.DB.Preload("Credentials").Model(&models.User{}).Where(&models.User{Username: form.Username}).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "User with this Username and Password not found", http.StatusForbidden)
			return
		} else {
			//TODO: normal handler

			panic(err)
		}
	}

	hash, _ := helpers.DeriveKey(form.Password, user.Credentials.PasswordSalt)

	if !reflect.DeepEqual(hash, user.Credentials.PasswordHash) {
		http.Error(w, "User with this Username and Password not found", http.StatusForbidden)
		return
	}

	session := models.Sessions{
		User:             user,
		FirstUserAgent:   r.UserAgent(),
		CurrentUserAgent: r.UserAgent(),
		FirstIP:          r.RemoteAddr,
		CurrentIP:        r.RemoteAddr,
		Status:           models.Active,
	}

	err = h.DB.Create(&session).Error
	if err != nil {
		panic(err)
	}
	err = h.DB.Save(&session).Error
	if err != nil {
		panic(err)
	}

	accessToken, refreshToken, err := h.JWTUtils.InitTokenPair(&session)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(loginResponse{AccessToken: accessToken, RefreshToken: refreshToken})
	return
}
