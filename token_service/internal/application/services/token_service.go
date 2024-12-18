package services

import (
	"fmt"
	"rcarvalho-pb/mottu-token_service/internal/application/dtos"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("minha-chave-secreta")

type TokenService struct {
}

func NewTokenService() *TokenService {
	return &TokenService{}
}

type Claims struct {
	UserID   int64
	Username string
	jwt.RegisteredClaims
}

func (t *TokenService) GenerateJWT(user *dtos.UserDTO) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID:   user.Id,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "mottu-app",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", nil
	}

	return tokenString, nil
}

func (t *TokenService) ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signature method")
		}

		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
