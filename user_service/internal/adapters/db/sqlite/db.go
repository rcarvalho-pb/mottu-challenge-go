package sqlite

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rcarvalho-pb/mottu-user_service/internal/domain/model"
)

// GetUserById(int64) (*User, error)
// GetUserByUsername(string) ([]*User, error)
// CreateUser(User) error
// UpdateUser(User) error

var dbTimeout = 10 * time.Second

type DB struct {
	DB *sqlx.DB
}

func GetDB() *DB {
	db := connectToDB()
	if db == nil {
		log.Fatal("couldn't connect to DB")
	}

	return &DB{db}
}

func connectToDB() *sqlx.DB {
	count := 0
	for count < 10 {
		db, err := openDB()
		if err == nil {
			return db
		}

		count++
		time.Sleep(1 * time.Second)
	}

	return nil
}

func openDB() (*sqlx.DB, error) {
	conn, err := sqlx.Open("sqlite3", "../data-storage/db.db")
	if err != nil {
		return nil, err
	}
	if err = conn.Ping(); err != nil {
		return nil, err
	}

	return conn, nil
}

func (db *DB) GetAllUsers() ([]*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `
	SELECT 
	id, username, password, name, birth_date, cnpj, cnh, cnh_type, cnh_file_path, active_location, created_at, updated_at, active
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
	id, username, password, name, birth_date, cnpj, cnh, cnh_type, cnh_file_path, active_location, created_at, updated_at, active
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

func (db *DB) GetUserByUsername(username string) ([]*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	searchedUsername := fmt.Sprintf("%s%%", username)
	stmt := `SELECT * FROM tb_users WHERE username LIKE ?`
	rows, err := db.DB.QueryContext(ctx, stmt, searchedUsername)
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
