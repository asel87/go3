package patronymic

import "errors"

type Patronymic string

func (p Patronymic) String() string {
	return string(p)
}

func New(patronymic string) (*Patronymic, error) {
	if len([]rune(patronymic)) > 100 {
		return nil, errors.New("patronymic must be less or equal to 100 characters")
	}
	p := Patronymic(patronymic)
	return &p, nil
}
