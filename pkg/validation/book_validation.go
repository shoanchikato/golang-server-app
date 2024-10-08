package validation

import (
	e "app/pkg/errors"
	m "app/pkg/model"
	r "app/pkg/repo"
	s "app/pkg/service"
)

type BookValidator interface {
	Add(book *m.Book) error
	AddAll(books *[]*m.Book) error
	Edit(id int, newBook *m.Book) error
	GetAll(lastId, limit int) (*[]m.Book, error)
	GetOne(id int) (*m.Book, error)
	Remove(id int) error
}

type bookValidator struct {
	Repo    r.BookRepo
	service s.ValidationService
}

func NewBookValidator(repo r.BookRepo, service s.ValidationService) BookValidator {
	return &bookValidator{repo, service}
}

// Add
func (v *bookValidator) Add(book *m.Book) error {
	err := v.service.Validate(book)
	if err != nil {
		return e.NewValidationError(e.ErrAddValidation, err)
	}

	err = v.Repo.Add(book)
	if err != nil {
		return err
	}

	return nil
}

// AddAll
func (v *bookValidator) AddAll(books *[]*m.Book) error {
	newBooks := *books
	errs := make([]error, 0, len(newBooks))
	for i := range newBooks {
		err := v.service.Validate(newBooks[i])
		if err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return e.NewValidationError(e.ErrAddAllValidation, errs...)
	}

	err := v.Repo.AddAll(books)
	if err != nil {
		return err
	}

	return nil
}

// Edit
func (v *bookValidator) Edit(id int, newBook *m.Book) error {
	err := v.service.Validate(newBook)
	if err != nil {
		return e.NewValidationError(e.ErrEditValidation, err)
	}

	err = v.Repo.Edit(id, newBook)
	if err != nil {
		return err
	}

	return nil
}

// GetAll
func (v *bookValidator) GetAll(lastId, limit int) (*[]m.Book, error) {
	return v.Repo.GetAll(lastId, limit)
}

// GetOne
func (v *bookValidator) GetOne(id int) (*m.Book, error) {
	return v.Repo.GetOne(id)
}

// Remove
func (v *bookValidator) Remove(id int) error {
	return v.Repo.Remove(id)
}
