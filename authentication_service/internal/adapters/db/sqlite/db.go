package sqlite

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rcarvalho-pb/mottu-authentication_service/internal/application/dtos"
	"github.com/rcarvalho-pb/mottu-authentication_service/internal/domain/model"
)

type DB struct {
	DB *sql.DB
}

func GetDB() *DB {
	db := ConnectToDB()
	if db == nil {
		log.Fatal("couldn't connect to DB")
	}

	return &DB{db}
}

func ConnectToDB() *sql.DB {
	count := 0
	for count < 10 {
		db, err := OpenDB()
		if err == nil {
			return db
		}

		count++
		time.Sleep(1 * time.Second)
	}

	return nil
}

func OpenDB() (*sql.DB, error) {
	conn, err := sql.Open("sqlite3", "../data-storage/db.db")
	if err != nil {
		return nil, err
	}
	if err = conn.Ping(); err != nil {
		return nil, err
	}

	return conn, nil
}

func (db DB) FindUserByUsername(req dtos.UserRequest) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stmt := "select * from tb_users where username = ?"

	row := db.DB.QueryRowContext(ctx, stmt, req.Username)

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
	); err != nil {
		return nil, err
	}

	return &user, nil
}
