package permission

import (
	e "app/pkg/errors"
	"errors"
)

type PermissionAuthorization interface {
	Add(userID int, permission *Permission) error
	AddAll(userID int, permissions *[]*Permission) error
	Edit(userID int, id int, newPermission *Permission) error
	GetAll(userID int) (*[]Permission, error)
	GetOne(userID int, id int) (*Permission, error)
	Remove(userID int, id int) error
}

type permissionAuthorization struct {
	auth AuthorizationService
	v    PermissionValidator
}

func NewPermissionAuthorization(
	auth AuthorizationService,
	v PermissionValidator,
) PermissionAuthorization {
	return &permissionAuthorization{auth, v}
}

// Add
func (p *permissionAuthorization) Add(userID int, permission *Permission) error {
	err := p.auth.CheckForAuthorization(userID, PermissionAdd.Name)
	if err != nil {
		return errors.Join(e.ErrBookDomain, e.ErrOnAdd, err)
	}

	return p.v.Add(permission)
}

// AddAll
func (p *permissionAuthorization) AddAll(userID int, permissions *[]*Permission) error {
	err := p.auth.CheckForAuthorization(userID, PermissionAdd.Name)
	if err != nil {
		return errors.Join(e.ErrPermissionDomain, e.ErrOnAdd, err)
	}

	return p.v.AddAll(permissions)
}

// Edit
func (p *permissionAuthorization) Edit(userID int, id int, newPermission *Permission) error {
	err := p.auth.CheckForAuthorization(userID, PermissionAdd.Name)
	if err != nil {
		return errors.Join(e.ErrPermissionDomain, e.ErrOnAdd, err)
	}

	return p.v.Edit(id, newPermission)
}

// GetAll
func (p *permissionAuthorization) GetAll(userID int) (*[]Permission, error) {
	err := p.auth.CheckForAuthorization(userID, PermissionAdd.Name)
	if err != nil {
		return nil, errors.Join(e.ErrPermissionDomain, e.ErrOnAdd, err)
	}

	return p.v.GetAll()
}

// GetOne
func (p *permissionAuthorization) GetOne(userID int, id int) (*Permission, error) {
	err := p.auth.CheckForAuthorization(userID, PermissionAdd.Name)
	if err != nil {
		return nil, errors.Join(e.ErrPermissionDomain, e.ErrOnAdd, err)
	}

	return p.v.GetOne(id)
}

// Remove
func (p *permissionAuthorization) Remove(userID int, id int) error {
	err := p.auth.CheckForAuthorization(userID, PermissionAdd.Name)
	if err != nil {
		return errors.Join(e.ErrPermissionDomain, e.ErrOnAdd, err)
	}

	return p.v.Remove(id)
}
