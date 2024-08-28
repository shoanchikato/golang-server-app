package model

type Book struct {
	ID       int    `json:"id"`
	Name     string `json:"name" valid:"required~name is required"`
	Year     int    `json:"year" valid:"required~year is required"`
	AuthorID int    `json:"author_id"`
}

func NewBook(name string, year, authorID int) *Book {
	return &Book{0, name, year, authorID}
}
