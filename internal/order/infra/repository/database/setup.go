package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func Load() *sql.DB {
	db, err := sql.Open("sqlite3", "./orders.db")
	if err != nil {
		panic(err)
	}

	return db
}
