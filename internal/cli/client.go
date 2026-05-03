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
	   - category entries: currently open category or with flag and [category_name]
	        Usage: List -entries (or '-e') [category_name]
	*open: Opens a category
	   Usage: open [category_name]
	*add: Adds an entry to the currently open category
	   Usage: add [entry_name]
	*delete: Deletes an entry from the currently open category
	   Usage: delete [entry_name]
	   - delete category: deletes the category from database
	      Usage: delete -t [category_name]
	*get: Gets an entry from the currently open category
	   Usage: get [entry_name]
	*update: Updates the amount of an entry from the currently open category
	   Usage: update [entry_name] [new_amount]
	*plus/minus: Adds or subtracts 1 amount from an entry from the currently open category
	   Usage: plus/minus [entry_name]
	*help: Shows available commands
	   Usage: help
	*quit: Stops the program
	   Usage: quit
	`)
}