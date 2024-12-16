package jwtutils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (_ *JWTUtils) generateAccessToken(userId uint, sessionId uint, tokenId uint, issuedAt *time.Time, expiresAt *time.Time, secret any) (string, error) {
	claims := AccessClaims{
		UserID:    userId,
		SessionID: sessionId,
		TokenID:   tokenId,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(*issuedAt),
			ExpiresAt: jwt.NewNumericDate(*expiresAt),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secret)

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (_ *JWTUtils) generateRefreshToken(userId uint, sessionId uint, tokenId uint, issuedAt *time.Time, notBefore *time.Time, expiresAt *time.Time, secret any) (string, error) {
	claims := RefreshClaims{
		UserID:    userId,
		SessionID: sessionId,
		TokenID:   tokenId,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(*issuedAt),
			NotBefore: jwt.NewNumericDate(*notBefore),
			ExpiresAt: jwt.NewNumericDate(*expiresAt),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secret)

	if err != nil {
		return "", err
	}

	return signedToken, err
}
