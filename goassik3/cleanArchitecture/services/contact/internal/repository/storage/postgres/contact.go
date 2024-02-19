package postgres

import "assik3/services/contact/internal/domain/contact"

func (r *Repository) CreateContact(contact *contact.Contact) (*contact.Contact, error) {
	panic("implement later")
}

func (r *Repository) ReadContactByID(contactID int) (*contact.Contact, error) {
	panic("implement later")
}

func (r *Repository) UpdateContact(contactID int, contact *contact.Contact) (*contact.Contact, error) {
	panic("implement later")
}

func (r *Repository) DeleteContact(contactID int) error {
	panic("implement later")
}
