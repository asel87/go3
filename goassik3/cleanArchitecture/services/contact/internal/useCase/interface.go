package useCase

import (
	"assik3/services/contact/internal/domain/contact"
	"assik3/services/contact/internal/domain/group"
)

type Contact interface {
	Create(contact *contact.Contact) (*contact.Contact, error)
	ReadByID(contactID int) (*contact.Contact, error)
	Update(contactID int, contact *contact.Contact) (*contact.Contact, error)
	Delete(contactID int) error
}

type Group interface {
	Create(group *group.Group) (*group.Group, error)
	ReadByID(groupID int) (*group.Group, error)
	Update(groupID int, group *group.Group) (*group.Group, error)
	Delete(groupID int) error

	ContactGroup
}

type ContactGroup interface {
	CreateContantIntoGroup(groupID int, contact *contact.Contact) (*contact.Contact, error)
	AddContactToGroup(groupID int, contactID int) error
	DeleteContantFromGroup(groupID int, contactID int) error
}
