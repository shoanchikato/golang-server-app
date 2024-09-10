package model

import (
	"fmt"

	valid "github.com/go-ozzo/ozzo-validation/v4"
)

type Book struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Year     int    `json:"year"`
	AuthorId int    `json:"author_id"`
}

func NewBook(name string, year, authorId int) *Book {
	return &Book{0, name, year, authorId}
}

func (b Book) String() string {
	return fmt.Sprintf(
		`%d, "%s", %d, %d`,
		b.Id, b.Name, b.Year, b.AuthorId,
	)
}

func (b *Book) Validate() error {
	return valid.ValidateStruct(b,
		valid.Field(&b.Name, valid.Required.Error("username is required")),
		valid.Field(&b.Year, valid.Required.Error("password is required")),
		valid.Field(&b.AuthorId, valid.Required.Error("author_id is required")),
	)
}
