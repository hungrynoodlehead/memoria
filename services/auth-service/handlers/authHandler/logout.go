package authhandler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/hungrynoodlehead/photos/services/auth-service/models"
	"gorm.io/gorm"
)

// @Description	Verify a JWT autrhorization
// @Tags			authentication
// @Router			/auth/logout [get]
// @Security		JWT Bearer
func (h *AuthHandler) logout(w http.ResponseWriter, r *http.Request) {
	claims, err := h.JWTUtils.RetrieveClaimsFromContext(r.Context())

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

	var session models.Sessions

	err = h.DB.Model(&models.Sessions{}).Find(&models.Sessions{}, claims.SessionID).First(&session).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Session associated to this token not found", http.StatusNotFound)
			return
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	}

	if session.Status == models.Terminated {
		http.Error(w, "This session is terminated", http.StatusForbidden)
		return
	} else if session.Status == models.Disabled {
		http.Error(w, "This session is inactive", http.StatusForbidden)
		return
	}

	var tokenPair models.TokenPairs
	err = h.DB.Find(&models.TokenPairs{}, claims.TokenID).First(&tokenPair).Error

	if err != nil {
		// TODO: Not found token error
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	tokenPair.Valid = false
	session.Status = models.Terminated
	err = h.DB.Save(&tokenPair).Error
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	err = h.DB.Save(&session).Error
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "Successful logout")
	return
}
