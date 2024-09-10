package model

type Post struct {
	Id     int    `json:"id"`
	Title  string `json:"title" validate:"required~title is required"`
	Body   string `json:"body" validate:"required~body is required"`
	UserId int    `json:"user_id" validate:"required~user_id is required"`
}

func NewPost(title, body string, userId int) *Post {
	return &Post{0, title, body, userId}
}
