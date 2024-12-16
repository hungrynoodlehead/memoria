package authhandler

import (
	"fmt"
	"net/http"

	"github.com/hungrynoodlehead/photos/services/auth-service/helpers/jwtutils"
)

// @Description	Verify a JWT autrhorization
// @Tags			authentication
// @Router			/auth/verify [get]
// @Security		JWT Bearer
func (h *AuthHandler) verify(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value("claims").(*jwtutils.AccessClaims)
	if !ok || claims == nil {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	
	fmt.Fprint(w, true)
}
