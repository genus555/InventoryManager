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
	query := fmt.Sprintf("SELECT amount FROM %s WHERE name = ?", tableName)
	err := db.database.QueryRow(query, item).Scan(&amount)
	if err == sql.ErrNoRows {return 0, fmt.Errorf("Item \"%s\" in category \"%s\" doesn't exist.", item, tableName)} else if err != nil {return 0, err}
	return amount, nil
}

func (db *DB) AddEntry(entry string) error {
	var exists string
	query := fmt.Sprintf("SELECT name FROM %s WHERE name = ?", db.TableName)
	if err := db.database.QueryRow(query, entry).Scan(&exists); err == nil {
		return fmt.Errorf("Entry \"%s\" already exists in category \"%s\".", entry, db.TableName)
	}

	query = fmt.Sprintf("INSERT INTO %s (name) VALUES (?)", db.TableName)
	_, err := db.database.Exec(query, entry)
	if err != nil {return err}
	return nil
}

func (db *DB) DeleteEntry(entry string) error {
	var exists string
	query := fmt.Sprintf("SELECT name FROM %s WHERE name = ?", db.TableName)
	if err := db.database.QueryRow(query, entry).Scan(&exists); err != nil {
		if err == sql.ErrNoRows{
			return fmt.Errorf("Entry \"%s\" doesn't exists in category \"%s\".", entry, db.TableName)
		}
		return err
	}

	query = fmt.Sprintf("DELETE FROM %s WHERE name = ?", db.TableName)
	_, err := db.database.Exec(query, entry)
	if err != nil {return err}

	fmt.Printf("Entry \"%s\" has been deleted from category \"%s\".\n", entry, db.TableName)
	return nil
}

func (db *DB) GetEntry(entry string) (int, error) {
	var amount int
	query := fmt.Sprintf("SELECT amount FROM %s WHERE name = ?", db.TableName)
	err := db.database.QueryRow(query, entry).Scan(&amount)
	if err == sql.ErrNoRows {
		return 0, fmt.Errorf("\"%s\" in \"%s\" doesn't exist.", entry, db.TableName)
	} else if err != nil {return 0, err}
	return amount, nil
}

func (db *DB) UpdateEntry(entry string, amount int) error {
	var exists string
	query := fmt.Sprintf("SELECT name FROM %s WHERE name = ?", db.TableName)
	if err := db.database.QueryRow(query, entry).Scan(&exists); err != nil {
		if err == sql.ErrNoRows{
			return fmt.Errorf("Entry \"%s\" doesn't exists in category \"%s\".", entry, db.TableName)
		}
		return err
	}

	query = fmt.Sprintf("UPDATE %s SET amount = ? WHERE name = ?", db.TableName)
	_, err := db.database.Exec(query, amount, entry)
	if err != nil {return err}

	fmt.Printf("Entry \"%s\" has been updated.\n", entry)
	return nil
}

func (db *DB) PlusMinus(entry string, PM bool) error {
	initial_amount, err := db.GetEntry(entry)
	if err != nil {return err}

	amount := initial_amount
	if PM {amount += 1} else {amount -= 1}

	query := fmt.Sprintf("UPDATE %s SET amount = ? WHERE name = ?", db.TableName)
	_, err = db.database.Exec(query, amount, entry)
	if err != nil {return err}

	fmt.Printf("Entry \"%s\" in %s has been updated: %d --> %d\n", entry, db.TableName, initial_amount, amount)
	return nil
}