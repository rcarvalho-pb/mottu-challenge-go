package dtos

import "github.com/golang-jwt/jwt/v5"

type UserDTO struct {
	Id       int64
	Username string
	Role     string
}

type Claims struct {
	UserID   int64
	Username string
	UserRole string
	jwt.RegisteredClaims
}
