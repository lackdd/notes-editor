package notes

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ShowNotes(filename string) {
	// Attempt to open the file specified by the filename
	readFile, err := os.Open(filename)
	if err != nil {
		// If there's an error opening the file, print an error message and return
		fmt.Println("Error opening file", err)
		return
	}
	// Ensure the file is closed once the function finishes executing
	defer readFile.Close()

	// Initialize a scanner to read through the file line by line
	fileScanner := bufio.NewScanner(readFile)
	// Split the file content by lines (default behavior, but made explicit here)
	fileScanner.Split(bufio.ScanLines)

	// Print the header for the notes section
	fmt.Printf("\n\033[1m\033[4mNotes:\033[0m\n")
	
	// Loop through each line in the file
	for fileScanner.Scan() {
		// Read the encrypted note
		encryptedNote := fileScanner.Text()
		// Decrypt the note using ROT13
		decryptedNote := rot13(encryptedNote)
		// Print the decrypted note
		fmt.Println(decryptedNote)
	}

	// Check if there was any error during scanning (e.g., read error)
	if err := fileScanner.Err(); err != nil {
		// If an error occurred while scanning, print an error message
		fmt.Println("Error reading file", err)
	}

	// Create a new scanner to wait for the user's input to continue
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("\n\033[33mPress Enter to continue\033[0m")
	
	// Wait for the user to press Enter to continue
	for {
        scanner.Scan() 
        operation := strings.TrimSpace(scanner.Text())
        if operation == "" {
			// If the user presses Enter without any additional input, clear the terminal and return
			ClearTerminal()
            return 
        } else {
			// If the user types something instead of just pressing Enter, prompt them to only press Enter
			fmt.Println("\033[31mInvalid input. Please just press Enter to continue.\033[0m")
		}
    }
}