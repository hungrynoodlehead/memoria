package middlewares

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/hungrynoodlehead/photos/services/auth-service/helpers/jwtutils"
	"github.com/hungrynoodlehead/photos/services/auth-service/models"
	"gorm.io/gorm"
)

type AuthMiddleware struct {
	DB       *gorm.DB
	jwtUtils *jwtutils.JWTUtils
}

func NewAuthMiddleware(db *gorm.DB, jwtUtils *jwtutils.JWTUtils) *AuthMiddleware {
	return &AuthMiddleware{DB: db, jwtUtils: jwtUtils}
}

func (mw *AuthMiddleware) Verify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")

		if token == "" {
			http.Error(w, "Access token is required", http.StatusBadRequest)
			return
		}

		claims, err := mw.jwtUtils.VerifyAccessToken(token)

		if err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) {
				http.Error(w, "Token is expired", http.StatusForbidden)
				return
			} else if errors.Is(err, jwt.ErrTokenUsedBeforeIssued) {
				http.Error(w, "Token used before issued", http.StatusForbidden)
				return
			} else if errors.Is(err, jwt.ErrTokenMalformed) {
				http.Error(w, "Token is malformed", http.StatusForbidden)
				return
			} else if errors.Is(err, jwt.ErrTokenInvalidClaims) {
				http.Error(w, "Token is invalid", http.StatusForbidden)
				return
			} else if errors.Is(err, jwt.ErrSignatureInvalid) {
				http.Error(w, "Token signature is invalid", http.StatusForbidden)
				return
			} else if errors.Is(err, jwtutils.ErrTokenDisabled) {
				http.Error(w, "Token is disabled", http.StatusForbidden)
				return
			} else if errors.Is(err, jwtutils.ErrTokenNotFound) {
				http.Error(w, "Token not found in database", http.StatusForbidden)
				return
			} else {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}
		}

		var session models.Sessions

		err = mw.DB.Find(&models.Sessions{}, claims.SessionID).First(&session).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				http.Error(w, "Session not found in database", http.StatusForbidden)
				return
			} else {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}
		}

		if session.Status != models.Active {
			http.Error(w, "This session is not active", http.StatusForbidden)
			return
		}

		ctx := context.WithValue(r.Context(), "claims", claims)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
