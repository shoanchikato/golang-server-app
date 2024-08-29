package authorization

import (
	e "app/pkg/errors"
	m "app/pkg/model"
	p "app/pkg/permission"
	s "app/pkg/service"
	v "app/pkg/validation"
	"errors"
)

type User interface {
	Add(userId int, user *m.Author) error
	AddAll(userId int, users *[]*m.Author) error
	Edit(userId int, id int, newAuthor *m.Author) error
	GetAll(userId, lastId, limit int) (*[]m.Author, error)
	GetOne(userId int, id int) (*m.Author, error)
	Remove(userId int, id int) error
}

type userAuthorization struct {
	auth      s.AuthorizationService
	validator v.AuthorValidator
}

func NewUser(auth s.AuthorizationService, validator v.AuthorValidator) User {
	return &userAuthorization{auth, validator}
}

// Add
func (u *userAuthorization) Add(userId int, user *m.Author) error {
	err := u.auth.CheckForAuthorization(userId, p.UserAdd.Name)
	if err != nil {
		return errors.Join(e.ErrAuthorDomain, e.ErrOnAdd, err)
	}

	return u.validator.Add(user)
}

// AddAll
func (u *userAuthorization) AddAll(userId int, users *[]*m.Author) error {
	err := u.auth.CheckForAuthorization(userId, p.UserAddAll.Name)
	if err != nil {
		return errors.Join(e.ErrAuthorDomain, e.ErrOnAddAll, err)
	}

	return u.validator.AddAll(users)
}

// Edit
func (u *userAuthorization) Edit(userId int, id int, newAuthor *m.Author) error {
	err := u.auth.CheckForAuthorization(userId, p.UserEdit.Name)
	if err != nil {
		return errors.Join(e.ErrAuthorDomain, e.ErrOnEdit, err)
	}

	return u.validator.Edit(id, newAuthor)
}

// GetAll
func (u *userAuthorization) GetAll(userId, lastId, limit int) (*[]m.Author, error) {
	err := u.auth.CheckForAuthorization(userId, p.UserGetAll.Name)
	if err != nil {
		return nil, errors.Join(e.ErrAuthorDomain, e.ErrOnGetAll, err)
	}

	return u.validator.GetAll(lastId, limit)
}

// GetOne
func (u *userAuthorization) GetOne(userId int, id int) (*m.Author, error) {
	err := u.auth.CheckForAuthorization(userId, p.UserGetOne.Name)
	if err != nil {
		return nil, errors.Join(e.ErrAuthorDomain, e.ErrOnGetOne, err)
	}

	return u.validator.GetOne(id)
}

// Remove
func (u *userAuthorization) Remove(userId int, id int) error {
	err := u.auth.CheckForAuthorization(userId, p.UserRemove.Name)
	if err != nil {
		return errors.Join(e.ErrAuthorDomain, e.ErrOnRemove, err)
	}

	return u.validator.Remove(id)
}
