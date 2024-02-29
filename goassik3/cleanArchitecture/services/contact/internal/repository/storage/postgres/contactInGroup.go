package postgres

import (
	"assik3/services/contact/internal/domain/contact"
)

func (r *Repository) CreateContantIntoGroup(groupID int, contact *contact.Contact) (*contact.Contact, error) {
	stmt := "INSERT INTO contacts (phoneNumber, name, surname, patronymic) VALUES(?, ?, ?, ?)"

	_, err := r.db.Exec(stmt, contact.PhoneNumber(), contact.Name(), contact.Surname(), contact.Patronymic())
	if err != nil {
		return nil, err
	}

	stmt = "INSERT INTO contactGroup (contactId, groupId) VALUES(?, ?)"

	_, err = r.db.Exec(stmt, contact.ID(), groupID)
	if err != nil {
		return nil, err
	}

	return contact, nil
}

func (r *Repository) AddContactToGroup(groupID int, contactID int) error {
	stmt := "INSERT INTO contactGroup (contactId, groupId) VALUES(?, ?)"

	_, err := r.db.Exec(stmt, contactID, groupID)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeleteContantFromGroup(groupID int, contactID int) error {
	stmt := "DELETE FROM contactGroup WHERE contactId = ? AND groupId = ?"

	_, err := r.db.Exec(stmt, contactID, groupID)
	if err != nil {
		return err
	}

	return nil
}
