package phonenumber

import (
    "fmt"
    "errors"
    "strings"
    "unicode"
)

var ErrInvalidPhoneNumber = errors.New("invalid phone number")

func Number(phoneNumber string) (string, error) {

	cleaned := strings.Map(func(r rune) rune {
		if unicode.IsDigit(r) {
			return r
		}
		return -1
	}, phoneNumber)
	if len(cleaned) == 11 && cleaned[0] == '1' {
		cleaned = cleaned[1:]
	}
	if len(cleaned) != 10 {
		return "", ErrInvalidPhoneNumber
	}
    if cleaned[0] < '2' || cleaned[0] > '9' {

		return "", ErrInvalidPhoneNumber
	}
	if cleaned[3] < '2' || cleaned[3] > '9' {
		return "", ErrInvalidPhoneNumber
	}
    
	return cleaned, nil
}

func AreaCode(phoneNumber string) (string, error) {
	number, err := Number(phoneNumber)

    if err != nil {
        return "", ErrInvalidPhoneNumber
    }

    return number[:3], nil
}

func Format(phoneNumber string) (string, error) {
	number, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("(%s) %s-%s",
		number[:3],
		number[3:6],
		number[6:],
	), nil
}
