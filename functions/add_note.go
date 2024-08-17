package notes

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func AddANote(fileName string) error {
	fmt.Print("Enter the note: ")
    file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
        return fmt.Errorf("failed to open file: %v", err)
    }
    
    defer file.Close()
	
	reader := bufio.NewReader(os.Stdin)
    input, err := reader.ReadString('\n')
    if err != nil {
        return fmt.Errorf("invalid input: %v", err)
    }
    input = strings.TrimSpace(input)
    _, err = file.WriteString(input + "\n")
	if err != nil {
        return fmt.Errorf("failed to write to file: %v", err)
    }
   
    fmt.Println("Text successfully added to the file!")
    return nil
}