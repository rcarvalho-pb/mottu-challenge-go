package model

import "time"

type MotorcycleRepository interface {
	GetMotorcycleById(int64) (*Motorcycle, error)
	GetMotorcycleByUserId(int64) (*Motorcycle, error)
	GetAllMotorcycles() ([]*Motorcycle, error)
	GetAllActiveMotorcycle() ([]*Motorcycle, error)
	GetMotorcyclesByYear(int64) ([]*Motorcycle, error)
	GetMotorcycleByModel(string) ([]*Motorcycle, error)
	CreateMotorcycle(Motorcycle) error
	UpdateMotorcycle(Motorcycle) error
	DeleteMotorcycleById(int64) error
	LocateMotorcycle(int64) error
	UnlocateMotorcycle(int64) error
}

type Motorcycle struct {
	Id        int64     `json:"id" db:"id"`
	UserId    int64     `json:"user_id" db:"user_id"`
	Year      int64     `json:"year" db:"year"`
	Model     string    `json:"model" db:"model"`
	Plate     string    `json:"plate" db:"plate"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	IsLocated bool      `json:"is_located" db:"is_located"`
	Active    bool      `json:"active" db:"active"`
}

func (m *Motorcycle) UpdateTime() {
	m.UpdatedAt = time.Now()
}
