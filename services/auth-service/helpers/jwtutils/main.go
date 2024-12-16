package jwtutils

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/hungrynoodlehead/photos/services/auth-service/utils"
	"gorm.io/gorm"
)

type JWTUtils struct {
	DB     *gorm.DB
	Config *utils.Config
}

func NewJWTUtils(db *gorm.DB, config *utils.Config) *JWTUtils {
	return &JWTUtils{DB: db, Config: config}
}

type AccessClaims struct {
	jwt.RegisteredClaims
	UserID    uint `json:"uid"`
	SessionID uint `json:"sid"`
	TokenID   uint `json:"tid"`
}
type RefreshClaims struct {
	jwt.RegisteredClaims
	UserID    uint `json:"id"`
	SessionID uint `json:"sid"`
	TokenID   uint `json:"tid"`
}

func (u *JWTUtils) getSigningSecret() []byte {
	return []byte(u.Config.GetString("jwt_token_secret"))
}

