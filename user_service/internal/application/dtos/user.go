package dtos

import (
	"time"
)

type ComparePasswordsDTO struct {
	HashedPassword string
	Password       string
}

type UserDTO struct {
	Id             int64     `json:"id"`
	Username       string    `json:"username"`
	Password       string    `json:"password"`
	Role           int       `json:"role"`
	Name           string    `json:"name"`
	BirthDate      time.Time `json:"birth_date"`
	CNPJ           string    `json:"cnpj"`
	CNH            string    `json:"cnh"`
	CNHType        string    `json:"cnh_type"`
	CNHFile        []byte    `json:"cnh_file_path"`
	ActiveLocation bool      `json:"active_location"`
}

type NewUserPasswordDTO struct {
	Id          int64  `json:"id"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}
