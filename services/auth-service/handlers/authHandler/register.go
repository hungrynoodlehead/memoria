package authhandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hungrynoodlehead/photos/services/auth-service/helpers"
	"github.com/hungrynoodlehead/photos/services/auth-service/models"
)

//	@Description	Register a new user
//	@Tags			authentication
//	@Router			/auth/register [post]
//	@Accept			json
//	@Produce		json
//	@Param			body	body		authhandler.register.registerForm	true	"User registration form"
//	@Success		200		{object}	authhandler.register.registerResponse
//	@Failure		400		{string}	string	"Bad request"
//	@Failure		403		{string}	string	"User already exists"
//	@Failure		500		{string}	string	"Cannot generate tokens"
func (h *AuthHandler) register(w http.ResponseWriter, r *http.Request) {
	type registerForm struct {
		Username string `json:"username" validate:"required"`
		Email    string `json:"email" validate:"optional"`
		Password string `json:"password" validate:"required"`
	}

	type registerResponse struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}

	var form registerForm
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

	var count int64
	h.DB.Model(&models.User{}).Where(&models.User{Username: form.Username}).Count(&count)
	if count != 0 {
		http.Error(w, "User already exists", http.StatusForbidden)
		return
	}

	user := models.User{Username: form.Username}
	if form.Email != "" {
		h.DB.Model(&models.User{}).Where(&models.User{Email: &form.Email}).Count(&count)
		if count != 0 {
			http.Error(w, "User with this email already exists", http.StatusForbidden)
			return
		}
		user.Email = &form.Email
	}

	hash, salt := helpers.DeriveKey(form.Password, nil)
	user.Credentials = models.Credentials{PasswordHash: hash, PasswordSalt: salt}
	h.DB.Create(&user)
	h.DB.Save(&user)

	session := models.Sessions{
		User:             user,
		FirstUserAgent:   r.Header.Get("User-Agent"),
		CurrentUserAgent: r.Header.Get("User-Agent"),
		FirstIP:          r.Header.Get("X-Forwarded-For"),
		CurrentIP:        r.Header.Get("X-Forwarded-For"),
		Status:           models.Active,
	}
	h.DB.Create(&session)
	h.DB.Save(&session)

	accessToken, refreshToken, err := h.JWTUtils.InitTokenPair(&session)
	if err != nil {
		http.Error(w, "Cannot generate tokens", http.StatusInternalServerError)
		fmt.Print(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(registerResponse{AccessToken: accessToken, RefreshToken: refreshToken})
}
