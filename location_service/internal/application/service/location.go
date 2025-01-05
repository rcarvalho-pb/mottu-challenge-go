package service

import (
	"github.com/rcarvalho-pb/mottu-location_service/internal/adapters/db/sqlite"
	"github.com/rcarvalho-pb/mottu-location_service/internal/application/dtos"
	"github.com/rcarvalho-pb/mottu-location_service/internal/model"
	rpc_client "github.com/rcarvalho-pb/mottu-location_service/internal/rpc/client"
)

type LocationService struct {
	db *sqlite.DB
}

func New(db *sqlite.DB) *LocationService {
	return &LocationService{
		db: db,
	}
}

func (ls *LocationService) CreateLocation(dto *dtos.NewLocationDTO) (err error) {
	var user *dtos.UserDTO
	if err = rpc_client.Call("", "UserService.GetUserById", &dto.UserId, &user); err != nil {
		return
	}
	newLocation := &model.Location{
		UserId:        dto.UserId,
		MotorcycleId:  dto.MotorcycleId,
		Price:         dto.Price,
		Days:          dto.Days,
		LocationDay:   dto.LocationDay,
		DevolutionDay: dto.DevolutionDay,
	}

	if err := ls.db.CreateLocation(newLocation); err != nil {
		return err
	}

	return nil
}

func (ls *LocationService) UpdateLocation(dto *dtos.LocationDTO) error {
	location := model.LocationFromDTO(dto)

	if err := ls.db.UpdateLocation(location); err != nil {
		return err
	}

	return nil
}

func (ls *LocationService) EndLocation(locationId int64) error {
	if err := ls.db.EndLocation(locationId); err != nil {
		return err
	}

	return nil
}

func (ls *LocationService) GetAllLocations() ([]*dtos.LocationDTO, error) {
	locations, err := ls.db.GetAllLocations()
	if err != nil {
		return nil, err
	}

	locationsDTO := make([]*dtos.LocationDTO, len(locations))
	for i, l := range locations {
		locationsDTO[i] = l.ToDTO()
	}

	return locationsDTO, nil
}

func (ls *LocationService) GetAllActiveLocations() ([]*dtos.LocationDTO, error) {
	locations, err := ls.db.GetAllActiveLocations()
	if err != nil {
		return nil, err
	}

	locationsDTO := make([]*dtos.LocationDTO, len(locations))
	for i, l := range locations {
		locationsDTO[i] = l.ToDTO()
	}

	return locationsDTO, nil
}

func (ls *LocationService) GetLocationById(locationsId int64) (*dtos.LocationDTO, error) {
	location, err := ls.db.GetLocationById(locationsId)
	if err != nil {
		return nil, err
	}

	return location.ToDTO(), nil
}

func (ls *LocationService) GetLocationByUserId(userId int64) ([]*dtos.LocationDTO, error) {
	locations, err := ls.db.GetLocationsByUserId(userId)
	if err != nil {
		return nil, err
	}

	locationsDTO := make([]*dtos.LocationDTO, len(locations))

	for i, l := range locations {
		locationsDTO[i] = l.ToDTO()
	}
	return locationsDTO, nil
}

func (ls *LocationService) GetLocationByMotorcycleId(motorcycleId int64) ([]*dtos.LocationDTO, error) {
	locations, err := ls.db.GetLocationsByMotorcycleId(motorcycleId)
	if err != nil {
		return nil, err
	}

	locationsDTO := make([]*dtos.LocationDTO, len(locations))

	for i, l := range locations {
		locationsDTO[i] = l.ToDTO()
	}
	return locationsDTO, nil
}
