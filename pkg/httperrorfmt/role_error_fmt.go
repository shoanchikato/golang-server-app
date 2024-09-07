package httperrorfmt

import (
	a "app/pkg/authorization"
	m "app/pkg/model"
	s "app/pkg/service"
)

type RoleHttpErrorFmt interface {
	Add(userId int, role *m.Role) error
	AddAll(userId int, roles *[]*m.Role) error
	Edit(userId int, id int, newAuthor *m.Role) error
	GetAll(userId, lastId, limit int) (*[]m.Role, error)
	GetOne(userId int, id int) (*m.Role, error)
	Remove(userId int, id int) error
}

type roleHttpErrorFmt struct {
	authorization a.RoleAuthorization
	service       s.HttpErrorFmt
}

func NewRoleHttpErrorFmt(authorization a.RoleAuthorization, service s.HttpErrorFmt) RoleHttpErrorFmt {
	return &roleHttpErrorFmt{authorization, service}
}

// Add
func (r *roleHttpErrorFmt) Add(userId int, role *m.Role) error {
	err := r.authorization.Add(userId, role)
	if err != nil {
		return r.service.GetError(err)
	}

	return nil
}

// AddAll
func (r *roleHttpErrorFmt) AddAll(userId int, roles *[]*m.Role) error {
	err := r.authorization.AddAll(userId, roles)
	if err != nil {
		return r.service.GetError(err)
	}

	return nil
}

// Edit
func (r *roleHttpErrorFmt) Edit(userId int, id int, newAuthor *m.Role) error {
	err := r.authorization.Edit(userId, id, newAuthor)
	if err != nil {
		return r.service.GetError(err)
	}

	return nil
}

// GetAll
func (r *roleHttpErrorFmt) GetAll(userId int, lastId int, limit int) (*[]m.Role, error) {
	roles, err := r.authorization.GetAll(userId, lastId, limit)
	if err != nil {
		return nil, r.service.GetError(err)
	}

	return roles, nil
}

// GetOne
func (r *roleHttpErrorFmt) GetOne(userId int, id int) (*m.Role, error) {
	roles, err := r.authorization.GetOne(userId, id)
	if err != nil {
		return nil, r.service.GetError(err)
	}

	return roles, nil
}

// Remove
func (r *roleHttpErrorFmt) Remove(userId int, id int) error {
	err := r.authorization.Remove(userId, id)
	if err != nil {
		return r.service.GetError(err)
	}

	return nil
}
