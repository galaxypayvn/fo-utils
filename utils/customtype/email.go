package customtype

import "regexp"

// Email is a custom type that extends string
type Email string

// Validate checks if the email address has a valid format using regex
func (e Email) Validate() bool {
	// Regular expression for validating an email
	var emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(string(e))
}

// String returns the string representation of the Email
func (e Email) String() string {
	return string(e)
}

// Domain extracts the domain part of the email address
func (e Email) Domain() string {
	// Split the email by '@' and return the part after it if possible
	parts := regexp.MustCompile(`@`).Split(string(e), 2)
	if len(parts) == 2 {
		return parts[1]
	}
	return ""
}
