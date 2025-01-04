package service

import (
	"fmt"

	"github.com/rcarvalho-pb/mottu-broker_service/internal/model"
	rpc_client "github.com/rcarvalho-pb/mottu-broker_service/internal/rpc/client"
)

type motorcycleService struct{}

func newMotorcycleService() *motorcycleService {
	return &motorcycleService{}
}

func (ms *motorcycleService) createMotorcycle(newMotorcycle *model.MotorcycleDTO) error {
	if newMotorcycle == nil {
		return fmt.Errorf("motorcycle can't be null")
	}
	if err := rpc_client.Call(addrs.MotorcycleAddr, "MotorcycleService.CreateMotorcycle", &newMotorcycle, &struct{}{}); err != nil {
		return err
	}
	return nil
}

func (ms *motorcycleService) updateMotorcycle(motorcycle *model.MotorcycleDTO) error {
	if motorcycle == nil {
		return fmt.Errorf("motorcycle can't be null")
	}
	if err := rpc_client.Call(addrs.MotorcycleAddr, "MotorcycleService.UpdateMotorcycle", &motorcycle, &struct{}{}); err != nil {
		return err
	}
	return nil
}

func (ms *motorcycleService) getAllMotorcycles() ([]*model.MotorcycleDTO, error) {
	var motorcycles []*model.MotorcycleDTO
	if err := rpc_client.Call(addrs.MotorcycleAddr, "MotorcycleService.GetAllMotorcycles", &struct{}{}, &motorcycles); err != nil {
		return nil, err
	}
	return motorcycles, nil
}

func (ms *motorcycleService) getAllActiveMotorcycles() ([]*model.MotorcycleDTO, error) {
	var motorcycles []*model.MotorcycleDTO
	if err := rpc_client.Call(addrs.MotorcycleAddr, "MotorcycleService.GetAllActiveMotorcycles", &struct{}{}, &motorcycles); err != nil {
		return nil, err
	}
	return motorcycles, nil
}

func (ms *motorcycleService) getMotorcycleById(motorcycleId int64) (*model.MotorcycleDTO, error) {
	var motorcycle *model.MotorcycleDTO
	if err := rpc_client.Call(addrs.MotorcycleAddr, "MotorcycleService.GetMotorcycleById", &motorcycleId, &motorcycle); err != nil {
		return nil, err
	}
	return motorcycle, nil
}

func (ms *motorcycleService) getMotorcycleByYear(motorcycleId int64) ([]*model.MotorcycleDTO, error) {
	var motorcycles []*model.MotorcycleDTO
	if err := rpc_client.Call(addrs.MotorcycleAddr, "MotorcycleService.GetAllMotorcycleByYear", &motorcycleId, &motorcycles); err != nil {
		return nil, err
	}
	return motorcycles, nil
}

func (ms *motorcycleService) getMotorcycleByModel(motorcycleId int64) ([]*model.MotorcycleDTO, error) {
	var motorcycles []*model.MotorcycleDTO
	if err := rpc_client.Call(addrs.MotorcycleAddr, "MotorcycleService.GetAllMotorcycleByModel", &motorcycleId, &motorcycles); err != nil {
		return nil, err
	}
	return motorcycles, nil
}

func (ms *motorcycleService) locateMotorcycle(motorcycleId int64) error {
	if err := rpc_client.Call(addrs.MotorcycleAddr, "MotorcycleService.LocateMotorcycle", &motorcycleId, &struct{}{}); err != nil {
		return err
	}
	return nil
}

func (ms *motorcycleService) unlocateMotorcycle(motorcycleId int64) error {
	if err := rpc_client.Call(addrs.MotorcycleAddr, "MotorcycleService.UnlocateMotorcycle", &motorcycleId, &struct{}{}); err != nil {
		return err
	}
	return nil
}
