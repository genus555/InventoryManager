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
			if db.TableName == "" {return fmt.Errorf("No category is currently open.")}
			err := db.ListEntries(db.TableName)
			if err != nil {return err}
		} else {
			tableName, err := db.CheckTable(inputs[2])
			if err != nil {return err}
			err = db.ListEntries(tableName)
			if err != nil {return err}
		}
	} else {return fmt.Errorf("Unknown flag")}

	return nil
}

func HandleOpen(db *database.DB, inputs []string) error {
	if len(inputs) < 2 {return fmt.Errorf("Incorrect usage. Usage: open [category_name]")}
	table, err := db.CheckTable(inputs[1])
	if err != nil {return err}
	db.TableName = table
	fmt.Printf("Category \"%s\" has been opened.\n", db.TableName)
	return nil
}

func HandleCreateEntry(db *database.DB, inputs []string) error {
	return nil
}