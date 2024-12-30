package dtos

import "time"

type NewLocationDTO struct {
	UserId        int64     `json:"user_id" db:"user_id"`
	MotorcycleId  int64     `json:"motorcycle_id" db:"motorcycle_id"`
	Price         float64   `json:"price" db:"price"`
	Days          int64     `json:"days" db:"days"`
	LocationDay   time.Time `json:"location_day" db:"location_day"`
	DevolutionDay time.Time `json:"devolution_day" db:"devolution_day"`
	Fine          float64   `json:"fine" db:"fine"`
}
