package model

import valid "github.com/go-ozzo/ozzo-validation/v4"

type Post struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserId int    `json:"user_id"`
}

func NewPost(title, body string, userId int) *Post {
	return &Post{0, title, body, userId}
}

func (p *Post) Validate() error {
	return valid.ValidateStruct(p,
		valid.Field(&p.Title, valid.Required.Error("title is required")),
		valid.Field(&p.Body, valid.Required.Error("body is required")),
		valid.Field(&p.UserId, valid.Required.Error("user_id is required")),
	)
}
