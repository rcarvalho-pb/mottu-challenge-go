package dtos

import "github.com/golang-jwt/jwt/v5"

type UserDTO struct {
	Id       int64
	Username string
}

type Claims struct {
	UserID   int64
	Username string
	jwt.RegisteredClaims
}
