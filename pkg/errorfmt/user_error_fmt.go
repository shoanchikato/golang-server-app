package errorfmt

import (
	a "app/pkg/authorization"
	m "app/pkg/model"
	s "app/pkg/service"
)

type UserErrorFmt interface {
	Add(userId int, user *m.User) error
	AddAll(userId int, users *[]*m.User) error
	Edit(userId int, id int, user *m.User) error
	GetAll(userId, lastId, limit int) (*[]m.User, error)
	GetOne(userId int, id int) (*m.User, error)
	Remove(userId int, id int) error
}

type userErrorFmt struct {
	auth    a.UserAuthorization
	service s.ErrorFmt
}

func NewUserErrorFmt(auth a.UserAuthorization, service s.ErrorFmt) UserErrorFmt {
	return &userErrorFmt{auth, service}
}

// Add
func (u *userErrorFmt) Add(userId int, user *m.User) error {
	err := u.auth.Add(userId, user)
	if err != nil {
		return u.service.GetError(err)
	}

	return nil
}

// AddAll
func (u *userErrorFmt) AddAll(userId int, users *[]*m.User) error {
	err := u.auth.AddAll(userId, users)
	if err != nil {
		return u.service.GetError(err)
	}

	return nil
}

// Edit
func (u *userErrorFmt) Edit(userId int, id int, user *m.User) error {
	err := u.auth.Edit(userId, id, user)
	if err != nil {
		return u.service.GetError(err)
	}

	return nil
}

// GetAll
func (u *userErrorFmt) GetAll(userId int, lastId int, limit int) (*[]m.User, error) {
	users, err := u.auth.GetAll(userId, lastId, limit)
	if err != nil {
		return nil, u.service.GetError(err)
	}

	return users, nil
}

// GetOne
func (u *userErrorFmt) GetOne(userId int, id int) (*m.User, error) {
	user, err := u.auth.GetOne(userId, id)
	if err != nil {
		return nil, u.service.GetError(err)
	}

	return user, nil
}

// Remove
func (u *userErrorFmt) Remove(userId int, id int) error {
	err := u.auth.Remove(userId, id)
	if err != nil {
		return u.service.GetError(err)
	}

	return nil
}
