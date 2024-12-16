package jwtutils

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/hungrynoodlehead/photos/services/auth-service/models"
	"gorm.io/gorm"
)

func (u *JWTUtils) VerifyAccessToken(accessToken string) (*AccessClaims, error) {
	token, err := jwt.ParseWithClaims(accessToken, &AccessClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrTokenUnverifiable
		}

		return u.getSigningSecret(), nil
	}, jwt.WithExpirationRequired(), jwt.WithIssuedAt())

	if err != nil {
		return nil, err
	} else if claims, ok := token.Claims.(*AccessClaims); ok {
		if err := validateAccessClaims(claims); err != nil {
			return nil, err
		}

		var token models.TokenPairs

		err := u.DB.Find(&models.TokenPairs{}, claims.TokenID).First(&token).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, ErrTokenNotFound
			} else {
				return nil, err
			}
		}

		if !token.Valid {
			return nil, ErrTokenDisabled
		}

		return claims, nil
	} else {
		return nil, jwt.ErrTokenInvalidClaims
	}
}

func validateAccessClaims(c *AccessClaims) error {
	if c.UserID == 0 {
		return jwt.ErrTokenRequiredClaimMissing
	}
	if c.SessionID == 0 {
		return jwt.ErrTokenRequiredClaimMissing
	}
	if c.TokenID == 0 {
		return jwt.ErrTokenRequiredClaimMissing
	}
	return nil
}

func (u *JWTUtils) VerifyRefreshToken(refreshToken string) (*RefreshClaims, error) {
	token, err := jwt.ParseWithClaims(refreshToken, &RefreshClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrTokenUnverifiable
		}

		return u.getSigningSecret(), nil
	}, jwt.WithExpirationRequired(), jwt.WithIssuedAt())

	if err != nil {
		return nil, err
	} else if claims, ok := token.Claims.(*RefreshClaims); ok {
		if err := validateRefreshClaims(claims); err != nil {
			return nil, err
		}
		return claims, nil
	} else {
		return nil, jwt.ErrTokenInvalidClaims
	}
}

func validateRefreshClaims(c *RefreshClaims) error {
	if c.UserID == 0 {
		return jwt.ErrTokenRequiredClaimMissing
	} else if c.SessionID == 0 {
		return jwt.ErrTokenRequiredClaimMissing
	} else if c.TokenID == 0 {
		return jwt.ErrTokenRequiredClaimMissing
	}
	return nil
}
