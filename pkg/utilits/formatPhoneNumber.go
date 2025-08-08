package utilits

import (
	"errors"
	"strings"
)

var (
	ErrInvalidPhoneLength = errors.New("invalid phone number length")
	ErrInvalidPhoneFormat = errors.New("invalid phone number format")
)

func FormatPhoneNumber(phone string) (string, error) {
	var cleaned strings.Builder
	for _, r := range phone {
		if r >= '0' && r <= '9' {
			cleaned.WriteRune(r)
		}
	}
	digits := cleaned.String()

	if len(digits) < 10 || len(digits) > 15 {
		return "", ErrInvalidPhoneLength
	}

	if strings.HasPrefix(digits, "8") && len(digits) == 11 {
		return "+7" + digits[1:], nil
	}

	if len(digits) == 10 {
		return "+7" + digits, nil
	}

	if len(digits) >= 11 {
		if digits[0] == '7' && len(digits) == 11 {
			return "+" + digits, nil
		}
		return "+" + digits, nil
	}

	return "", ErrInvalidPhoneFormat
}
