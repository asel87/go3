package name

import (
	"errors"
	"strings"
)

type Name string

func (n Name) String() string {
	return string(n)
}

func New(name string) (*Name, error) {
	if len([]rune(name)) > 250 {
		return nil, errors.New("group name must be less or equal to 250 characters")
	}
	n := Name(name)
	return &n, nil
}

func (n Name) IsEmpty() bool {
	return len(strings.TrimSpace(string(n))) == 0
}
