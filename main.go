package main

import (
	"fmt"
	"log"
	"database/sql"

	_			"modernc.org/sqlite"
	database	"github.com/genus555/InventoryManager/internal/database"
)

func main() {
	fmt.Println("Welcome to Inventory Manager")
	db_file, err := sql.Open("sqlite", "./inventory.db")
	if err != nil {
		log.Fatalf("Problem creating database: %v", err)
	}
	defer db_file.Close()

	db := database.NewDB(db_file)
	fmt.Println(db)
}