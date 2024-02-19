package group

import (
	"errors"

	"assik3/services/contact/internal/domain/group/name"
)

type Group struct {
	id   int
	name name.Name
}

func New(id int, name name.Name) (*Group, error) {

	if name.IsEmpty() {
		return nil, errors.New("group name is required")
	}

	return &Group{
		id:   id,
		name: name,
	}, nil
}

func (g Group) ID() int {
	return g.id
}

func (g Group) Name() name.Name {
	return g.name
}
