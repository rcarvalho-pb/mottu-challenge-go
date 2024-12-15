package sqlite

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
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
