package group

import "assik3/services/contact/internal/domain/group"

func (uc *UseCase) Create(group *group.Group) (*group.Group, error) {
	createdGroup, err := uc.adapterStorage.CreateGroup(group)
	if err != nil {
		return nil, err
	}
	return createdGroup, nil
}

func (uc *UseCase) ReadByID(groupID int) (*group.Group, error) {
	foundGroup, err := uc.adapterStorage.ReadGroupByID(groupID)
	if err != nil {
		return nil, err
	}
	return foundGroup, nil
}

func (uc *UseCase) Update(groupID int, group *group.Group) (*group.Group, error) {
	updatedGroup, err := uc.adapterStorage.UpdateGroup(groupID, group)
	if err != nil {
		return nil, err
	}
	return updatedGroup, nil
}

func (uc *UseCase) Delete(groupID int) error {
	err := uc.adapterStorage.DeleteGroup(groupID)
	if err != nil {
		return err
	}
	return nil
}
