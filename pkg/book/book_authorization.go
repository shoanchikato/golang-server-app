package book

import (
	a "app/pkg/authorization"
	e "app/pkg/errors"
	"errors"
)

type BookAuthorization interface {
	Add(userID int, book *Book) error
	AddAll(userID int, books *[]*Book) error
	Edit(userID int, id int, newBook *Book) error
	GetAll(userID int) (*[]Book, error)
	GetOne(userID int, id int) (*Book, error)
	Remove(userID int, id int) error
}

type bookAuthorization struct {
	auth a.AuthorizationService
	v    BookValidator
}

func NewBookAuthorization(auth a.AuthorizationService, v BookValidator) BookAuthorization {
	return &bookAuthorization{auth, v}
}

// Add
func (b *bookAuthorization) Add(userID int, book *Book) error {
	err := b.auth.CheckForAuthorization(userID, BookAdd.Name)
	if err != nil {
		return errors.Join(e.ErrBookDomain, e.ErrOnAdd, err)
	}

	return b.v.Add(book)
}

// AddAll
func (b *bookAuthorization) AddAll(userID int, books *[]*Book) error {
	err := b.auth.CheckForAuthorization(userID, BookAddAll.Name)
	if err != nil {
		return errors.Join(e.ErrBookDomain, e.ErrOnAddAll, err)
	}

	return b.v.AddAll(books)
}

// Edit
func (b *bookAuthorization) Edit(userID int, id int, newBook *Book) error {
	err := b.auth.CheckForAuthorization(userID, BookEdit.Name)
	if err != nil {
		return errors.Join(e.ErrBookDomain, e.ErrOnEdit, err)
	}

	return b.v.Edit(id, newBook)
}

// GetAll
func (b *bookAuthorization) GetAll(userID int) (*[]Book, error) {
	err := b.auth.CheckForAuthorization(userID, BookGetAll.Name)
	if err != nil {
		return nil, errors.Join(e.ErrBookDomain, e.ErrOnGetAll, err)
	}

	return b.v.GetAll()
}

// GetOne
func (b *bookAuthorization) GetOne(userID int, id int) (*Book, error) {
	err := b.auth.CheckForAuthorization(userID, BookGetOne.Name)
	if err != nil {
		return nil, errors.Join(e.ErrBookDomain, e.ErrOnGetOne, err)
	}

	return b.v.GetOne(id)
}

// Remove
func (b *bookAuthorization) Remove(userID int, id int) error {
	err := b.auth.CheckForAuthorization(userID, BookRemove.Name)
	if err != nil {
		return errors.Join(e.ErrBookDomain, e.ErrOnRemove, err)
	}

	return b.v.Remove(id)
}
