package sqlite

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	DB *sqlx.DB
}

func GetDB(dbLocation string) *DB {
	db := connectToDB(dbLocation)
	if db == nil {
		log.Fatal("couldn't connect to DB")
	}

	return &DB{db}
}

func connectToDB(dbLocation string) *sqlx.DB {
	count := 0
	for count < 10 {
		db, err := openDB(dbLocation)
		if err == nil {
			return db
		}

		count++
		time.Sleep(1 * time.Second)
	}

	return nil
}

func openDB(dbLocation string) (*sqlx.DB, error) {
	conn, err := sqlx.Open("sqlite3", dbLocation)
	if err != nil {
		return nil, err
	}
	if err = conn.Ping(); err != nil {
		return nil, err
	}

	return conn, nil
}
