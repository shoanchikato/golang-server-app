package authorization

import (
	e "app/pkg/errors"
	m "app/pkg/model"
	p "app/pkg/permission"
	s "app/pkg/service"
	v "app/pkg/validation"
	"errors"
)

type BookAuthorization interface {
	Add(userId int, book *m.Book) error
	AddAll(userId int, books *[]*m.Book) error
	Edit(userId int, id int, newBook *m.Book) error
	GetAll(userId, lastId, limit int) (*[]m.Book, error)
	GetOne(userId int, id int) (*m.Book, error)
	Remove(userId int, id int) error
}

type bookAuthorization struct {
	auth s.AuthorizationService
	v    v.BookValidator
}

func NewBookAuthorization(auth s.AuthorizationService, v v.BookValidator) BookAuthorization {
	return &bookAuthorization{auth, v}
}

// Add
func (b *bookAuthorization) Add(userId int, book *m.Book) error {
	err := b.auth.CheckForAuthorization(userId, p.BookAdd.Name)
	if err != nil {
		return errors.Join(e.ErrBookDomain, e.ErrOnAdd, err)
	}

	return b.v.Add(book)
}

// AddAll
func (b *bookAuthorization) AddAll(userId int, books *[]*m.Book) error {
	err := b.auth.CheckForAuthorization(userId, p.BookAddAll.Name)
	if err != nil {
		return errors.Join(e.ErrBookDomain, e.ErrOnAddAll, err)
	}

	return b.v.AddAll(books)
}

// Edit
func (b *bookAuthorization) Edit(userId int, id int, newBook *m.Book) error {
	err := b.auth.CheckForAuthorization(userId, p.BookEdit.Name)
	if err != nil {
		return errors.Join(e.ErrBookDomain, e.ErrOnEdit, err)
	}

	return b.v.Edit(id, newBook)
}

// GetAll
func (b *bookAuthorization) GetAll(userId, lastId, limit int) (*[]m.Book, error) {
	err := b.auth.CheckForAuthorization(userId, p.BookGetAll.Name)
	if err != nil {
		return nil, errors.Join(e.ErrBookDomain, e.ErrOnGetAll, err)
	}

	return b.v.GetAll(lastId, limit)
}

// GetOne
func (b *bookAuthorization) GetOne(userId int, id int) (*m.Book, error) {
	err := b.auth.CheckForAuthorization(userId, p.BookGetOne.Name)
	if err != nil {
		return nil, errors.Join(e.ErrBookDomain, e.ErrOnGetOne, err)
	}

	return b.v.GetOne(id)
}

// Remove
func (b *bookAuthorization) Remove(userId int, id int) error {
	err := b.auth.CheckForAuthorization(userId, p.BookRemove.Name)
	if err != nil {
		return errors.Join(e.ErrBookDomain, e.ErrOnRemove, err)
	}

	return b.v.Remove(id)
}
