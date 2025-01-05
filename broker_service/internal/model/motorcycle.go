package model

import "time"

type MotorcycleDTO struct {
	Id        int64     `json:"id"`
	Year      int64     `json:"year"`
	Model     string    `json:"model"`
	Plate     string    `json:"plate"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsLocated bool      `json:"is_located"`
	Active    bool      `json:"active"`
}
