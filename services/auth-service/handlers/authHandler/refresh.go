package authhandler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/hungrynoodlehead/photos/services/auth-service/helpers/jwtutils"
)

// @Description	Get new token pair
// @Tags			authentication
// @Router			/auth/refresh [post]
// @Accept			json
// @Produce		json
// @Param			body	body		authhandler.refresh.refreshForm	true	"User registration form"
// @Success		200		{object}	authhandler.refresh.refreshResponse
func (h *AuthHandler) refresh(w http.ResponseWriter, r *http.Request) {
	type refreshForm struct {
		Token string `json:"refresh_token" validate:"required"`
	}

	type refreshResponse struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}

	var form refreshForm

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(&form)

	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
	}

	if form.Token == "" {
		http.Error(w, "Refresh token cannot be empty", http.StatusBadRequest)
	}

	accessToken, refreshToken, err := h.JWTUtils.Refresh(form.Token)

	if err != nil {
		if errors.Is(err, jwtutils.ErrTokenAlreadyBeenUsed) {
			//TODO:
			http.Error(w, "Token already been used", http.StatusForbidden)
			return
		}
	}

	res := refreshResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&res)
}
