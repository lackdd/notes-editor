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

    // Initialize a scanner to read through the file line by line
    scanner := bufio.NewScanner(file)

    // Print the header for the notes section only once
    fmt.Printf("\n\033[1m\033[4mNotes:\033[0m\n")

    // Read and decrypt each note from the file
    var decryptedNotes []string
    for scanner.Scan() {
        encryptedNote := scanner.Text()
        decryptedNote := rot13(encryptedNote) // Decrypt the note using ROT13
        decryptedNotes = append(decryptedNotes, decryptedNote)
    }

    // Check if any error occurred while scanning the file
    if err := scanner.Err(); err != nil {
        fmt.Println("Failed to read file.")
        return fmt.Errorf("failed to read file: %v", err)
    }

    // If no notes are found, notify the user and return
    if len(decryptedNotes) == 0 {
        fmt.Println("\033[31mNo notes to delete.\033[0m")
        return nil
    }

    // Display all decrypted notes with their corresponding numbers
    for i, note := range decryptedNotes {
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
    if err != nil || index < 0 || index > len(decryptedNotes) {
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

    // Delete the selected note from the slice of decrypted notes
    decryptedNotes = append(decryptedNotes[:index-1], decryptedNotes[index:]...)

    // Open the file for writing to save the updated notes (encrypted)
    file, err = os.Create(fileName)
    if err != nil {
        fmt.Println("Failed to open file for writing.")
        return fmt.Errorf("failed to open file for writing: %v", err)
    }
    defer file.Close()

    // Encrypt and write the remaining notes back to the file
    for _, note := range decryptedNotes {
        encryptedNote := rot13(note) // Encrypt the note before writing
        _, err = file.WriteString(encryptedNote + "\n")
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