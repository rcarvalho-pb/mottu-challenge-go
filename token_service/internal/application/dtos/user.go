package dtos

import "github.com/golang-jwt/jwt/v5"

type UserDTO struct {
	Id       int64
	Username string
	Role     int
}

type Claims struct {
	UserID   int64
	Username string
	UserRole int
	jwt.RegisteredClaims
}
