package database

import (
	"database/sql"
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

func (db *DB) ListEntries(tableName string) error {
	query := fmt.Sprintf("SELECT name FROM %s", tableName)
	entries, err := db.database.Query(query)
	if err != nil {return err}
	defer entries.Close()
	fmt.Printf("Category %s:\n", tableName)

	for entries.Next() {
		var entry string
		if err := entries.Scan(&entry); err != nil {return err}
		amount, err := db.GetAmount(tableName, entry)
		if err != nil {return err}
		fmt.Printf("    - %s: %d\n", entry, amount)
	}
	return nil
}

func (db *DB) CheckTable(tableName string) (string, error) {
	var exists string
	if err := db.database.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name = ?", tableName).Scan(&exists); err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("Category \"%s\" doesn't exist", tableName)
		}
		return "", err
	}
	return exists, nil
}

func (db *DB) GetAmount(tableName, item string) (int, error) {
	var amount int
	query := fmt.Sprintf("SELECT amount FROM %s WHERE name = ?", tableName).Scan(&amount)
	err := db.database.QueryRow(query, item)
	if err == sql.ErrNoRows {return 0, fmt.Errorf("Item \"%s\" in category \"%s\" doesn't exist.", item, tableName)} else if err != nil {return 0, err}
	return amount, nil
}