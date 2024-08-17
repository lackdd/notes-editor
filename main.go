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

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func AddANote(string) error {
    file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
   
    defer file.Close()
	
	reader := bufio.NewReader(os.Stdin)
    input, err := reader.ReadString('\n')
   
    input = strings.TrimSpace(input)
    _, err = file.WriteString(input + "\n")
   
    fmt.Println("Text successfully added to the file!")
    return nil
}