package main

import (
	a "app/pkg/auth"
	au "app/pkg/author"
	b "app/pkg/book"
	p "app/pkg/post"
	r "app/pkg/repo"
	s "app/pkg/service"
	u "app/pkg/user"
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "small.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rw := &sync.RWMutex{}
	dbU := r.NewDBUtil(db, rw)
	en := s.NewEncryptionService()

	signingSecret := "my secret"
	_ = s.NewJWTService(
		&signingSecret,
		time.Duration(10*time.Second),
		time.Duration(1*time.Minute),
	)

	_ = a.Di(db, rw, dbU, en)
	_ = au.Di(db, rw, dbU)
	_ = b.Di(db, rw, dbU)
	_ = p.Di(db, rw, dbU)
	_ = u.Di(db, rw, dbU, en)

	// _, _, _, _ = Data()

	// err = jwt.GetAccessToken()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%v\n", err)
}

func Data() (
	[]*au.Author,
	[]*b.Book,
	[]*p.Post,
	[]*u.User,
) {
	posts := []*p.Post{
		p.NewPost("one", "one body", 1),
		p.NewPost("two", "two body", 1),
		p.NewPost("three", "three body", 1),
		p.NewPost("four", "four body", 1),
		p.NewPost("five", "four", 1),
	}

	authors := []*au.Author{
		au.NewAuthor("John", "Doe"),
		au.NewAuthor("Jane", "Doe"),
		au.NewAuthor("James", "Doe"),
	}

	books := []*b.Book{
		b.NewBook("one book", 2010, 1),
		b.NewBook("two book", 2018, 2),
		b.NewBook("three book", 2004, 3),
		b.NewBook("four book", 2014, 1),
	}

	users := []*u.User{
		u.NewUser("John", "Doe", "john_doe", "john@doe.com", "password1"),
		u.NewUser("Jenny", "Doe", "jenny_doe", "jenny@doe.com", "password2"),
		u.NewUser("Jim", "Doe", "jim_doe", "jim@doe.com", "password3"),
	}

	return authors, books, posts, users
}
