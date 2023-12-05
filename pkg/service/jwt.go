package service

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

type JWTManager struct {
	secretKey string
}

type UserClaims struct {
	jwt.StandardClaims
	GuessMe string `json:"guess_me"`
}

func NewJWTManager(secretKey string) *JWTManager {
	return &JWTManager{secretKey}
}

func (manager *JWTManager) Verify(accessToken string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected token signing method")
			}

			return []byte(manager.secretKey), nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
