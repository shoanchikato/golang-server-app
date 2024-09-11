package httperrorfmt

import (
	a "app/pkg/authorization"
	m "app/pkg/model"
	s "app/pkg/service"
)

type AuthorHttpErrorFmt interface {
	Add(userId int, author *m.Author) error
	AddAll(userId int, authors *[]*m.Author) error
	Edit(userId int, id int, newAuthor *m.Author) error
	GetAll(userId, lastId, limit int) (*[]m.Author, error)
	GetOne(userId int, id int) (*m.Author, error)
	Remove(userId int, id int) error
}

type authorHttpErrorFmt struct {
	authorization a.AuthorAuthorization
	service       s.HttpErrorFmt
}

func NewAuthorHttpErrorFmt(authorization a.AuthorAuthorization, service s.HttpErrorFmt) AuthorHttpErrorFmt {
	return &authorHttpErrorFmt{authorization, service}
}

// Add
func (r *authorHttpErrorFmt) Add(userId int, author *m.Author) error {
	err := r.authorization.Add(userId, author)
	if err != nil {
		return r.service.GetError(err)
	}

	return nil
}

// AddAll
func (r *authorHttpErrorFmt) AddAll(userId int, authors *[]*m.Author) error {
	err := r.authorization.AddAll(userId, authors)
	if err != nil {
		return r.service.GetError(err)
	}

	return nil
}

// Edit
func (r *authorHttpErrorFmt) Edit(userId int, id int, newAuthor *m.Author) error {
	err := r.authorization.Edit(userId, id, newAuthor)
	if err != nil {
		return r.service.GetError(err)
	}

	return nil
}

// GetAll
func (r *authorHttpErrorFmt) GetAll(userId int, lastId int, limit int) (*[]m.Author, error) {
	authors, err := r.authorization.GetAll(userId, lastId, limit)
	if err != nil {
		return nil, r.service.GetError(err)
	}

	return authors, nil
}

// GetOne
func (r *authorHttpErrorFmt) GetOne(userId int, id int) (*m.Author, error) {
	authors, err := r.authorization.GetOne(userId, id)
	if err != nil {
		return nil, r.service.GetError(err)
	}

	return authors, nil
}

// Remove
func (r *authorHttpErrorFmt) Remove(userId int, id int) error {
	err := r.authorization.Remove(userId, id)
	if err != nil {
		return r.service.GetError(err)
	}

	return nil
}
