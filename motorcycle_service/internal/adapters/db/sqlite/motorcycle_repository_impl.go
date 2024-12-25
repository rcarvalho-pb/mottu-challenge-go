package sqlite

import (
	"context"
	"time"

	"github.com/rcarvalho-pb/mottu-motorcycle_service/internal/model"
)

// GetAllMotorcycles() ([]*Motorcycle, error)
// GetAllActiveMotorcycle() ([]*Motorcycle, error)
// GetMotorcyclesByYear(int64) ([]*Motorcycle, error)
// GetMotorcycleByModel(string) ([]*Motorcycle, error)

// user_id INTEGER,
// year INT NOT NULL,
// model TEXT NOT NULL,
// plate TEXT NOT NULL,
// created_at TIMESTAMP DEFAUL CURRENT_TIMESTAMP,
// updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
// is_located BOOLEAN DEFAULT FALSE,

var dbTimeout = 10 * time.Second

func (db *DB) GetAllMotorcycles() ([]*model.Motorcycle, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `SELECT * FROM tb_motorcycle`

	var motorcycles []*model.Motorcycle

	db.DB.SelectContex(ctx, motorcycles, stmt)
}

func (db *DB) UpdateMotorcycle(motorcycle *model.Motorcycle) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `UPDATE tb_motorcycles
	SET year = :year, model = :model, plate = :plate
	WHERE id = :id`

	motorcycle.UpdateTime()

	if _, err := db.DB.NamedExecContext(ctx, stmt, motorcycle); err != nil {
		return err
	}

	return nil
}

func (db *DB) LocateMotorcycle(motorcycleId, userId int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	motorcycle, err := db.GetMotorcycleById(motorcycleId)
	if err != nil {
		return err
	}

	stmt := `UPDATE tb_motorcycles 
	SET user_id = :user_id, updated_at = :updated_at, is_located = :is_located 
	WHERE id = :id`

	motorcycle.UserId = userId
	motorcycle.UpdateTime()
	motorcycle.IsLocated = true

	tx := db.DB.MustBegin()
	tx.NamedExecContext(ctx, stmt, motorcycle)
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (db *DB) UnlocateMotorcycle(userId int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	motorcycle, err := db.GetMotorcycleByUserId(userId)
	if err != nil {
		return err
	}

	motorcycle.UserId = 0
	motorcycle.UpdateTime()
	motorcycle.IsLocated = false

	stmt := `UPDATE tb_motorcycles 
	SET user_id = :user_id, updated_at = :updated_at, is_located = :is_located 
	WHERE id = :id`

	tx := db.DB.MustBegin()
	tx.NamedExecContext(ctx, stmt, motorcycle)
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (db *DB) GetMotorcycleById(motorcycleId int64) (*model.Motorcycle, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `SELECT * FROM tb_motorcycles WHERE id = ?`

	var motorcycle *model.Motorcycle

	if err := db.DB.GetContext(ctx, motorcycle, stmt, motorcycleId); err != nil {
		return nil, err
	}

	return motorcycle, nil
}

func (db *DB) CreateMotorcycle(motorcycle model.Motorcycle) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `INSERT INTO tb_motorcycles (year, model, plate) VALUES (:year, :model, :plate)`

	if _, err := db.DB.NamedExecContext(ctx, stmt, &motorcycle); err != nil {
		return err
	}

	return nil
}

func (db *DB) DeleteMotorcycleById(motorcycleId int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	motorcycle, err := db.GetMotorcycleById(motorcycleId)
	if err != nil {
		return err
	}

	motorcycle.Active = false
	motorcycle.UpdateTime()

	stmt := `UPDATE tb_motorcycles
	SET active = :active, updated_at = :updated_at
	WHERE id = :id`

}

func (db *DB) GetMotorcycleByUserId(motorcycleId int64) (*model.Motorcycle, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `SELECT * FROM tb_motorcycles WHERE id = ?`

	var motorcycle model.Motorcycle

	if err := db.DB.GetContext(ctx, &motorcycle, stmt, motorcycleId); err != nil {
		return nil, err
	}

	return &motorcycle, nil
}
