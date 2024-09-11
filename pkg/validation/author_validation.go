package validation

import (
	e "app/pkg/errors"
	m "app/pkg/model"
	r "app/pkg/repo"
	s "app/pkg/service"
)

type AuthorValidator interface {
	Add(author *m.Author) error
	AddAll(authors *[]*m.Author) error
	Edit(id int, newAuthor *m.Author) error
	GetAll(lastId, limit int) (*[]m.Author, error)
	GetOne(id int) (*m.Author, error)
	Remove(id int) error
}

type authorValidator struct {
	Repo    r.AuthorRepo
	service s.ValidationService
}

func NewAuthorValidator(repo r.AuthorRepo, service s.ValidationService) AuthorValidator {
	return &authorValidator{repo, service}
}

// Add
func (v *authorValidator) Add(author *m.Author) error {
	err := v.service.Validate(author)
	if err != nil {
		return e.NewValidationError(e.ErrAddValidation, err)
	}

	err = v.Repo.Add(author)
	if err != nil {
		return err
	}

	return nil
}

// AddAll
func (v *authorValidator) AddAll(authors *[]*m.Author) error {
	newAuthors := *authors
	errs := make([]error, 0, len(newAuthors))
	for i := range newAuthors {
		err := newAuthors[i].Validate()
		if err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return e.NewValidationError(e.ErrAddAllValidation, errs...)
	}

	err := v.Repo.AddAll(authors)
	if err != nil {
		return err
	}

	return nil
}

// Edit
func (v *authorValidator) Edit(id int, newAuthor *m.Author) error {
	err := newAuthor.Validate()
	if err != nil {
		return e.NewValidationError(e.ErrEditValidation, err)
	}

	err = v.Repo.Edit(id, newAuthor)
	if err != nil {
		return err
	}

	return nil
}

// GetAll
func (v *authorValidator) GetAll(lastId, limit int) (*[]m.Author, error) {
	return v.Repo.GetAll(lastId, limit)
}

// GetOne
func (v *authorValidator) GetOne(id int) (*m.Author, error) {
	return v.Repo.GetOne(id)
}

// Remove
func (v *authorValidator) Remove(id int) error {
	return v.Repo.Remove(id)
}
