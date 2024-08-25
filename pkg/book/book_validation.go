package book

import (
	e "app/pkg/errors"
	"fmt"
	"strings"

	valid "github.com/asaskevich/govalidator"
)

type BookValidator interface {
	Add(book *Book) error
	AddAll(books *[]*Book) error
	Edit(id int, newBook *Book) error
	GetAll() (*[]Book, error)
	GetOne(id int) (*Book, error)
	Remove(id int) error
}

type bookValidator struct {
	Repo BookRepo
}

func NewBookValidator(repo BookRepo) BookValidator {
	return &bookValidator{repo}
}

// Add
func (v *bookValidator) Add(book *Book) error {
	_, err := valid.ValidateStruct(book)
	if err != nil {
		return e.NewValidationError(e.ErrAddValidation, err.Error())
	}

	err = v.Repo.Add(book)
	if err != nil {
		return err
	}

	return nil
}

// AddAll
func (v *bookValidator) AddAll(books *[]*Book) error {
	errs := []string{}
	newBooks := *books
	for i := 0; i < len(newBooks); i++ {
		_, err := valid.ValidateStruct(newBooks[i])
		if err != nil {
			errStr := fmt.Sprintf("\n[%d] %s", i, err.Error())
			errs = append(errs, errStr)
		}
	}

	if len(errs) > 0 {
		newErrors := strings.Join(errs, "")
		return e.NewValidationError(e.ErrAddAllValidation, newErrors)
	}

	err := v.Repo.AddAll(books)
	if err != nil {
		return err
	}

	return nil
}

// Edit
func (v *bookValidator) Edit(id int, newBook *Book) error {
	_, err := valid.ValidateStruct(newBook)
	if err != nil {
		return e.NewValidationError(e.ErrEditValidation, err.Error())
	}

	err = v.Repo.Edit(id, newBook)
	if err != nil {
		return err
	}

	return nil
}

// GetAll
func (v *bookValidator) GetAll() (*[]Book, error) {
	return v.Repo.GetAll()
}

// GetOne
func (v *bookValidator) GetOne(id int) (*Book, error) {
	return v.Repo.GetOne(id)
}

// Remove
func (v *bookValidator) Remove(id int) error {
	return v.Repo.Remove(id)
}
