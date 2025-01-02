package model

import "time"

type LocationDTO struct {
	Id            int64     `json:"id"`
	UserId        int64     `json:"user_id"`
	MotorcycleId  int64     `json:"motorcycle_id"`
	Price         float64   `json:"price"`
	Days          int64     `json:"days"`
	LocationDay   time.Time `json:"location_day"`
	DevolutionDay time.Time `json:"devolution_day"`
	Fine          float64   `json:"fine"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type NewLocationDTO struct {
	Id            int64     `json:"id"`
	UserId        int64     `json:"user_id"`
	MotorcycleId  int64     `json:"motorcycle_id"`
	Price         float64   `json:"price"`
	Days          int64     `json:"days"`
	LocationDay   time.Time `json:"location_day"`
	DevolutionDay time.Time `json:"devolution_day"`
	Fine          float64   `json:"fine"`
}
