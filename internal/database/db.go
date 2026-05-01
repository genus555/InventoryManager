package database

import (
	//"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

func (db *DB) CreateTable(inputs []string) error {
	tableParams := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		amount INTEGER DEFAULT 0
		)
	`, inputs[1])
	_, err:= db.database.Exec(tableParams)
	if err != nil {fmt.Errorf("Problem creating database")}

	return nil
}

func (db *DB) ListTables() error {
	rows, err := db.database.Query("SELECT name FROM sqlite_master WHERE type='table' AND name NOT LIKE 'sqlite_%';")
	if err != nil {return err}
	defer rows.Close()

	fmt.Println("Categories:")
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			return err
		}
		fmt.Printf("    - %s\n", tableName)
	}
	if err := rows.Err(); err != nil {return err}
	return nil
}