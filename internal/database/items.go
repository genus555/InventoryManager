package database

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

type DB struct {
	database	*sql.DB
	TableName	string
}

const (
	PLUS = true
	MINUS = false
)

const (
	LOW = 1
)

func NewDB(db *sql.DB) *DB {
	database := DB{
		database:	db,
		TableName:	"",
	}
	return &database
}