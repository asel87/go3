package storage

import (
	"assik3/services/contact/internal/domain/contact"
	"assik3/services/contact/internal/domain/group"
)

type Storage interface {
	Contact
	Group
}

type Contact interface {
	CreateContact(contact *contact.Contact) (*contact.Contact, error)
	ReadContactByID(contactID int) (*contact.Contact, error)
	UpdateContact(contactID int, contact *contact.Contact) (*contact.Contact, error)
	DeleteContact(contactID int) error
}

type Group interface {
	CreateGroup(group *group.Group) (*group.Group, error)
	ReadGroupByID(groupID int) (*group.Group, error)
	UpdateGroup(groupID int, group *group.Group) (*group.Group, error)
	DeleteGroup(groupID int) error

	ContactGroup
}

type ContactGroup interface {
	CreateContantIntoGroup(groupID int, contact *contact.Contact) (*contact.Contact, error)
	AddContactToGroup(groupID int, contactID int) error
	DeleteContantFromGroup(groupID int, contactID int) error
}
