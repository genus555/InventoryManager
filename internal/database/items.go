package database

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

type DB struct {
	database	*sql.DB
	TableName	string
}

func NewDB(db *sql.DB) *DB {
	database := DB{
		database:	db,
		TableName:	"",
	}
	return &database
}