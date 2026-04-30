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
	fmt.Println("List of available commands:")
	fmt.Println("	*help:\n	     Shows available commands\n	     Usage: help")
	fmt.Println("	*quit:\n	     Stops the program\n	     Usage: quit")
}