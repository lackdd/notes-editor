package notes

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func DeleteANote(fileName string) error {
	// Attempt to open the file for reading
	file, err := os.Open(fileName)
	if err != nil {
		// If there's an error opening the file, print an error message and return
		fmt.Println("Failed to open file.")
		return fmt.Errorf("failed to open file: %v", err)
	}
	// Ensure the file is closed when the function exits, preventing resource leaks
	defer file.Close()

	// Read all notes from the file into a slice of strings
	var notes []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		notes = append(notes, scanner.Text())
	}

	// Check if any error occurred while scanning the file
	if err := scanner.Err(); err != nil {
		fmt.Println("Failed to read file.")
		return fmt.Errorf("failed to read file: %v", err)
	}

	// If no notes are found, notify the user and return
	if len(notes) == 0 {
		fmt.Println("\033[31mNo notes to delete.\033[0m")
		return nil
	}

	// Display all notes with their corresponding numbers
	fmt.Println("\033[1m\033[4mNotes:\033[0m")
	for i, note := range notes {
		fmt.Printf("%d: %s\n", i+1, note) // Print each note with its corresponding number
	}

	// Prompt the user to enter the number of the note to delete, or 0 to cancel
	fmt.Print("\n\033[33mEnter the number of note to delete or press 0 to cancel:\033[0m ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Invalid input.")
		return fmt.Errorf("invalid input: %v", err)
	}
	input = strings.TrimSpace(input)  // Remove any leading or trailing whitespace from the input
	index, err := strconv.Atoi(input) // Convert the string input to an integer
	if err != nil || index < 0 || index > len(notes) {
		// If the input is not a valid number or out of range, notify the user
		ClearTerminal()
		fmt.Println("\033[31mInvalid note number, note not deleted.\033[0m")
		return fmt.Errorf("invalid note number")
	}
	if index == 0 {
		// If the user entered 0, cancel the deletion
		fmt.Println("Note deletion cancelled.")
		return nil
	}

	// Delete the selected note from the slice
	notes = append(notes[:index-1], notes[index:]...)

	// Recreate the file to save the updated notes (without the deleted note)
	file, err = os.Create(fileName) // Open the file for writing
	if err != nil {
		fmt.Println("Failed to open file for writing.")
		return fmt.Errorf("failed to open file for writing: %v", err)
	}
	defer file.Close()

	// Write the remaining notes back to the file
	for _, note := range notes {
		_, err = file.WriteString(note + "\n")
		if err != nil {
			fmt.Println("Failed to write to file.")
			return fmt.Errorf("failed to write to file: %v", err)
		}
	}

	// Clear the terminal and notify the user that the note was successfully deleted
	ClearTerminal()
	fmt.Println("Note successfully deleted!")
	return nil
}