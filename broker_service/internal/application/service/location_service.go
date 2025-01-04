package service

import (
	"fmt"

	"github.com/rcarvalho-pb/mottu-broker_service/internal/model"
	rpc_client "github.com/rcarvalho-pb/mottu-broker_service/internal/rpc/client"
)

type locationService struct{}

func newLocationService() *locationService {
	return &locationService{}
}

func (ls *locationService) getLocationById(locationId int64) (*model.LocationDTO, error) {
	var location *model.LocationDTO
	if err := rpc_client.Call(addrs.LocationAddr, "LocationService.GetLocationById", &locationId, &location); err != nil {
		return nil, err
	}
	return location, nil
}

func (ls *locationService) createLocation(newLocation *model.NewLocationDTO) error {
	if newLocation == nil {
		return fmt.Errorf("new location can't be null")
	}
	if err := rpc_client.Call(addrs.LocationAddr, "LocationService.CreateLocation", &newLocation, &struct{}{}); err != nil {
		return err
	}
	return nil
}

func (ls *locationService) updateLocation(location *model.LocationDTO) error {
	if location == nil {
		return fmt.Errorf("location can't be null")
	}
	if err := rpc_client.Call(addrs.LocationAddr, "LocationService.UpdateLocation", &location, &struct{}{}); err != nil {
		return err
	}
	return nil
}

func (ls *locationService) endLocation(locationId int64) error {
	if err := rpc_client.Call(addrs.LocationAddr, "LocationService.EndLocation", &locationId, &struct{}{}); err != nil {
		return err
	}
	return nil
}

func (ls *locationService) getAllLocations() ([]*model.LocationDTO, error) {
	var locations []*model.LocationDTO
	if err := rpc_client.Call(addrs.LocationAddr, "LocationService.GetAllLocations", &struct{}{}, &locations); err != nil {
		return nil, err
	}
	return locations, nil
}

func (ls *locationService) getAllActiveLocations() ([]*model.LocationDTO, error) {
	var locations []*model.LocationDTO
	if err := rpc_client.Call(addrs.LocationAddr, "LocationService.GetAllActiveLocations", &struct{}{}, &locations); err != nil {
		return nil, err
	}
	return locations, nil
}

func (ls *locationService) getLocationsByUserId(userId int64) ([]*model.LocationDTO, error) {
	var (
		locations []*model.LocationDTO
		err       error
	)
	if err = rpc_client.Call(addrs.LocationAddr, "LocationService.GetLocationsByUserId", &userId, &locations); err != nil {
		return locations, err
	}
	return locations, err
}

func (ls *locationService) getLocationsByMotorcycleId(motorcycleId int64) ([]*model.LocationDTO, error) {
	var (
		locations []*model.LocationDTO
		err       error
	)
	if err = rpc_client.Call(addrs.LocationAddr, "LocationService.GetLocationsByMotorcycleId", &motorcycleId, &locations); err != nil {
		return locations, err
	}
	return locations, err
}
