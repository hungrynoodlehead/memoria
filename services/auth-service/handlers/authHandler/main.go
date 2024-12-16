package authhandler

import (
	"github.com/go-chi/chi/v5"
	"github.com/hungrynoodlehead/photos/services/auth-service/helpers/jwtutils"
	"github.com/hungrynoodlehead/photos/services/auth-service/middlewares"
	"gorm.io/gorm"
)

type AuthHandler struct {
	*chi.Mux
	DB             *gorm.DB
	JWTUtils       *jwtutils.JWTUtils
	AuthMiddleware *middlewares.AuthMiddleware
}

func NewAuthHandler(db *gorm.DB, jwtutils *jwtutils.JWTUtils, authMiddleware *middlewares.AuthMiddleware) *AuthHandler {
	ah := AuthHandler{Mux: chi.NewRouter(), DB: db, JWTUtils: jwtutils, AuthMiddleware: authMiddleware}

	ah.Post("/register", ah.register)
	ah.Post("/login", ah.login)

	ah.Group(func(r chi.Router) {
		r.Use(ah.AuthMiddleware.Verify)
		r.Get("/verify", ah.verify)
		r.Get("/logout", ah.logout)
	})

	return &ah
}
