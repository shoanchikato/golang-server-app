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
	Add(userId int, role *m.Author) error
	AddAll(userId int, roles *[]*m.Author) error
	Edit(userId int, id int, newAuthor *m.Author) error
	GetAll(userId int) (*[]m.Author, error)
	GetOne(userId int, id int) (*m.Author, error)
	Remove(userId int, id int) error
}

type roleAuthorization struct {
	auth      s.AuthorizationService
	validator v.AuthorValidator
}

func NewRole(auth s.AuthorizationService, validator v.AuthorValidator) Role {
	return &roleAuthorization{auth, validator}
}

// Add
func (r *roleAuthorization) Add(userId int, role *m.Author) error {
	err := r.auth.CheckForAuthorization(userId, p.RoleAdd.Name)
	if err != nil {
		return errors.Join(e.ErrAuthorDomain, e.ErrOnAdd, err)
	}

	return r.validator.Add(role)
}

// AddAll
func (r *roleAuthorization) AddAll(userId int, roles *[]*m.Author) error {
	err := r.auth.CheckForAuthorization(userId, p.RoleAddAll.Name)
	if err != nil {
		return errors.Join(e.ErrAuthorDomain, e.ErrOnAddAll, err)
	}

	return r.validator.AddAll(roles)
}

// Edit
func (r *roleAuthorization) Edit(userId int, id int, newAuthor *m.Author) error {
	err := r.auth.CheckForAuthorization(userId, p.RoleEdit.Name)
	if err != nil {
		return errors.Join(e.ErrAuthorDomain, e.ErrOnEdit, err)
	}

	return r.validator.Edit(id, newAuthor)
}

// GetAll
func (r *roleAuthorization) GetAll(userId int) (*[]m.Author, error) {
	err := r.auth.CheckForAuthorization(userId, p.RoleGetAll.Name)
	if err != nil {
		return nil, errors.Join(e.ErrAuthorDomain, e.ErrOnGetAll, err)
	}

	return r.validator.GetAll()
}

// GetOne
func (r *roleAuthorization) GetOne(userId int, id int) (*m.Author, error) {
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
