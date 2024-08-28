package authorization

import (
	e "app/pkg/errors"
	m "app/pkg/model"
	pe "app/pkg/permission"
	s "app/pkg/service"
	v "app/pkg/validation"
	"errors"
)

type PermissionAuthorization interface {
	Add(userId int, permission *m.Permission) error
	AddAll(userId int, permissions *[]*m.Permission) error
	Edit(userId int, id int, newPermission *m.Permission) error
	GetAll(userId int) (*[]m.Permission, error)
	GetOne(userId int, id int) (*m.Permission, error)
	Remove(userId int, id int) error
}

type permissionAuthorization struct {
	auth      s.AuthorizationService
	validator v.PermissionValidator
}

func NewPermissionAuthorization(
	auth s.AuthorizationService,
	validator v.PermissionValidator,
) PermissionAuthorization {
	return &permissionAuthorization{auth, validator}
}

// Add
func (p *permissionAuthorization) Add(userId int, permission *m.Permission) error {
	err := p.auth.CheckForAuthorization(userId, pe.PermissionAdd.Name)
	if err != nil {
		return errors.Join(e.ErrBookDomain, e.ErrOnAdd, err)
	}

	return p.validator.Add(permission)
}

// AddAll
func (p *permissionAuthorization) AddAll(userId int, permissions *[]*m.Permission) error {
	err := p.auth.CheckForAuthorization(userId, pe.PermissionAdd.Name)
	if err != nil {
		return errors.Join(e.ErrPermissionDomain, e.ErrOnAdd, err)
	}

	return p.validator.AddAll(permissions)
}

// Edit
func (p *permissionAuthorization) Edit(userId int, id int, newPermission *m.Permission) error {
	err := p.auth.CheckForAuthorization(userId, pe.PermissionAdd.Name)
	if err != nil {
		return errors.Join(e.ErrPermissionDomain, e.ErrOnAdd, err)
	}

	return p.validator.Edit(id, newPermission)
}

// GetAll
func (p *permissionAuthorization) GetAll(userId int) (*[]m.Permission, error) {
	err := p.auth.CheckForAuthorization(userId, pe.PermissionAdd.Name)
	if err != nil {
		return nil, errors.Join(e.ErrPermissionDomain, e.ErrOnAdd, err)
	}

	return p.validator.GetAll()
}

// GetOne
func (p *permissionAuthorization) GetOne(userId int, id int) (*m.Permission, error) {
	err := p.auth.CheckForAuthorization(userId, pe.PermissionAdd.Name)
	if err != nil {
		return nil, errors.Join(e.ErrPermissionDomain, e.ErrOnAdd, err)
	}

	return p.validator.GetOne(id)
}

// Remove
func (p *permissionAuthorization) Remove(userId int, id int) error {
	err := p.auth.CheckForAuthorization(userId, pe.PermissionAdd.Name)
	if err != nil {
		return errors.Join(e.ErrPermissionDomain, e.ErrOnAdd, err)
	}

	return p.validator.Remove(id)
}
