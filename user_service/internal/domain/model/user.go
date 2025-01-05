package model

import (
	"database/sql"
	"slices"
	"time"

	"github.com/rcarvalho-pb/mottu-user_service/internal/application/dtos"
)

var ROLES = []string{"admin", "common"}

type Role int

const (
	Admin Role = iota
	Common
)

func (r Role) String() string {
	return ROLES[r]
}

func RoleToCod(role string) Role {
	return Role(slices.Index(ROLES, role))
}

type UserRepository interface {
	GetAllUsers() ([]*User, error)
	GetAllActiveUsers() ([]*User, error)
	GetUserById(int64) (*User, error)
	GetUserByUsername(string) (*User, error)
	CreateUser(*User) error
	UpdateUser(*User) error
}

type User struct {
	Id             int64          `json:"id" db:"id"`
	Username       string         `json:"username" db:"username"`
	Password       string         `json:"password" db:"password"`
	Role           Role           `json:"role" db:"role"`
	Name           string         `json:"name" db:"name"`
	BirthDate      time.Time      `json:"birth_date" db:"birth_date"`
	CNPJ           string         `json:"cnpj" db:"cnpj"`
	CNH            string         `json:"cnh" db:"cnh"`
	CNHType        string         `json:"cnh_type" db:"cnh_type"`
	CNHFilePath    sql.NullString `json:"cnh_file_path" db:"cnh_file_path"`
	ActiveLocation bool           `json:"active_location" db:"active_location"`
	CreatedAt      time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at" db:"updated_at"`
	Active         bool           `json:"active" db:"active"`
}

func (u *User) UpdateTime() {
	u.UpdatedAt = time.Now()
}

func UserFromDTO(dto *dtos.UserDTO) *User {
	return &User{
		Id:        dto.Id,
		Username:  dto.Username,
		Password:  dto.Password,
		Role:      RoleToCod(dto.Role),
		Name:      dto.Username,
		BirthDate: dto.BirthDate,
		CNPJ:      dto.CNPJ,
		CNH:       dto.CNH,
		CNHType:   dto.CNHType,
		CNHFilePath: sql.NullString{
			Valid:  true,
			String: "teste",
		},
		ActiveLocation: dto.ActiveLocation,
	}
}

func (u *User) ToDTO() *dtos.UserDTO {
	return &dtos.UserDTO{
		Id:             u.Id,
		Username:       u.Username,
		Password:       u.Password,
		Role:           u.Role.String(),
		Name:           u.Name,
		BirthDate:      u.BirthDate,
		CNPJ:           u.CNPJ,
		CNH:            u.CNH,
		CNHType:        u.CNHType,
		ActiveLocation: u.ActiveLocation,
	}
}
