package internal

import "fmt"

// Underline returns the given text with ANSI escape codes to underline it in the terminal.
func Underline(text string) string {
	return fmt.Sprintf("\033[4m%s\033[0m", text)
}
