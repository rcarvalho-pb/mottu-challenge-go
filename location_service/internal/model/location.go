package model

import (
	"time"

	"github.com/rcarvalho-pb/mottu-location_service/internal/application/dtos"
)

type LocationRepository interface {
	CreateLocation(*Location) error
	UpdateLocation(*Location) error
	EndLocation(int64) error
	GetAllLocations() ([]*Location, error)
	GetLocationById(int64) (*Location, error)
	GetAllActiveLocations() ([]*Location, error)
	GetLocationsByUserId(int64) ([]*Location, error)
	GetLocationsByMotorcycleId(int64) ([]*Location, error)
}

type Location struct {
	Id            int64     `db:"id"`
	UserId        int64     `db:"user_id"`
	MotorcycleId  int64     `db:"motorcycle_id"`
	Price         float64   `db:"price"`
	Days          int64     `db:"days"`
	LocationDay   time.Time `db:"location_day"`
	DevolutionDay time.Time `db:"devolution_day"`
	Fine          float64   `db:"fine"`
	Active        bool      `db:"active"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}

func (l *Location) ToDTO() *dtos.LocationDTO {
	return &dtos.LocationDTO{
		Id:            l.Id,
		UserId:        l.UserId,
		MotorcycleId:  l.MotorcycleId,
		Price:         l.Price,
		Days:          l.Days,
		LocationDay:   l.LocationDay,
		DevolutionDay: l.DevolutionDay,
		Fine:          l.Fine,
		CreatedAt:     l.CreatedAt,
		UpdatedAt:     l.UpdatedAt,
	}
}

func LocationFromDTO(dto *dtos.LocationDTO) *Location {
	return &Location{
		Id:            dto.Id,
		UserId:        dto.UserId,
		MotorcycleId:  dto.MotorcycleId,
		Price:         dto.Price,
		Days:          dto.Days,
		LocationDay:   dto.LocationDay,
		DevolutionDay: dto.DevolutionDay,
		Fine:          dto.Fine,
		CreatedAt:     dto.CreatedAt,
		UpdatedAt:     dto.UpdatedAt,
	}
}
