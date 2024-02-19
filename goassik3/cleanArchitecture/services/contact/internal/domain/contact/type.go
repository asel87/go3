package contact

import (
	"errors"

	"assik3/pkg/type/phoneNumber"
	"assik3/services/contact/internal/domain/contact/name"
	"assik3/services/contact/internal/domain/contact/patronymic"
	"assik3/services/contact/internal/domain/contact/surname"
)

type Contact struct {
	id          int
	phoneNumber phoneNumber.PhoneNumber
	name        name.Name
	surname     surname.Surname
	patronymic  patronymic.Patronymic
}

func New(
	id int,
	phoneNumber phoneNumber.PhoneNumber,
	name name.Name,
	surname surname.Surname,
	patronymic patronymic.Patronymic) (*Contact, error) {

	if phoneNumber.IsEmpty() {
		return nil, errors.New("phone number is required")
	}

	return &Contact{
		id:          id,
		phoneNumber: phoneNumber,
		name:        name,
		surname:     surname,
		patronymic:  patronymic,
	}, nil
}

func (c Contact) ID() int {
	return c.id
}

func (c Contact) PhoneNumber() phoneNumber.PhoneNumber {
	return c.phoneNumber
}

func (c Contact) Name() name.Name {
	return c.name
}

func (c Contact) Surname() surname.Surname {
	return c.surname
}

func (c Contact) Patronymic() patronymic.Patronymic {
	return c.patronymic
}
