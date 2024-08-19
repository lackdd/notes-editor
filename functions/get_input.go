package notes

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetInput() string {
	// Create a new scanner to read user input from the standard input (stdin).
	scanner := bufio.NewScanner(os.Stdin)

	// Enter a loop to continually prompt the user until valid input is received.
	for {
		// Print the menu options to the user.
		fmt.Printf("\nPlease choose an operation:\n1. Show notes.\n2. Add a note.\n3. Delete a note.\n4. Exit.\n")
		
		// Read the user's input.
		scanner.Scan()
		// Trim any leading or trailing whitespace from the input.
		operation := strings.TrimSpace(scanner.Text())
		
		// Check if the user's input is one of the valid options ("1", "2", "3", or "4").
		if operation == "1" || operation == "2" || operation == "3" || operation == "4" {
			// If valid, return the user's choice.
			return operation
		} else {
			// If invalid, clear the terminal and display an error message.
			ClearTerminal()
			fmt.Println("\033[31mInvalid input:\033[0m ", operation)
		}
	}
}