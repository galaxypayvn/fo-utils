package utmail

import (
	"regexp"
)

// ValidEmail is a function that validates if the input string is a valid email address.
func ValidEmail(input string) bool {
	pattern := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(pattern)
	if !re.MatchString(input) {
		return false
	}
	return true
}
