package notes

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func AddANote(fileName string) error {
    // Open the file in append and write-only mode. If the file doesn't exist, it won't be created.
    file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        // Print an error message if the file can't be opened and return the error.
        fmt.Println("Failed to open the file.")
        fmt.Printf("Failed to open the file: %v\n", err)
        return err
    }
    
    // Ensure the file is closed when the function exits, preventing resource leaks.
    defer file.Close()

    // Get the current timestamp for the note.
    timestamp := getUserTimestamp()

    // Create a new reader to read input from the user.
    reader := bufio.NewReader(os.Stdin)
    
    // Prompt the user to enter the title of the note.
    fmt.Print("\n\033[33mEnter the title of the note:\033[0m ")
    title, err := reader.ReadString('\n')
    if err != nil {
        // Print an error message if there's an issue with the input and return the error.
        fmt.Println("Invalid input.")
        return fmt.Errorf("invalid input: %v", err)
    }
    // Remove any leading or trailing whitespace from the title.
    title = strings.TrimSpace(title)

    // Prompt the user to enter tags for the note, separated by commas.
    fmt.Print("\n\033[33mEnter the tags for the note (comma-separated):\033[0m ")
    tags, err := reader.ReadString('\n')
    if err != nil {
        // Print an error message if there's an issue with the input and return the error.
        fmt.Println("Invalid input.")
        return fmt.Errorf("invalid input: %v", err)
    }
    // Remove any leading or trailing whitespace from the tags.
    tags = strings.TrimSpace(tags)
    
    // Prompt the user to enter the note content.
    fmt.Print("\n\033[33mEnter the note:\033[0m\n")
    input, err := reader.ReadString('\n')
    if err != nil {
        // Print an error message if there's an issue with the input and return the error.
        fmt.Println("Invalid input.")
        return fmt.Errorf("invalid input: %v", err)
    }
    // Remove any leading or trailing whitespace from the note content.
    input = strings.TrimSpace(input)
    
    // Check if the note content is empty and notify the user if it is.
    if input == "" {
        ClearTerminal()
        fmt.Println("\033[31mError: The note cannot be empty.\033[0m")
        return fmt.Errorf("empty note provided")
    }

    // Format the note with its metadata (timestamp, title, tags, and content).
    noteWithMetadata := fmt.Sprintf("[%s]\nTitle: %s\nTags: %s\nNote: %s\n", timestamp, title, tags, input)
    // Encrypt the note content using the ROT13 cipher.
    encryptedNote := rot13(noteWithMetadata)

    // Write the encrypted note to the file.
    _, err = file.WriteString(encryptedNote + "\n")
    if err != nil {
        // Print an error message if there's an issue with writing to the file and return the error.
        fmt.Println("Failed to write to file.")
        return fmt.Errorf("failed to write to file: %v", err)
    }

    // Clear the terminal and notify the user that the note was successfully added.
    ClearTerminal()
    fmt.Println("Text successfully added to the file!")
    return nil
}