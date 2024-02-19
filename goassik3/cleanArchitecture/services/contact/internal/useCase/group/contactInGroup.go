package group

import (
	"assik3/services/contact/internal/domain/contact"
)

func (uc *UseCase) CreateContantIntoGroup(groupID int, contact *contact.Contact) (*contact.Contact, error) {
	createdContact, err := uc.adapterStorage.CreateContantIntoGroup(groupID, contact)
	if err != nil {
		return nil, err
	}
	return createdContact, nil
}

func (uc *UseCase) AddContactToGroup(groupID int, contactID int) error {
	err := uc.adapterStorage.AddContactToGroup(groupID, contactID)
	if err != nil {
		return err
	}
	return nil
}

func (uc *UseCase) DeleteContantFromGroup(groupID int, contactID int) error {
	err := uc.adapterStorage.DeleteContantFromGroup(groupID, contactID)
	if err != nil {
		return err
	}
	return nil
}
