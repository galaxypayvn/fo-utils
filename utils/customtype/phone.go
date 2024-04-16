package customtype

import "regexp"

const (
	VietNamPhoneFormat = `^\+(8480|8495|8486|8487|8496|8497|8498|8432|8433|8434|8435|8436|8437|8438|8439|8488|8491|8494|8483|8484|8485|8481|8482|8489|8490|8493|8470|8479|8477|8476|8478|8492|8456|8458|8499|8459|8455|8452|8415)+([0-9]{7})\b$`
)

// PhoneFormatType Define Phone format type using iota
type PhoneFormatType int

const (
	International  PhoneFormatType = iota // +84xxxxxxxxx
	PrefixWithZero                        // 0xxxxxxxxx
)

// Phone is a custom type that extends string
type Phone string

// Validate checks if the phone number is in a valid Vietnamese format using regex
func (p Phone) Validate(locale string) bool {
	phoneFormat := map[string]string{}
	phoneFormat["VN"] = VietNamPhoneFormat

	// Regular expression for validating a phone number
	var phoneRegex = regexp.MustCompile(locale)
	return phoneRegex.MatchString(string(p))
}

// String returns the string representation of the Phone
func (p Phone) String() string {
	return string(p)
}

// Format ensures the phone number is in international format
// 0 : International format
// 1 : PrefixWithZero 0xxxxxxxxx format (For VN)
func (p Phone) Format(formatType PhoneFormatType, locale string) string {
	if locale == "" {
		locale = "VN"
	}
	if !p.Validate(locale) {
		return "Invalid format"
	}

	switch formatType {
	case International:
		return p.String()
	case PrefixWithZero:
		return "0" + p.String()[3:]
	default:
		return p.String()
	}
}
