package main

import (
	"fmt"
	"notes/functions"
	"os"
)



func main() {
    // Check if exactly one command-line argument is provided and if the argument is not "help".
	if len(os.Args) != 2 || os.Args[1] == "help" {
        // If the argument is missing or "help", print usage instructions and exit.
		fmt.Printf("Hello, to use this tool please add exactly one argument as a collection name to the commandline as in './notestool (your_collection_name)'")
		return
	} else {
        // Verify the password before proceeding. If the password is incorrect, exit.
		if !notes.VerifyPassword() {
			fmt.Println("Invalid password.")
			return
		}

        // Print a welcome message with some formatting (bold and underlined text).
		fmt.Printf("\n\033[1m\033[4mHello and welcome to the notes tool!\033[0m\n")

        // Create the filename using the provided collection name argument.
		filename := fmt.Sprintf("%s.txt", os.Args[1])

        // Check if the file exists. If it doesn't exist, create a new file.
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			file, err := os.Create(filename)
			if err != nil {
                // If there's an error creating the file, print an error message and exit.
				fmt.Println("Error creating file.")
				fmt.Printf("Error creating file: %v\n", err)
				return
			}
            // Close the file after creating it.
			file.Close()
		}

        // Enter an infinite loop to interact with the user.
		for {
            // Get the user's choice of operation (show notes, add a note, delete a note, exit).
			answer := notes.GetInput()

            // Clear the terminal screen after each operation.
			notes.ClearTerminal()

            // Handle the user's input based on their choice.
			if answer == "1" {
                // Show notes if the user selected "1".
				notes.ShowNotes(filename)
			} else if answer == "2" {
                // Add a new note if the user selected "2".
				notes.AddANote(filename)
			} else if answer == "3" {
                // Delete an existing note if the user selected "3".
				notes.DeleteANote(filename)
			} else if answer == "4" {
                // Exit the program if the user selected "4".
				notes.ClearTerminal()
				os.Exit(0)
			}
		}
	}
}