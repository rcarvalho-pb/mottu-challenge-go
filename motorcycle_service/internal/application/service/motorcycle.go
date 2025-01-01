package service

import (
	"github.com/rcarvalho-pb/mottu-motorcycle_service/internal/application/dtos"
	"github.com/rcarvalho-pb/mottu-motorcycle_service/internal/model"
)

type MotorcycleService struct {
	repository model.MotorcycleRepository
}

func New(repository model.MotorcycleRepository) *MotorcycleService {
	return &MotorcycleService{
		repository: repository,
	}
}

func (ms *MotorcycleService) GetMotorcycleById(motorcycleId int64) (*dtos.MotorcycleDTO, error) {
	motorcycle, err := ms.repository.GetMotorcycleById(motorcycleId)
	if err != nil {
		return nil, err
	}

	return motorcycle.ToDTO(), err
}

func (ms *MotorcycleService) GetAllMotorcycles() ([]*dtos.MotorcycleDTO, error) {
	motorcycles, err := ms.repository.GetAllMotorcycles()
	if err != nil {
		return nil, err
	}

	motorcyclesDTO := make([]*dtos.MotorcycleDTO, len(motorcycles))
	for i, m := range motorcycles {
		motorcyclesDTO[i] = m.ToDTO()
	}

	return motorcyclesDTO, nil
}

func (ms *MotorcycleService) GetAllActiveMotorcycles() ([]*dtos.MotorcycleDTO, error) {
	motorcycles, err := ms.repository.GetAllActiveMotorcycles()
	if err != nil {
		return nil, err
	}

	motorcyclesDTO := make([]*dtos.MotorcycleDTO, len(motorcycles))
	for i, m := range motorcycles {
		motorcyclesDTO[i] = m.ToDTO()
	}

	return motorcyclesDTO, nil
}

func (ms *MotorcycleService) GetMotorcyclesByYear(year int64) ([]*dtos.MotorcycleDTO, error) {
	motorcycles, err := ms.repository.GetMotorcyclesByYear(year)
	if err != nil {
		return nil, err
	}

	motorcyclesDTO := make([]*dtos.MotorcycleDTO, len(motorcycles))

	for i, m := range motorcycles {
		motorcyclesDTO[i] = m.ToDTO()
	}

	return motorcyclesDTO, nil
}

func (ms *MotorcycleService) GetMotorcyclesByModel(model string) ([]*dtos.MotorcycleDTO, error) {
	motorcycles, err := ms.repository.GetMotorcyclesByModel(model)
	if err != nil {
		return nil, err
	}

	motorcyclesDTO := make([]*dtos.MotorcycleDTO, len(motorcycles))

	for i, m := range motorcycles {
		motorcyclesDTO[i] = m.ToDTO()
	}

	return motorcyclesDTO, nil
}

func (ms *MotorcycleService) CreateMotorcycle(req *dtos.NewMotorcycleRequest) error {
	if err := ms.repository.CreateMotorcycle(req); err != nil {
		return err
	}

	return nil
}

func (ms *MotorcycleService) UpdateMotorcycle(dto *dtos.MotorcycleDTO) error {
	motorcycle, err := ms.repository.GetMotorcycleById(dto.Id)
	if err != nil {
		return err
	}

	motorcycle.Year = dto.Year
	motorcycle.Model = dto.Model
	motorcycle.Plate = dto.Plate
	motorcycle.UpdateTime()

	if err := ms.repository.UpdateMotorcycle(motorcycle); err != nil {
		return err
	}

	return nil
}

func (ms *MotorcycleService) DeleteMotorcycleById(motorcycleId int64) error {
	motorcycle, err := ms.repository.GetMotorcycleById(motorcycleId)
	if err != nil {
		return err
	}

	motorcycle.UpdateTime()
	motorcycle.Active = false

	if err = ms.repository.UpdateMotorcycle(motorcycle); err != nil {
		return err
	}

	return nil
}

func (ms *MotorcycleService) LocateMotorcycle(motorcycleId int64) error {
	motorcycle, err := ms.repository.GetMotorcycleById(motorcycleId)
	if err != nil {
		return err
	}

	motorcycle.UpdateTime()
	motorcycle.IsLocated = true

	if err = ms.repository.UpdateMotorcycle(motorcycle); err != nil {
		return err
	}

	return nil
}

func (ms *MotorcycleService) UnlocateMotorcycle(motorcycleId int64) error {
	motorcycle, err := ms.repository.GetMotorcycleById(motorcycleId)
	if err != nil {
		return err
	}

	motorcycle.UpdateTime()
	motorcycle.IsLocated = false

	if err = ms.repository.UpdateMotorcycle(motorcycle); err != nil {
		return err
	}

	return nil
}
