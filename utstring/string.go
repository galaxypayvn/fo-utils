package utstring

import (
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"strings"
	"unicode"
)

// isMark returns true if the rune is a non-spacing mark
func isMark(r rune) bool {
	return unicode.Is(unicode.Mn, r)
}

func TransformString(input string, uppercase bool) string {
	input = strings.TrimSpace(input)

	// Remove invalid UTF-8 sequences
	input = strings.Map(func(r rune) rune {
		if r == '\uFFFD' {
			return -1 // Remove the Unicode replacement character
		}
		return r
	}, input)

	// Create a transformation pipeline
	t := transform.Chain(
		norm.NFD,                     // Canonical decomposition
		transform.RemoveFunc(isMark), // Remove non-spacing marks
		norm.NFC,                     // Canonical composition
	)

	// Apply the transformation pipeline
	result, _, err := transform.String(t, input)
	if err != nil {
		return ""
	}

	// Replace specific characters
	result = strings.ReplaceAll(result, "Đ", "D")
	result = strings.ReplaceAll(result, "đ", "d")

	// Convert case
	if uppercase {
		return strings.ToUpper(result)
	}
	return strings.ToLower(result)
}

// RemoveDuplicates removes duplicate elements from a slice of strings.
func RemoveDuplicatesInArrayString(input []string) []string {
	uniqueMap := make(map[string]struct{})
	for _, item := range input {
		uniqueMap[item] = struct{}{}
	}

	uniqueSlice := make([]string, 0, len(uniqueMap))
	for key := range uniqueMap {
		uniqueSlice = append(uniqueSlice, key)
	}

	return uniqueSlice
}
