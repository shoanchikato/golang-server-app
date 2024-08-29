package authorization

import (
	e "app/pkg/errors"
	m "app/pkg/model"
	p "app/pkg/permission"
	s "app/pkg/service"
	v "app/pkg/validation"
	"errors"
)

type Role interface {
	Add(userId int, role *m.Role) error
	AddAll(userId int, roles *[]*m.Role) error
	Edit(userId int, id int, newAuthor *m.Role) error
	GetAll(userId, lastId, limit int) (*[]m.Role, error)
	GetOne(userId int, id int) (*m.Role, error)
	Remove(userId int, id int) error
}

type roleAuthorization struct {
	auth      s.AuthorizationService
	validator v.RoleValidator
}

func NewRoleAuthorization(auth s.AuthorizationService, validator v.RoleValidator) Role {
	return &roleAuthorization{auth, validator}
}

// Add
func (r *roleAuthorization) Add(userId int, role *m.Role) error {
	err := r.auth.CheckForAuthorization(userId, p.RoleAdd.Name)
	if err != nil {
		return errors.Join(e.ErrAuthorDomain, e.ErrOnAdd, err)
	}

	return r.validator.Add(role)
}

// AddAll
func (r *roleAuthorization) AddAll(userId int, roles *[]*m.Role) error {
	err := r.auth.CheckForAuthorization(userId, p.RoleAddAll.Name)
	if err != nil {
		return errors.Join(e.ErrAuthorDomain, e.ErrOnAddAll, err)
	}

	return r.validator.AddAll(roles)
}

// Edit
func (r *roleAuthorization) Edit(userId int, id int, newAuthor *m.Role) error {
	err := r.auth.CheckForAuthorization(userId, p.RoleEdit.Name)
	if err != nil {
		return errors.Join(e.ErrAuthorDomain, e.ErrOnEdit, err)
	}

	return r.validator.Edit(id, newAuthor)
}

// GetAll
func (r *roleAuthorization) GetAll(userId, lastId, limit int) (*[]m.Role, error) {
	err := r.auth.CheckForAuthorization(userId, p.RoleGetAll.Name)
	if err != nil {
		return nil, errors.Join(e.ErrAuthorDomain, e.ErrOnGetAll, err)
	}

	return r.validator.GetAll(lastId, limit)
}

// GetOne
func (r *roleAuthorization) GetOne(userId int, id int) (*m.Role, error) {
	err := r.auth.CheckForAuthorization(userId, p.RoleGetOne.Name)
	if err != nil {
		return nil, errors.Join(e.ErrAuthorDomain, e.ErrOnGetOne, err)
	}

	return r.validator.GetOne(id)
}

// Remove
func (r *roleAuthorization) Remove(userId int, id int) error {
	err := r.auth.CheckForAuthorization(userId, p.RoleRemove.Name)
	if err != nil {
		return errors.Join(e.ErrAuthorDomain, e.ErrOnRemove, err)
	}

	return r.validator.Remove(id)
}
