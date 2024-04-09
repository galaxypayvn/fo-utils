package utphone

import (
	"regexp"
	"strings"
)

// ValidateVietnamesePhoneNumber is a function that validates a Vietnamese phone number.
// It takes a string as an input, which should be the phone number to validate.
// The function uses a regular expression to check if the phone number matches the Vietnamese phone number format.
// It returns true if the phone number is valid and false otherwise.
func ValidateVietnamesePhoneNumber(phoneNumber string) bool {
	VietnamesePhoneNumberRegex := `^(\+84|0)(3[2-9]|5[689]|7[06-9]|8[1-9]|9[0-4]|9[6-9])[0-9]{7}$`

	re := regexp.MustCompile(VietnamesePhoneNumberRegex)
	if !re.MatchString(phoneNumber) {
		return false
	}
	return true
}

func ConvertVNPhoneFormat(phone string) string {
	if phone != "" {
		if strings.HasPrefix(phone, "84") {
			phone = "+" + phone
		}
		if strings.HasPrefix(phone, "0") {
			phone = "+84" + phone[1:]
		}
	}
	return phone
}
