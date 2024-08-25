package service

import (
	e "app/pkg/errors"
	m "app/pkg/model"
	r "app/pkg/repo"
	"fmt"
	"strings"

	valid "github.com/asaskevich/govalidator"
)

type AuthorValidator interface {
	Add(author *m.Author) error
	AddAll(authors *[]*m.Author) error
	Edit(id int, newAuthor *m.Author) error
	GetAll() (*[]m.Author, error)
	GetOne(id int) (*m.Author, error)
	Remove(id int) error
}

type authorValidator struct {
	Repo r.AuthorRepo
}

func NewAuthorValidator(repo r.AuthorRepo) AuthorValidator {
	return &authorValidator{repo}
}

// Add
func (v *authorValidator) Add(author *m.Author) error {
	_, err := valid.ValidateStruct(author)
	if err != nil {
		return e.NewValidationError(e.ErrAddValidation, err.Error())
	}

	err = v.Repo.Add(author)
	if err != nil {
		return err
	}

	return nil
}

// AddAll
func (v *authorValidator) AddAll(authors *[]*m.Author) error {
	errs := []string{}
	newAuthors := *authors
	for i := 0; i < len(newAuthors); i++ {
		_, err := valid.ValidateStruct(newAuthors[i])
		if err != nil {
			errStr := fmt.Sprintf("\n[%d] %s", i, err.Error())
			errs = append(errs, errStr)
		}
	}

	if len(errs) > 0 {
		newErrors := strings.Join(errs, "")
		return e.NewValidationError(e.ErrAddAllValidation, newErrors)
	}

	err := v.Repo.AddAll(authors)
	if err != nil {
		return err
	}

	return nil
}

// Edit
func (v *authorValidator) Edit(id int, newAuthor *m.Author) error {
	_, err := valid.ValidateStruct(newAuthor)
	if err != nil {
		return e.NewValidationError(e.ErrEditValidation, err.Error())
	}

	err = v.Repo.Edit(id, newAuthor)
	if err != nil {
		return err
	}

	return nil
}

// GetAll
func (v *authorValidator) GetAll() (*[]m.Author, error) {
	return v.Repo.GetAll()
}

// GetOne
func (v *authorValidator) GetOne(id int) (*m.Author, error) {
	return v.Repo.GetOne(id)
}

// Remove
func (v *authorValidator) Remove(id int) error {
	return v.Repo.Remove(id)
}
