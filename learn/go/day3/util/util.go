package util

import "fmt"

// FormatDuration converts minutes into a human-readable hours/minutes string.
func FormatDuration(minutes int) string {
	hours := minutes / 60
	mins := minutes % 60
	return fmt.Sprintf("%dh %dm", hours, mins)
}

// ParseScore converts a string score to an integer.
func ParseScore(input string) (int, error) {
	var score int
	_, err := fmt.Sscanf(input, "%d", &score)
	if err != nil {
		return 0, fmt.Errorf("invalid score %q: %w", input, err)
	}
	return score, nil
}

// Greet returns a greeting for the given name.
func Greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}
