package service

import "github.com/rcarvalho-pb/mottu-location_service/internal/adapters/db/sqlite"

type LocationService struct {
	db *sqlite.DB
}

func New(db *sqlite.DB) *LocationService {
	return &LocationService{
		db: db,
	}
}
