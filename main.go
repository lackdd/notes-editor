package main

import (
	"fmt"
	"notes/functions"
)

func main() {
	fmt.Println("hello")
	fmt.Println(notes.Test())
}
package sprint

package sprint

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func AddANote(filename string) error {
    file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
    
    defer file.Close()
	
    fmt.Print("Enter the note: ")
	reader := bufio.NewReader(os.Stdin)
    input, err := reader.ReadString('\n')
    if err != nil {
        return fmt.Errorf("Invalid input: %v", err)
    }
    input = strings.TrimSpace(input)
    _, err = file.WriteString(input + "\n")
   
    fmt.Println("Text successfully added to the file!")
    return nil
}