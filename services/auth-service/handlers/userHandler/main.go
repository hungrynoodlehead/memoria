package userhandler

import (
	"github.com/hungrynoodlehead/photos/services/auth-service/helpers/jwtutils"
	"gorm.io/gorm"
)

type UserHandler struct {
	DB       *gorm.DB
	JWTUtils *jwtutils.JWTUtils
}

func NewUserHandler(db *gorm.DB, jwtutils *jwtutils.JWTUtils) UserHandler {
	return UserHandler{DB: db, JWTUtils: jwtutils}
	//TODO: bind handlers
}
