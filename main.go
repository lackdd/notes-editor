package main

import (
	"fmt"
	notes "notes/functions"
	"os"
)

func main() {
	if len(os.Args) != 2 || os.Args[1] == "help" {
		fmt.Printf(`Hello, to use this tool please add exactly one argument as a textfile's name to the commandline as in './notestool (your_textfile_name)'
You will be given a menu, where you can choose, which action to take-
If you choose 1. Show notes, the program will show all the notes in the specified collection.
If you choose 2. Add a note, the program will let you add a new note to the collection and if the collection doesn't exist, also create it.
If you choose 3. Delete a note, the program will delete a note in the collection with your specified line.
Choosing 4. Exit will close the program.`)
		return
	} else {
		fmt.Printf("\nHello and welcome to the notes tool.\n")
		for {
			filename := fmt.Sprintf("%s.txt" ,os.Args[1])
			answer := notes.GetInput()
			if answer == "1" {
				notes.ShowNotes(filename)
			} else if answer == "2" {
				notes.AddANote(filename)
			} else if answer == "3" {
				notes.DeleteANote(filename)
			} else if answer == "4" {
				os.Exit(0)
			}
		}
	}
}
