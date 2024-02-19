package name

import "errors"

type Name string

func (n Name) String() string {
	return string(n)
}

func New(name string) (*Name, error) {
	if len([]rune(name)) > 100 {
		return nil, errors.New("name must be less or equal to 100 characters")
	}
	n := Name(name)
	return &n, nil
}
