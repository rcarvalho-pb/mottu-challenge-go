package sqlite

import (
	"context"

	"github.com/rcarvalho-pb/mottu-location_service/internal/model"
)

// UserId        int64     `json:"user_id" db:"user_id"`
// MotorcycleId  int64     `json:"motorcycle_id" db:"motorcycle_id"`
// Price         float64   `json:"price" db:"price"`
// Days          int64     `json:"days" db:"days"`
// LocationDay   time.Time `json:"location_day" db:"location_day"`
// DevolutionDay time.Time `json:"devolution_day" db:"devolution_day"`
// Fine          float64   `json:"fine" db:"fine"`

func (db *DB) CreateLocation(location *model.Location) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `INSERT INTO tb_locations (user_id, motorcycle_id, price, days, location_day, devolution_day, fine)
	VALUES 
	(:user_id, :motorcycle_id, :price, :days, :location_day, :devolution_day, :fine)`

	if _, err := db.db.NamedExecContext(ctx, stmt, &location); err != nil {
		return err
	}

	return nil
}

func (db *DB) UpdateLocation(location *model.Location) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `UPDATE tb_locations
	SET price = :price, days = :days, devolution_day = :devolution_day, fine = :fine
	WHERE user_id = :user_id AND motorcycle_id = :motorcycle_id AND active = TRUE`

	if _, err := db.db.NamedExecContext(ctx, stmt, &location); err != nil {
		return err
	}

	return nil
}

func (db *DB) GetLocationsByUserId(userId int64) ([]*model.Location, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `SELECT * FROM tb_locations WHERE user_id = :user_id AND active = TRUE`

	var locations []*model.Location

	if err := db.db.SelectContext(ctx, &locations, stmt, userId); err != nil {
		return nil, err
	}

	return locations, nil
}

func (db *DB) GetLocationsByMotorcycleId(motorcycleId int64) ([]*model.Location, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `SELECT * FROM tb_locations WHERE motorcycle_id = :motorcycle_id AND active = TRUE`

	var locations []*model.Location

	if err := db.db.SelectContext(ctx, &locations, stmt, motorcycleId); err != nil {
		return nil, err
	}

	return locations, nil
}

func (db *DB) GetAllActiveLocations() ([]*model.Location, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `SELECT * FROM tb_locations WHERE active = TRUE`

	var locations []*model.Location

	if err := db.db.SelectContext(ctx, &locations, stmt); err != nil {
		return nil, err
	}

	return locations, nil

}

func (db *DB) GetAllLocations() ([]*model.Location, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `SELECT * FROM tb_locations`

	var locations []*model.Location

	if err := db.db.SelectContext(ctx, &locations, stmt); err != nil {
		return nil, err
	}

	return locations, nil
}
