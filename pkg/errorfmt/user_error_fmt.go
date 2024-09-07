package errorfmt

import (
	a "app/pkg/authorization"
	m "app/pkg/model"
	s "app/pkg/service"
)

type UserHttpErrorFmt interface {
	Add(userId int, user *m.User) error
	AddAll(userId int, users *[]*m.User) error
	Edit(userId int, id int, user *m.EditUser) error
	GetAll(userId, lastId, limit int) (*[]m.User, error)
	GetOne(userId int, id int) (*m.User, error)
	Remove(userId int, id int) error
}

type userHttpErrorFmt struct {
	auth    a.UserAuthorization
	service s.HttpErrorFmt
}

func NewUserHttpErrorFmt(auth a.UserAuthorization, service s.HttpErrorFmt) UserHttpErrorFmt {
	return &userHttpErrorFmt{auth, service}
}

// Add
func (u *userHttpErrorFmt) Add(userId int, user *m.User) error {
	err := u.auth.Add(userId, user)
	if err != nil {
		return u.service.GetError(err)
	}

	return nil
}

// AddAll
func (u *userHttpErrorFmt) AddAll(userId int, users *[]*m.User) error {
	err := u.auth.AddAll(userId, users)
	if err != nil {
		return u.service.GetError(err)
	}

	return nil
}

// Edit
func (u *userHttpErrorFmt) Edit(userId int, id int, user *m.EditUser) error {
	err := u.auth.Edit(userId, id, user)
	if err != nil {
		return u.service.GetError(err)
	}

	return nil
}

// GetAll
func (u *userHttpErrorFmt) GetAll(userId int, lastId int, limit int) (*[]m.User, error) {
	users, err := u.auth.GetAll(userId, lastId, limit)
	if err != nil {
		return nil, u.service.GetError(err)
	}

	return users, nil
}

// GetOne
func (u *userHttpErrorFmt) GetOne(userId int, id int) (*m.User, error) {
	user, err := u.auth.GetOne(userId, id)
	if err != nil {
		return nil, u.service.GetError(err)
	}

	return user, nil
}

// Remove
func (u *userHttpErrorFmt) Remove(userId int, id int) error {
	err := u.auth.Remove(userId, id)
	if err != nil {
		return u.service.GetError(err)
	}

	return nil
}
