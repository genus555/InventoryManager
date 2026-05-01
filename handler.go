package main

import (
	"fmt"

	database	"github.com/genus555/InventoryManager/internal/database"
)

func HandleCreateTable(db *database.DB, inputs []string) error {
	if len(inputs) < 2 {
		return fmt.Errorf("Incorrect usage\nUsage: new (or 'n') [category_name]")
	}

	err := db.CreateTable(inputs)
	if err != nil {return err}
	return nil
}

func HandleList(db *database.DB, inputs []string) error {
	if len(inputs) == 1 || (len(inputs) >= 2 && (inputs[1] == "-c" || inputs[1] == "-categories")) {
		err := db.ListTables()
		if err != nil {return err}
		return nil
	}
	if inputs[1] == "-e" || inputs[1] == "-entries" {
		if len(inputs) == 2 {
			fmt.Println("Add way to check open table and list entries")
		} else {
			fmt.Printf("Add way to check if table \"%s\" exists and list entries\n", inputs[2])
		}
	} else {return fmt.Errorf("Unknown flag")}

	return nil
}