package postgres

import (
	"assik3/services/contact/internal/domain/contact"
	"database/sql"
	"errors"
	"fmt"
)

func (r *Repository) CreateContact(contact *contact.Contact) (*contact.Contact, error) {
	stmt := "INSERT INTO contacts (phoneNumber, name, surname, patronymic) VALUES(?, ?, ?, ?)"

	_, err := r.db.Exec(stmt, contact.PhoneNumber(), contact.Name(), contact.Surname(), contact.Patronymic())
	if err != nil {
		return nil, err
	}

	return contact, nil
}

func (r *Repository) ReadContactByID(contactID int) (*contact.Contact, error) {
	stmt := "SELECT id, phoneNumber, name, surname, patronymic FROM contacts WHERE id = ?"

	row := r.db.QueryRow(stmt, contactID)

	var c contact.Contact

	err := row.Scan(c.ID(), c.PhoneNumber(), c.Name(), c.Surname(), c.Patronymic())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &contact.Contact{}, fmt.Errorf("no rows found")
		} else {
			return &contact.Contact{}, err
		}
	}

	return &c, nil
}

func (r *Repository) UpdateContact(contactID int, contact *contact.Contact) (*contact.Contact, error) {
	stmt := `UPDATE contacts SET phoneNumber = ?, name = ?, surname = ?, patronymic = ? WHERE id = ?`

	_, err := r.db.Exec(stmt, contact.PhoneNumber(), contact.Name(), contact.Surname(), contact.Patronymic(), contactID)
	if err != nil {
		return nil, err
	}

	return contact, nil
}

func (r *Repository) DeleteContact(contactID int) error {
	stmt := "DELETE FROM contacts WHERE id = ?"

	_, err := r.db.Exec(stmt, contactID)
	if err != nil {
		return err
	}

	return nil
}
