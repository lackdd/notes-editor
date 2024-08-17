package notes

import (
	"bufio"
	"fmt"
	"os"
)

func ShowNotes(filename string) {
	readFile, err := os.Open(filename)

	if err != nil {
		fmt.Println("Error opening file", err)
		return
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
	}
	if err := fileScanner.Err(); err != nil {
		fmt.Println("Error reading file", err)
	}
	readFile.Close()
	return
}