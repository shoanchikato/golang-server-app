package model

type Post struct {
	ID     int    `json:"id"`
	Title  string `json:"title" valid:"required~title is required"`
	Body   string `json:"body" valid:"required~body is required"`
	UserID int    `json:"user_id" valid:"required~user_id is required"`
}

func NewPost(title, body string, userID int) *Post {
	return &Post{0, title, body, userID}
}
