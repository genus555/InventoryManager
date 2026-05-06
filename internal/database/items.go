package database

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

type DB struct {
	database	*sql.DB
	TableName	string
	cache		map[string]int
	cacheOrder	[]string
}

const (
	PLUS = true
	MINUS = false
)

const (
	LOW = 1
	DEFAULTAMOUNT = 0
)

func NewDB(db *sql.DB) *DB {
	return &DB{
		database:	db,
		TableName:	"",
		cache:		make(map[string]int),
		cacheOrder:	[]string{},
	}
}