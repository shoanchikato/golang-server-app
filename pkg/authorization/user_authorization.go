package authorization

import (
	e "app/pkg/errors"
	m "app/pkg/model"
	p "app/pkg/permission"
	s "app/pkg/service"
	v "app/pkg/validation"
	"errors"
)

type UserAuthorization interface {
	Add(userId int, user *m.User) error
	AddAll(userId int, users *[]*m.User) error
	Edit(userId int, id int, newAuthor *m.EditUser) error
	GetAll(userId, lastId, limit int) (*[]m.User, error)
	GetOne(userId int, id int) (*m.User, error)
	Remove(userId int, id int) error
}

type userAuthorization struct {
	auth      s.AuthorizationService
	validator v.UserValidator
}

func NewUserAuthorization(auth s.AuthorizationService, validator v.UserValidator) UserAuthorization {
	return &userAuthorization{auth, validator}
}

// Add
func (u *userAuthorization) Add(userId int, user *m.User) error {
	err := u.auth.CheckForAuthorization(userId, p.UserAdd.Name)
	if err != nil {
		return errors.Join(e.ErrUserDomain, e.ErrOnAdd, err)
	}

	return u.validator.Add(user)
}

// AddAll
func (u *userAuthorization) AddAll(userId int, users *[]*m.User) error {
	err := u.auth.CheckForAuthorization(userId, p.UserAddAll.Name)
	if err != nil {
		return errors.Join(e.ErrUserDomain, e.ErrOnAddAll, err)
	}

	return u.validator.AddAll(users)
}

// Edit
func (u *userAuthorization) Edit(userId int, id int, user *m.EditUser) error {
	err := u.auth.CheckForAuthorization(userId, p.UserEdit.Name)
	if err != nil {
		return errors.Join(e.ErrUserDomain, e.ErrOnEdit, err)
	}

	return u.validator.Edit(id, user)
}

// GetAll
func (u *userAuthorization) GetAll(userId, lastId, limit int) (*[]m.User, error) {
	err := u.auth.CheckForAuthorization(userId, p.UserGetAll.Name)
	if err != nil {
		return nil, errors.Join(e.ErrUserDomain, e.ErrOnGetAll, err)
	}

	return u.validator.GetAll(lastId, limit)
}

// GetOne
func (u *userAuthorization) GetOne(userId int, id int) (*m.User, error) {
	err := u.auth.CheckForAuthorization(userId, p.UserGetOne.Name)
	if err != nil {
		return nil, errors.Join(e.ErrUserDomain, e.ErrOnGetOne, err)
	}

	return u.validator.GetOne(id)
}

// Remove
func (u *userAuthorization) Remove(userId int, id int) error {
	err := u.auth.CheckForAuthorization(userId, p.UserRemove.Name)
	if err != nil {
		return errors.Join(e.ErrUserDomain, e.ErrOnRemove, err)
	}

	return u.validator.Remove(id)
}
