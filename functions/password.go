package notes
import (
	"bufio"
	"fmt"
	"strings"
	"os"
)

// passwordFile defines the filename where the hashed password will be stored.
const passwordFile = ".password"

// GetPassword prompts the user to enter a password and reads it from stdin.
func GetPassword(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	password, _ := reader.ReadString('\n') // Read the input until a newline is encountered.
	return strings.TrimSpace(password)     // Remove any leading/trailing whitespace.
}

// HashPassword hashes the password using ROT13.
// This is a simple example and not secure for real-world applications.
func HashPassword(password string) string {
	return rot13(password) // Apply the ROT13 cipher to the password.
}

// VerifyPassword checks if the input password matches the stored hash.
func VerifyPassword() bool {
	// Check if the password file exists.
	if _, err := os.Stat(passwordFile); os.IsNotExist(err) {
		// If the password file doesn't exist, prompt the user to set a new password.
		fmt.Println("No password set. Setting a new password.")
		setPassword := GetPassword("Enter a new password: ")
		hashedPassword := HashPassword(setPassword)               // Hash the new password.
		os.WriteFile(passwordFile, []byte(hashedPassword), 0644) // Save the hashed password to the file.
		return true
	}

	// If the password file exists, read the stored hash.
	storedHash, _ := os.ReadFile(passwordFile)
	// Prompt the user to enter their password.
	inputPassword := GetPassword("Enter your password: ")
	// Hash the input password to compare with the stored hash.
	hashedInput := HashPassword(inputPassword)

	// Compare the hashed input password with the stored hash.
	return hashedInput == string(storedHash)
}

// rot13 applies the ROT13 cipher to a string.
// ROT13 shifts each letter by 13 places in the alphabet.
// It is its own inverse, meaning applying it twice returns the original text.
func rot13(s string) string {
	result := make([]rune, len(s))
	for i, c := range s {
		switch {
		case 'a' <= c && c <= 'z': // If the character is a lowercase letter.
			result[i] = 'a' + (c-'a'+13)%26 // Shift by 13 within the lowercase letters.
		case 'A' <= c && c <= 'Z': // If the character is an uppercase letter.
			result[i] = 'A' + (c-'A'+13)%26 // Shift by 13 within the uppercase letters.
		default:
			result[i] = c // Non-alphabetic characters are not changed.
		}
	}
	return string(result) // Convert the result slice back to a string.
}