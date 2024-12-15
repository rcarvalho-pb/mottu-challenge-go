package model

import (
	"time"

	"github.com/rcarvalho-pb/mottu-user_service/internal/application/dtos"
)

type UserRepository interface {
	GetUserById(int64) (*User, error)
	GetUserByUsername(string) ([]*User, error)
	CreateUser(*User) error
	UpdateUser(*User) error
}

type User struct {
	ID             int       `json:"id" db:"id"`
	Username       string    `json:"username" db:"username"`
	Password       string    `json:"password" db:"password"`
	Name           string    `json:"name" db:"name"`
	BirthDate      time.Time `json:"birth_date" db:"birth_date"`
	CNPJ           string    `json:"cnpj" db:"cnpj"`
	CNH            string    `json:"cnh" db:"cnh"`
	CNHType        string    `json:"cnh_type" db:"cnh_type"`
	CNHFilePath    string    `json:"cnh_file_path" db:"cnh_file_path"`
	ActiveLocation bool      `json:"active_location" db:"active_location"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
	Active         bool      `json:"active" db:"active"`
}

func UserFromDTO(dto dtos.UserDTO) *User {
	return &User{}
}

func (u *User) ToDTO() *dtos.UserDTO {
	return &dtos.UserDTO{}
}
