package sqlite

import (
	"context"
	"time"

	"github.com/rcarvalho-pb/mottu-motorcycle_service/internal/application/dtos"
	"github.com/rcarvalho-pb/mottu-motorcycle_service/internal/model"
)

var dbTimeout = 10 * time.Second

func (db *DB) GetMotorcyclesByModel(motorcycleModel string) ([]*model.Motorcycle, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `SELECT * FROM tb_motorcycle WHERE active = true and model = ?`

	var motorcycles []*model.Motorcycle

	if err := db.DB.SelectContext(ctx, &motorcycles, stmt, motorcycleModel); err != nil {
		return nil, err
	}

	return motorcycles, nil
}

func (db *DB) GetMotorcyclesByYear(year int64) ([]*model.Motorcycle, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `SELECT * FROM tb_motorcycle WHERE active = true and year = ?`

	var motorcycles []*model.Motorcycle

	if err := db.DB.SelectContext(ctx, &motorcycles, stmt); err != nil {
		return nil, err
	}

	return motorcycles, nil
}

func (db *DB) GetAllMotorcycles() ([]*model.Motorcycle, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `SELECT * FROM tb_motorcycle`

	var motorcycles []*model.Motorcycle

	if err := db.DB.SelectContext(ctx, &motorcycles, stmt); err != nil {
		return nil, err
	}

	return motorcycles, nil
}

func (db *DB) GetAllActiveMotorcycles() ([]*model.Motorcycle, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `SELECT * FROM tb_motorcycle WHERE active = true`

	var motorcycles []*model.Motorcycle

	if err := db.DB.SelectContext(ctx, &motorcycles, stmt); err != nil {
		return nil, err
	}

	return motorcycles, nil
}

func (db *DB) UpdateMotorcycle(motorcycle *model.Motorcycle) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `UPDATE tb_motorcycles
	SET year = :year, model = :model, plate = :plate, updated_at = :updated_at, is_located = :is_located, active = :active
	WHERE id = :id`

	if _, err := db.DB.NamedExecContext(ctx, stmt, motorcycle); err != nil {
		return err
	}

	return nil
}

func (db *DB) LocateMotorcycle(motorcycleId int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	motorcycle, err := db.GetMotorcycleById(motorcycleId)
	if err != nil {
		return err
	}

	stmt := `UPDATE tb_motorcycles 
	SET updated_at = :updated_at, is_located = :is_located 
	WHERE id = :id`

	motorcycle.UpdateTime()
	motorcycle.IsLocated = true

	tx := db.DB.MustBegin()
	tx.NamedExecContext(ctx, stmt, motorcycle)
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (db *DB) UnlocateMotorcycle(motorcycleId int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	motorcycle, err := db.GetMotorcycleById(motorcycleId)
	if err != nil {
		return err
	}

	motorcycle.UpdateTime()
	motorcycle.IsLocated = false

	stmt := `UPDATE tb_motorcycles 
	SET updated_at = :updated_at, is_located = :is_located 
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

	stmt := `SELECT * FROM tb_motorcycles WHERE active = true AND id = ?`

	var motorcycle *model.Motorcycle

	if err := db.DB.GetContext(ctx, motorcycle, stmt, motorcycleId); err != nil {
		return nil, err
	}

	return motorcycle, nil
}

func (db *DB) GetMotorcycleByIdAdmin(motorcycleId int64) (*model.Motorcycle, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `SELECT * FROM tb_motorcycles WHERE id = ?`

	var motorcycle *model.Motorcycle

	if err := db.DB.GetContext(ctx, motorcycle, stmt, motorcycleId); err != nil {
		return nil, err
	}

	return motorcycle, nil
}

func (db *DB) CreateMotorcycle(request *dtos.NewMotorcycleRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `INSERT INTO tb_motorcycles (year, model, plate) VALUES (:year, :model, :plate)`

	if _, err := db.DB.NamedExecContext(ctx, stmt, &request); err != nil {
		return err
	}

	return nil
}
