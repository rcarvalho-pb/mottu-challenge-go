package model

import (
	"time"

	"github.com/rcarvalho-pb/mottu-motorcycle_service/internal/application/dtos"
)

type MotorcycleRepository interface {
	GetMotorcycleById(int64) (*Motorcycle, error)
	GetMotorcycleByUserId(int64) (*Motorcycle, error)
	GetAllMotorcycles() ([]*Motorcycle, error)
	GetAllActiveMotorcycles() ([]*Motorcycle, error)
	GetMotorcyclesByYear(int64) ([]*Motorcycle, error)
	GetMotorcyclesByModel(string) ([]*Motorcycle, error)
	CreateMotorcycle(*Motorcycle) error
	UpdateMotorcycle(*Motorcycle) error
}

type Motorcycle struct {
	Id        int64     `json:"id" db:"id"`
	UserId    int64     `json:"user_id" db:"user_id"`
	Year      int64     `json:"year" db:"year"`
	Model     string    `json:"model" db:"model"`
	Plate     string    `json:"plate" db:"plate"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	IsLocated bool      `json:"is_located" db:"is_located"`
	Active    bool      `json:"active" db:"active"`
}

func (m *Motorcycle) UpdateTime() {
	m.UpdatedAt = time.Now()
}

func (m *Motorcycle) ToDTO() *dtos.MotorcycleDTO {
	return &dtos.MotorcycleDTO{
		Id:        m.Id,
		UserId:    m.UserId,
		Year:      m.Year,
		Model:     m.Model,
		Plate:     m.Plate,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
		IsLocated: m.IsLocated,
		Active:    m.Active,
	}
}

func MotorcycleFromDTO(dto *dtos.MotorcycleDTO) *Motorcycle {
	return &Motorcycle{
		Id:        dto.Id,
		UserId:    dto.UserId,
		Year:      dto.Year,
		Model:     dto.Model,
		Plate:     dto.Plate,
		CreatedAt: dto.CreatedAt,
		UpdatedAt: dto.UpdatedAt,
		IsLocated: dto.IsLocated,
		Active:    dto.Active,
	}
}
