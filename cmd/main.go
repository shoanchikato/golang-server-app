package main

import (
	a "app/pkg/auth"
	au "app/pkg/author"
	aa "app/pkg/authorization"
	b "app/pkg/book"
	pe "app/pkg/permission"
	p "app/pkg/post"
	r "app/pkg/repo"
	rr "app/pkg/role"
	s "app/pkg/service"
	u "app/pkg/user"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
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

	signingSecret := "my secret"
	_ = s.NewJWTService(
		&signingSecret,
		time.Duration(10*time.Second),
		time.Duration(1*time.Minute),
	)

	rw := &sync.RWMutex{}
	dbU := r.NewDBUtil(db, rw)
	en := s.NewEncryptionService()

	// Repos
	uRepo := u.NewUserRepo(db, rw, dbU)
	aRepo := a.NewAuthRepo(db, rw, dbU)
	rRepo := rr.NewRoleRepo(db, rw, dbU)
	peRepo := pe.NewPermissionRepo(db, rw, dbU)
	auRepo := au.NewAuthorRepo(db, rw, dbU)
	bRepo := b.NewBookRepo(db, rw, dbU)
	pRepo := p.NewPostRepo(db, rw, dbU)
	pmRepo := pe.NewPermissionManagementRepo(db, rw, dbU, uRepo, rRepo, peRepo)

	// Services
	aEncrypt := a.NewAuthEncryption(aRepo, en)
	uEncrypt := u.NewUserEncryption(uRepo, en)
	auth := aa.NewAuthorizationService(pmRepo)

	// Validators
	_ = u.NewUserValidator(uEncrypt)
	_ = a.NewAuthValidator(aEncrypt)
	_ = rr.NewRoleValidator(rRepo)
	_ = pe.NewPermissionValidator(peRepo)
	_ = au.NewAuthorValidator(auRepo)
	_ = b.NewBookValidator(bRepo)
	pVal := p.NewPostValidator(pRepo)
	_ = pe.NewPermissionManagementValidator(pmRepo)

	// Authorization
	_ = p.NewPostAuthorization(pVal, auth)
	srv := b.NewBookAuthorization(auth, bRepo)

	// _, _, posts, _ := Data()

	values := os.Args
	value, err := strconv.Atoi(values[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	posts, err := srv.GetAll(value)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%v\n%v\n", value, posts)
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
		p.NewPost("five", "five body", 1),
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

	_ = []*pe.Permission{
		pe.NewPermission("post: app"),
		pe.NewPermission("post: add all"),
		pe.NewPermission("post: edit"),
		pe.NewPermission("post: remove"),
	}

	return authors, books, posts, users
}
