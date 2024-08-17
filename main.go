package main

import (
	"fmt"
	"notes/functions"
	"os"
)

func main() {
	for {
		filename := fmt.Sprintf("%s.txt" ,os.Args[1])
		answer := notes.GetInput()
		if answer == "1" {
			notes.ShowNotes(filename)
		} else if answer == "2" {
			notes.AddANote(filename)
		} else if answer == "4" {
			os.Exit(0)
		}
	}
}


    /*err := notes.AddANote(filename)
    if err != nil {
        fmt.Println("Error:", err)
    }*/
	//notes.AddANote(filename)
	//notes.ShowNotes("test.txt")