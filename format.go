package formatter

import "strings"

const (
	withGreen string = "\033[32m"
	withRed   string = "\033[31m"
	reset     string = "\033[0m"
)

// Format prints the input string in a proper format
func Format(input string) string {
	formatted := []string{withRed, input, reset}
	return strings.Join(formatted, "")
}
