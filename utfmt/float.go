package utfmt

import (
	"fmt"
	"strings"
)

// FormatWithThousandsSeparator formats a float64 amount by inserting a separator every three digits.
//
// Parameters:
//
// - amount: The float64 amount to be formatted.
//
// - separate: The rune to be used as the separator. If zero, a comma is used.
//
// Returns:
//
// - A string representing the formatted amount with separators.
func FormatWithThousandsSeparator(value float64, separate rune) string {
	if value < 0 {
		return "-" + FormatWithThousandsSeparator(-value, separate)
	}
	if separate == 0 {
		separate = ','
	}

	var (
		result   strings.Builder
		valueStr = fmt.Sprintf("%d", int(value))
		length   = len(valueStr)
	)
	for i, digit := range valueStr {
		if i > 0 && (length-i)%3 == 0 {
			result.WriteRune(separate)
		}
		result.WriteRune(digit)
	}
	return result.String()
}
