package jwtutils

import (
	"time"

	"github.com/hungrynoodlehead/photos/services/auth-service/models"
)

func (u *JWTUtils) InitTokenPair(session *models.Sessions) (string, string, error) {
	tokenPair := models.TokenPairs{
		Valid:   true,
		Session: *session,
	}

	u.DB.Create(&tokenPair)
	u.DB.Save(&tokenPair)

	secret := u.getSigningSecret()

	issuedAt := time.Now()
	expiresAt := issuedAt.Add(u.Config.GetDuration("jwt_access_token_duration"))

	accessToken, err := u.generateAccessToken(
		session.User.ID,
		session.ID,
		tokenPair.ID,
		&issuedAt,
		&expiresAt,
		secret,
	)
	if err != nil {
		// TODO: Logger
		return "", "", ErrGenToken
	}

	notBefore := expiresAt
	expiresAt = issuedAt.Add(u.Config.GetDuration("jwt_refresh_token_duration"))

	refreshToken, err := u.generateRefreshToken(
		session.User.ID,
		session.ID,
		tokenPair.ID,
		&issuedAt,
		&notBefore,
		&expiresAt,
		secret,
	)
	if err != nil {
		// TODO: Logger
		return "", "", ErrGenToken
	}

	return accessToken, refreshToken, nil
}
