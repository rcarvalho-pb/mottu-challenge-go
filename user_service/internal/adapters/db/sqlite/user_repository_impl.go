package sqlite

import (
	"context"
	"fmt"

	"github.com/rcarvalho-pb/mottu-user_service/internal/domain/model"
)

func (db *DB) GetAllUsers() ([]*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `
	SELECT 
	id, username, password, role, name, birth_date, cnpj, cnh, cnh_type, cnh_file_path, active_location, created_at, updated_at, active
	FROM tb_users`

	rows, err := db.DB.QueryContext(ctx, stmt)
	if err != nil {
		return nil, err
	}

	var users []*model.User

	for rows.Next() {
		var user model.User
		if err = rows.Scan(
			&user.ID,
			&user.Username,
			&user.Password,
			&user.Role,
			&user.Name,
			&user.BirthDate,
			&user.CNPJ,
			&user.CNH,
			&user.CNHType,
			&user.CNHFilePath,
			&user.ActiveLocation,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.Active,
		); err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return users, err
}

func (db *DB) GetAllActiveUsers() ([]*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `
	SELECT 
	id, username, password, role, name, birth_date, cnpj, cnh, cnh_type, cnh_file_path, active_location, created_at, updated_at, active
	FROM tb_users
	WHERE active = true`

	rows, err := db.DB.QueryContext(ctx, stmt)
	if err != nil {
		return nil, err
	}

	var users []*model.User

	for rows.Next() {
		var user model.User
		if err = rows.Scan(
			&user.ID,
			&user.Username,
			&user.Password,
			&user.Role,
			&user.Name,
			&user.BirthDate,
			&user.CNPJ,
			&user.CNH,
			&user.CNHType,
			&user.CNHFilePath,
			&user.ActiveLocation,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.Active,
		); err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return users, err

}

func (db *DB) GetUserById(userId int64) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `SELECT * FROM tb_users WHERE id = ?`

	row := db.DB.QueryRowContext(ctx, stmt, userId)

	var user model.User
	if err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Role,
		&user.Name,
		&user.BirthDate,
		&user.CNPJ,
		&user.CNH,
		&user.CNHType,
		&user.CNHFilePath,
		&user.ActiveLocation,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Active,
	); err != nil {
		return nil, err
	}

	return &user, nil
}

func (db *DB) GetUserByUsername(username string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `SELECT * FROM tb_users WHERE username = ?`
	row := db.DB.QueryRowContext(ctx, stmt, username)

	user := new(model.User)

	if err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Role,
		&user.Name,
		&user.BirthDate,
		&user.CNPJ,
		&user.CNH,
		&user.CNHType,
		&user.CNHFilePath,
		&user.ActiveLocation,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Active,
	); err != nil {
		fmt.Printf("Error: %s\n", err)
		return nil, err
	}

	return user, nil
}

func (db *DB) GetUsersByName(name string) ([]*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `SELECT * FROM tb_users WHERE name LIKE ?`

	searchUsername := fmt.Sprintf("%%%s%%", name)
	rows, err := db.DB.QueryContext(ctx, stmt, searchUsername)
	if err != nil {
		return nil, err
	}

	var users []*model.User

	for rows.Next() {
		var user model.User
		if err = rows.Scan(
			&user.ID,
			&user.Username,
			&user.Password,
			&user.Role,
			&user.Name,
			&user.BirthDate,
			&user.CNPJ,
			&user.CNH,
			&user.CNHType,
			&user.CNHFilePath,
			&user.ActiveLocation,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.Active,
		); err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return users, err

}

func (db *DB) CreateUser(user *model.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `INSERT INTO tb_users (username, password, name, birth_date, cnpj, cnh, cnh_type, cnh_file_path) VALUES (:username, :password, :name, :birth_date, :cnpj, :cnh, :cnh_type, :cnh_file_path)`

	res, err := db.DB.NamedExecContext(ctx, stmt, user)
	if err != nil {
		return err
	}

	fmt.Println(res.RowsAffected())

	return nil
}

func (db *DB) UpdateUser(user *model.User) error {
	return nil
}
