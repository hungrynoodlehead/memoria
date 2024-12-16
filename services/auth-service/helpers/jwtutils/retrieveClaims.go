package jwtutils

import (
	"context"
	"errors"
)

func (_ *JWTUtils) RetrieveClaimsFromContext(ctx context.Context) (*AccessClaims, error) {
	claims, ok := ctx.Value("claims").(*AccessClaims)

	if !ok {
		return nil, errors.New("Cannot retrieve claims")
	}

	return claims, nil
}
