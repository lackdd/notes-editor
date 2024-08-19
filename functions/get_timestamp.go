package notes
import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func getUserTimestamp() string {
	// Create a new reader to read input from the user.
	reader := bufio.NewReader(os.Stdin)

	// Prompt the user to enter the year.
	fmt.Print("Enter the year (e.g., 2024): ")
	year, _ := reader.ReadString('\n')
	// Remove any leading or trailing whitespace from the input (e.g., newline characters).
	year = strings.TrimSpace(year)

	// Prompt the user to enter the month.
	fmt.Print("Enter the month (01-12): ")
	month, _ := reader.ReadString('\n')
	// Remove any leading or trailing whitespace from the input.
	month = strings.TrimSpace(month)

	// Prompt the user to enter the day.
	fmt.Print("Enter the day (01-31): ")
	day, _ := reader.ReadString('\n')
	// Remove any leading or trailing whitespace from the input.
	day = strings.TrimSpace(day)

	// Return the formatted timestamp as "YYYY-MM-DD".
	return fmt.Sprintf("%s-%s-%s", year, month, day)
}