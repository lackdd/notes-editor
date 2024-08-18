package notes

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetInput() string {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Hello and welcome to the notes tool.")

	for {
		fmt.Println("Please choose an operation:\n1. Show notes.\n2. Add a note.\n3. Delete a note.\n4. Exit.")
		scanner.Scan()
		operation := strings.TrimSpace(scanner.Text())
		if operation == "1" || operation == "2" || operation == "3" || operation == "4" {
			return operation
		} else {
			fmt.Println("Invalid input: ", operation)
		}
	}
}