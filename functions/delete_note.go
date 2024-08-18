package notes

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func DeleteANote(fileName string) error {
	// file opening for reading
	file, err := os.Open(fileName) // opens the file
	if err != nil {                // checks if an error occurred while opening the file
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close() //Ensures that the file is closed when the function exits, preventing resource leaks

	// we can read all notes from file
	var notes []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		notes = append(notes, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("failed to read file: %v", err)
	}

	// if 0 notes = no notes to delete
	if len(notes) == 0 {
		fmt.Println("No notes to delete.")
		return nil
	}

	// showing all notes with numbers
	fmt.Println("Notes:")
	for i, note := range notes {
		fmt.Printf("%d: %s\n", i+1, note) //printing each with its corresponding number.
	}

	// waiting for the note number input
	fmt.Print("Enter the number of the note to delete: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("invalid input: %v", err)
	}
	input = strings.TrimSpace(input)  //Removes any leading or trailing whitespace from the input.
	index, err := strconv.Atoi(input) //Converts the string input to an integer.
	if err != nil || index < 1 || index > len(notes) {
		return fmt.Errorf("invalid note number")
	}

	// deleting the note
	notes = append(notes[:index-1], notes[index:]...)

	// doing like resave for file, delete old and add new one without note
	file, err = os.Create(fileName) //Opens the file for writing
	if err != nil {
		return fmt.Errorf("failed to open file for writing: %v", err)
	}
	defer file.Close()

	for _, note := range notes { //loop through the remaining notes.
		_, err = file.WriteString(note + "\n") // re add every note
		if err != nil {
			return fmt.Errorf("failed to write to file: %v", err)
		}
	}

	fmt.Println("Note successfully deleted!")
	return nil
}
