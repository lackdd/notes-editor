package notes

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func AddANote(fileName string) error {
    file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
    if err != nil {
        fmt.Printf("Failed to open the file: %v\n", err)
        return err
    }
    defer file.Close()

    timestamp := getUserTimestamp()
    reader := bufio.NewReader(os.Stdin)
    
    fmt.Print("\n\033[33mEnter the title of the note:\033[0m ")
    title, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Invalid input.")
        return fmt.Errorf("invalid input: %v", err)
    }
    title = strings.TrimSpace(title)

    fmt.Print("\n\033[33mEnter the tags for the note (comma-separated):\033[0m ")
    tags, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Invalid input.")
        return fmt.Errorf("invalid input: %v", err)
    }
    tags = strings.TrimSpace(tags)
    
    fmt.Print("\n\033[33mEnter the note:\033[0m\n")
    input, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Invalid input.")
        return fmt.Errorf("invalid input: %v", err)
    }
    input = strings.TrimSpace(input)
    
    if input == "" {
        fmt.Println("\033[31mError: The note cannot be empty.\033[0m")
        return fmt.Errorf("empty note provided")
    }

    // Format the note as a single line with separators.
    note := fmt.Sprintf("%s|%s|%s|%s", timestamp, title, tags, input)

    // Encrypt the note content using ROT13.
    encryptedNote := rot13(note)

    // Write the encrypted note to the file.
    _, err = file.WriteString(encryptedNote + "\n")
    if err != nil {
        fmt.Println("Failed to write to file.")
        return fmt.Errorf("failed to write to file: %v", err)
    }

    fmt.Println("Text successfully added to the file!")
    return nil
}