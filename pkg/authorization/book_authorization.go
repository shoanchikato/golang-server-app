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
	Add(userID int, book *m.Book) error
	AddAll(userID int, books *[]*m.Book) error
	Edit(userID int, id int, newBook *m.Book) error
	GetAll(userID int) (*[]m.Book, error)
	GetOne(userID int, id int) (*m.Book, error)
	Remove(userID int, id int) error
}

type bookAuthorization struct {
	auth s.AuthorizationService
	v    v.BookValidator
}

func NewBookAuthorization(auth s.AuthorizationService, v v.BookValidator) BookAuthorization {
	return &bookAuthorization{auth, v}
}

// Add
func (b *bookAuthorization) Add(userID int, book *m.Book) error {
	err := b.auth.CheckForAuthorization(userID, p.BookAdd.Name)
	if err != nil {
		return errors.Join(e.ErrBookDomain, e.ErrOnAdd, err)
	}

	return b.v.Add(book)
}

// AddAll
func (b *bookAuthorization) AddAll(userID int, books *[]*m.Book) error {
	err := b.auth.CheckForAuthorization(userID, p.BookAddAll.Name)
	if err != nil {
		return errors.Join(e.ErrBookDomain, e.ErrOnAddAll, err)
	}

	return b.v.AddAll(books)
}

// Edit
func (b *bookAuthorization) Edit(userID int, id int, newBook *m.Book) error {
	err := b.auth.CheckForAuthorization(userID, p.BookEdit.Name)
	if err != nil {
		return errors.Join(e.ErrBookDomain, e.ErrOnEdit, err)
	}

	return b.v.Edit(id, newBook)
}

// GetAll
func (b *bookAuthorization) GetAll(userID int) (*[]m.Book, error) {
	err := b.auth.CheckForAuthorization(userID, p.BookGetAll.Name)
	if err != nil {
		return nil, errors.Join(e.ErrBookDomain, e.ErrOnGetAll, err)
	}

	return b.v.GetAll()
}

// GetOne
func (b *bookAuthorization) GetOne(userID int, id int) (*m.Book, error) {
	err := b.auth.CheckForAuthorization(userID, p.BookGetOne.Name)
	if err != nil {
		return nil, errors.Join(e.ErrBookDomain, e.ErrOnGetOne, err)
	}

	return b.v.GetOne(id)
}

// Remove
func (b *bookAuthorization) Remove(userID int, id int) error {
	err := b.auth.CheckForAuthorization(userID, p.BookRemove.Name)
	if err != nil {
		return errors.Join(e.ErrBookDomain, e.ErrOnRemove, err)
	}

	return b.v.Remove(id)
}
