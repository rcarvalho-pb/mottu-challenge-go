package service

import (
	"fmt"

	"github.com/rcarvalho-pb/mottu-motorcycle_service/internal/application/dtos"
	"github.com/rcarvalho-pb/mottu-motorcycle_service/internal/model"
)

type motorcycleService struct {
	model.MotorcycleRepository
}

// GetMotorcycleById(int64) (*Motorcycle, error)
// GetMotorcycleByUserId(int64) (*Motorcycle, error)
// GetAllMotorcycles() ([]*Motorcycle, error)
// GetAllActiveMotorcycles() ([]*Motorcycle, error)
// GetMotorcyclesByYear(int64) ([]*Motorcycle, error)
// GetMotorcyclesByModel(string) ([]*Motorcycle, error)
// CreateMotorcycle(Motorcycle) error
// UpdateMotorcycle(Motorcycle) error
// DeleteMotorcycleById(int64) error
// LocateMotorcycle(int64) error
// UnlocateMotorcycle(int64) error

// UserId    int64     `json:"user_id" db:"user_id"`
// 	Year      int64     `json:"year" db:"year"`
// 	Model     string    `json:"model" db:"model"`
// 	Plate     string    `json:"plate" db:"plate"`
// 	CreatedAt time.Time `json:"created_at" db:"created_at"`
// 	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
// 	IsLocated bool      `json:"is_located" db:"is_located"`
// 	Active

func (ms *motorcycleService) CreateMotorcycle(dto *dtos.MotorcycleDTO) error {
	motorcycle := model.MotorcycleFromDTO(dto)

	if err := ms.MotorcycleRepository.CreateMotorcycle(motorcycle); err != nil {
		return err
	}

	return nil
}

func (ms *motorcycleService) UpdateMotorcycle(dto *dtos.MotorcycleDTO) error {
	motorcycle, err := ms.MotorcycleRepository.GetMotorcycleById(dto.Id)
	if err != nil {
		return err
	}

	motorcycle.Year = dto.Year
	motorcycle.Model = dto.Model
	motorcycle.Plate = dto.Plate
	motorcycle.UpdateTime()

	if err := ms.MotorcycleRepository.UpdateMotorcycle(motorcycle); err != nil {
		return err
	}

	return nil
}

func (ms *motorcycleService) DeleteMotorcycleById(motorcycleId int64) error {
	motorcycle, err := ms.MotorcycleRepository.GetMotorcycleById(motorcycleId)
	if err != nil {
		return err
	}

	if motorcycle.UserId != 0 {
		return fmt.Errorf("motorcycle is located, remove location before deleting")
	}

	motorcycle.UpdateTime()
	motorcycle.Active = false

	if err = ms.MotorcycleRepository.UpdateMotorcycle(motorcycle); err != nil {
		return err
	}

	return nil
}
