package model

import (
	"fmt"

	valid "github.com/go-ozzo/ozzo-validation/v4"
)

type Author struct {
	Id        int     `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Books     *[]Book `json:"books,omitempty"`
}

func (a Author) String() string {
	return fmt.Sprintf(
		`{%d, "%s", "%s", %v}`,
		a.Id, a.FirstName, a.LastName, a.Books,
	)
}

func NewAuthor(firstName, lastName string) *Author {
	return &Author{0, firstName, lastName, nil}
}

func (a *Author) Validate() error {
	return valid.ValidateStruct(a,
		valid.Field(&a.FirstName, valid.Required.Error("first_name is required")),
		valid.Field(&a.LastName, valid.Required.Error("last_name is required")),
	)
}
