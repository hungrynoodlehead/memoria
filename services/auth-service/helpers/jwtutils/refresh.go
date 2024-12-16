package jwtutils

import (
	"github.com/hungrynoodlehead/photos/services/auth-service/models"
)

func (u *JWTUtils) Refresh(refreshToken string) (string, string, error) {
	claims, err := u.VerifyRefreshToken(refreshToken)

	if err != nil {
		return "", "", err
	}

	var oldTokens models.TokenPairs

	err = u.DB.Model(&models.TokenPairs{}).First(&oldTokens, claims.TokenID).Error

	if !oldTokens.Valid {
		//TODO: Delete session
		return "", "", ErrTokenAlreadyBeenUsed
	}

	session := oldTokens.Session
	newAccessToken, newRefreshToken, err := u.InitTokenPair(&session)
	oldTokens.Valid = false
	err = u.DB.Save(&oldTokens).Error

	if err != nil {
		return "", "", err
	}

	return newAccessToken, newRefreshToken, nil
}
