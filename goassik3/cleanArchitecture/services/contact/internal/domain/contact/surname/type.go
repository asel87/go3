package surname

import "errors"

type Surname string

func (s Surname) String() string {
	return string(s)
}

func New(surname string) (*Surname, error) {
	if len([]rune(surname)) > 100 {
		return nil, errors.New("surname must be less or equal to 100 characters")
	}
	s := Surname(surname)
	return &s, nil
}
