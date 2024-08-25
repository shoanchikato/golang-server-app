package main

import (
	m "app/pkg/model"
	r "app/pkg/repo"
	s "app/pkg/service"
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

	authRepo := r.NewAuthRepo(db, rw, dbU)
	authorRepo := r.NewAuthorRepo(db, rw, dbU)
	bookRepo := r.NewBookRepo(db, rw, dbU)
	postRepo := r.NewPostRepo(db, rw, dbU)
	userRepo := r.NewUserRepo(db, rw, dbU)

	encrypt := s.NewEncryptionService()
	authEncrypt := s.NewAuthEncryption(authRepo, encrypt)
	userEncrypt := s.NewUserEncryption(userRepo, encrypt)

	signingSecret := "my secret"
	_ = s.NewJWTService(
		&signingSecret,
		time.Duration(10*time.Second),
		time.Duration(1*time.Minute),
	)

	_ = s.NewAuthValidator(authEncrypt)
	_ = s.NewAuthorValidator(authorRepo)
	_ = s.NewBookValidator(bookRepo)
	_ = s.NewPostValidator(postRepo)
	_ = s.NewUserValidator(userEncrypt)

	// _, _, _, _ = Data()

	// err = jwt.GetAccessToken()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%v\n", err)
}

func Data() (
	[]*m.Author,
	[]*m.Book,
	[]*m.Post,
	[]*m.User,
) {
	posts := []*m.Post{
		m.NewPost("one", "one body", 1),
		m.NewPost("two", "two body", 1),
		m.NewPost("three", "three body", 1),
		m.NewPost("four", "four body", 1),
		m.NewPost("five", "four", 1),
	}

	authors := []*m.Author{
		m.NewAuthor("John", "Doe"),
		m.NewAuthor("Jane", "Doe"),
		m.NewAuthor("James", "Doe"),
	}

	books := []*m.Book{
		m.NewBook("one book", 2010, 1),
		m.NewBook("two book", 2018, 2),
		m.NewBook("three book", 2004, 3),
		m.NewBook("four book", 2014, 1),
	}

	users := []*m.User{
		m.NewUser("John", "Doe", "john_doe", "john@doe.com", "password1"),
		m.NewUser("Jenny", "Doe", "jenny_doe", "jenny@doe.com", "password2"),
		m.NewUser("Jim", "Doe", "jim_doe", "jim@doe.com", "password3"),
	}

	return authors, books, posts, users
}
