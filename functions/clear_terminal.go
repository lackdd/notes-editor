package notes

import (
	"fmt"
)

func ClearTerminal() {
	// Use ANSI escape sequences to clear the terminal for Unix-based systems
	// and a specific escape sequence for Windows
	fmt.Print("\033[H\033[2J\033[3J") // Unix-based (Linux, macOS)
	fmt.Print("\033c")                // Windows
}