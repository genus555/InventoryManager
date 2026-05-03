package cli

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func GetInput() []string {
	fmt.Printf("> ")
	scanner := bufio.NewScanner(os.Stdin)
	scanned := scanner.Scan()
	if !scanned {return nil}
	input := scanner.Text()
	input = strings.TrimSpace(input)
	input = strings.ToLower(input)
	return strings.Fields(input)
}

func PrintCommands() {
	fmt.Println(`List of available commands:
	*new: Creates a new category in database.
	   Usage: new [category_name]
	*list: Lists categories or entries in a category
	   - categories: Default option or with flag
	       Usage: List -categories (or '-c')
	   - category entries: with flag lists all entries in [category_name]
	       Usage: List -entries (or '-e') [category_name]
	*open: Opens a category
	   Usage: open [category_name]
	*add: Adds an entry to the currently open category
	   Usage: add [entry_name]
	*delete: Deletes an entry
	   - default: deletes an entry from the currently open category
		   Usage: delete [entry_name]
	   - delete category: deletes the category from database
	       Usage: delete -t [category_name]
	*get: Gets an entry from the currently open category
	   Usage: get [entry_name]
	*update: Updates the amount of an entry from the currently open category
	   Usage: update [entry_name] [new_amount]
	*plus/minus: Adds or subtracts 1 amount from an entry from the currently open category
	   Usage: plus/minus [entry_name]
	*restock: lists entries that need restocking in currently open category
	   - default: lists low and out of stock entries
	       Usage: restock
	   - low: lists entries that are low on stock
	       Usage: restock -low (or '-l')
	   - empty: lists entries that are out of stock
	       Usage: restock -empty (or '-e')
	*help: Shows available commands
	   Usage: help
	*quit: Stops the program
	   Usage: quit
	`)
}