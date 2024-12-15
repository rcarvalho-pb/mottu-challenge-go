package model

import (
	"time"

	"github.com/rcarvalho-pb/mottu-user_service/internal/application/dtos"
)

type UserRepository interface {
	GetUserById(int64) (*User, error)
	GetUserByUsername(string) ([]*User, error)
	CreateUser(User) error
	UpdateUser(User) error
	DeleteUser(int64) error
}

type User struct {
	ID             int       `json:"id"`
	Username       string    `json:"username"`
	Password       string    `json:"password"`
	Name           string    `json:"name"`
	BirthDate      time.Time `json:"birth_date"`
	CNPJ           string    `json:"cnpj"`
	CNH            string    `json:"cnh"`
	CNHType        string    `json:"cnh_type"`
	CNHFilePath    string    `json:"cnh_file_path"`
	ActiveLocation bool      `json:"active_location"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Active         bool      `json:"active"`
}

func UserFromDTO(dto dtos.UserDTO) *User {
	return &User{}
}

func (u *User) ToDTO() *dtos.UserDTO {
	return &dtos.UserDTO{}
}
