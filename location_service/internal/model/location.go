package model

import "time"

type Location struct {
	UserId        int64     `db:"user_id"`
	MotorcycleId  int64     `db:"motorcycle_id"`
	Price         float64   `db:"price"`
	Days          int64     `db:"days"`
	LocationDay   time.Time `db:"location_day"`
	DevolutionDay time.Time `db:"devolution_day"`
	Fine          float64   `db:"fine"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}
