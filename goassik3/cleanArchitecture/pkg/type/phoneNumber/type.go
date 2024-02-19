package phoneNumber

import (
	"errors"
	"regexp"
	"strings"
)

type PhoneNumber string

func (p PhoneNumber) String() string {
	return string(p)
}

func New(phone string) (*PhoneNumber, error) {
	re := regexp.MustCompile(`^\d+$`)

	if !re.MatchString(phone) {
		return nil, errors.New("phone number contains invalid characters, only numbers are allowed")
	}

	p := PhoneNumber(phone)
	return &p, nil
}

// func (p PhoneNumber) Equal(phoneNumber PhoneNumber) bool {
// 	return p == phoneNumber
// }

func (p PhoneNumber) IsEmpty() bool {
	return len(strings.TrimSpace(string(p))) == 0
}
