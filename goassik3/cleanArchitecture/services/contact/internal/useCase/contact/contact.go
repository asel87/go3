package contact

import "assik3/services/contact/internal/domain/contact"

func (uc *UseCase) Create(contact *contact.Contact) (*contact.Contact, error) {
	createdContact, err := uc.adapterStorage.CreateContact(contact)
	if err != nil {
		return nil, err
	}
	return createdContact, nil
}

func (uc *UseCase) ReadByID(contactID int) (*contact.Contact, error) {
	foundContact, err := uc.adapterStorage.ReadContactByID(contactID)
	if err != nil {
		return nil, err
	}
	return foundContact, nil
}

func (uc *UseCase) Update(contactID int, contact *contact.Contact) (*contact.Contact, error) {
	updatedContact, err := uc.adapterStorage.UpdateContact(contactID, contact)
	if err != nil {
		return nil, err
	}
	return updatedContact, nil
}

func (uc *UseCase) Delete(contactID int) error {
	err := uc.adapterStorage.DeleteContact(contactID)
	if err != nil {
		return err
	}
	return nil
}
