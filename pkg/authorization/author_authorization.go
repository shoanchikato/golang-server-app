package authorization

import (
	e "app/pkg/errors"
	m "app/pkg/model"
	p "app/pkg/permission"
	s "app/pkg/service"
	v "app/pkg/validation"
	"errors"
)

type AuthorAuthorization interface {
	Add(userId int, author *m.Author) error
	AddAll(userId int, authors *[]*m.Author) error
	Edit(userId int, id int, newAuthor *m.Author) error
	GetAll(userId int) (*[]m.Author, error)
	GetOne(userId int, id int) (*m.Author, error)
	Remove(userId int, id int) error
}

type authorAuthorization struct {
	auth      s.AuthorizationService
	validator v.AuthorValidator
}

func NewAuthorAuthorization(auth s.AuthorizationService, validator v.AuthorValidator) AuthorAuthorization {
	return &authorAuthorization{auth, validator}
}

// Add
func (a *authorAuthorization) Add(userId int, author *m.Author) error {
	err := a.auth.CheckForAuthorization(userId, p.AuthorAdd.Name)
	if err != nil {
		return errors.Join(e.ErrAuthorDomain, e.ErrOnAdd, err)
	}

	return a.validator.Add(author)
}

// AddAll
func (a *authorAuthorization) AddAll(userId int, authors *[]*m.Author) error {
	err := a.auth.CheckForAuthorization(userId, p.AuthorAddAll.Name)
	if err != nil {
		return errors.Join(e.ErrAuthorDomain, e.ErrOnAddAll, err)
	}

	return a.validator.AddAll(authors)
}

// Edit
func (a *authorAuthorization) Edit(userId int, id int, newAuthor *m.Author) error {
	err := a.auth.CheckForAuthorization(userId, p.AuthorEdit.Name)
	if err != nil {
		return errors.Join(e.ErrAuthorDomain, e.ErrOnEdit, err)
	}

	return a.validator.Edit(id, newAuthor)
}

// GetAll
func (a *authorAuthorization) GetAll(userId int) (*[]m.Author, error) {
	err := a.auth.CheckForAuthorization(userId, p.AuthorGetAll.Name)
	if err != nil {
		return nil, errors.Join(e.ErrAuthorDomain, e.ErrOnGetAll, err)
	}

	return a.validator.GetAll()
}

// GetOne
func (a *authorAuthorization) GetOne(userId int, id int) (*m.Author, error) {
	err := a.auth.CheckForAuthorization(userId, p.AuthorGetOne.Name)
	if err != nil {
		return nil, errors.Join(e.ErrAuthorDomain, e.ErrOnGetOne, err)
	}

	return a.validator.GetOne(id)
}

// Remove
func (a *authorAuthorization) Remove(userId int, id int) error {
	err := a.auth.CheckForAuthorization(userId, p.AuthorRemove.Name)
	if err != nil {
		return errors.Join(e.ErrAuthorDomain, e.ErrOnRemove, err)
	}

	return a.validator.Remove(id)
}
