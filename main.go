package main

import (
	"fmt"
	notes "notes/functions"
	"os"
)

func main() {
	if len(os.Args) != 2 || os.Args[1] == "help" {
		fmt.Printf(`Hello, to use this tool please add exactly one argument as a collection name to the commandline as in './notestool (your_collection_name)'`)
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
