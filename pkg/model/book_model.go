package model

type Book struct {
	Id       int    `json:"id"`
	Name     string `json:"name" valid:"required~name is required"`
	Year     int    `json:"year" valid:"required~year is required"`
	AuthorId int    `json:"author_id"`
}

func NewBook(name string, year, authorId int) *Book {
	return &Book{0, name, year, authorId}
}
