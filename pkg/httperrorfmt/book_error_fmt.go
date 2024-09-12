package httperrorfmt

import (
	a "app/pkg/authorization"
	m "app/pkg/model"
	s "app/pkg/service"
)

type BookHttpErrorFmt interface {
	Add(userId int, book *m.Book) error
	AddAll(userId int, books *[]*m.Book) error
	Edit(userId int, id int, newBook *m.Book) error
	GetAll(userId, lastId, limit int) (*[]m.Book, error)
	GetOne(userId int, id int) (*m.Book, error)
	Remove(userId int, id int) error
}

type bookHttpErrorFmt struct {
	authorization a.BookAuthorization
	service       s.HttpErrorFmt
}

func NewBookHttpErrorFmt(authorization a.BookAuthorization, service s.HttpErrorFmt) BookHttpErrorFmt {
	return &bookHttpErrorFmt{authorization, service}
}

// Add
func (r *bookHttpErrorFmt) Add(userId int, book *m.Book) error {
	err := r.authorization.Add(userId, book)
	if err != nil {
		return r.service.GetError(err)
	}

	return nil
}

// AddAll
func (r *bookHttpErrorFmt) AddAll(userId int, books *[]*m.Book) error {
	err := r.authorization.AddAll(userId, books)
	if err != nil {
		return r.service.GetError(err)
	}

	return nil
}

// Edit
func (r *bookHttpErrorFmt) Edit(userId int, id int, newBook *m.Book) error {
	err := r.authorization.Edit(userId, id, newBook)
	if err != nil {
		return r.service.GetError(err)
	}

	return nil
}

// GetAll
func (r *bookHttpErrorFmt) GetAll(userId int, lastId int, limit int) (*[]m.Book, error) {
	books, err := r.authorization.GetAll(userId, lastId, limit)
	if err != nil {
		return nil, r.service.GetError(err)
	}

	return books, nil
}

// GetOne
func (r *bookHttpErrorFmt) GetOne(userId int, id int) (*m.Book, error) {
	books, err := r.authorization.GetOne(userId, id)
	if err != nil {
		return nil, r.service.GetError(err)
	}

	return books, nil
}

// Remove
func (r *bookHttpErrorFmt) Remove(userId int, id int) error {
	err := r.authorization.Remove(userId, id)
	if err != nil {
		return r.service.GetError(err)
	}

	return nil
}
