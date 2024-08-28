package main

import (
	a "app/pkg/authorization"
	e "app/pkg/encrypt"
	m "app/pkg/model"
	pe "app/pkg/permission"
	r "app/pkg/repo"
	s "app/pkg/service"
	v "app/pkg/validation"
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

	signingSecret := "my secret"
	_ = s.NewJWTService(
		&signingSecret,
		time.Duration(10*time.Second),
		time.Duration(1*time.Minute),
	)

	rw := &sync.RWMutex{}
	dbU := r.NewDBUtil(db, rw)

	// Repos
	uRepo := r.NewUserRepo(db, rw, dbU)
	aRepo := r.NewAuthRepo(db, rw, dbU)
	rRepo := r.NewRoleRepo(db, rw, dbU)
	peRepo := r.NewPermissionRepo(db, rw, dbU)
	auRepo := r.NewAuthorRepo(db, rw, dbU)
	bRepo := r.NewBookRepo(db, rw, dbU)
	pRepo := r.NewPostRepo(db, rw, dbU)
	pmRepo := r.NewPermissionManagementRepo(db, rw, dbU, uRepo, rRepo, peRepo)

	// Encrypt
	en := s.NewEncryptionService()
	aEncrypt := e.NewAuthEncryption(aRepo, en)
	uEncrypt := e.NewUserEncryption(uRepo, en)

	// Validators
	_ = v.NewUserValidator(uEncrypt)
	_ = v.NewAuthValidator(aEncrypt)
	_ = v.NewRoleValidator(rRepo)
	srv := v.NewPermissionValidator(peRepo)
	_ = v.NewAuthorValidator(auRepo)
	bVal := v.NewBookValidator(bRepo)
	pVal := v.NewPostValidator(pRepo)
	_ = v.NewPermissionManagementValidator(pmRepo)

	// Authorization
	auth := s.NewAuthorizationService(pmRepo)
	_ = a.NewPostAuthorization(auth, pVal)
	_ = a.NewBookAuthorization(auth, bVal)

	// _, _, _, _, _ = Data()

	pp, err := srv.GetAll()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(pp)
}

func Data() (
	[]*m.Author,
	[]*m.Book,
	[]*m.Post,
	[]*m.User,
	[]*m.Permission,
) {
	posts := []*m.Post{
		m.NewPost("one", "one body", 1),
		m.NewPost("two", "two body", 1),
		m.NewPost("three", "three body", 1),
		m.NewPost("four", "four body", 1),
		m.NewPost("five", "five body", 1),
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

	postPermissions := []*m.Permission{
		pe.PostAdd,
		pe.PostAddAll,
		pe.PostEdit,
		pe.PostGetOne,
		pe.PostGetAll,
		pe.PostRemove,
	}

	_ = []*m.Permission{
		pe.BookAdd,
		pe.BookAddAll,
		pe.BookEdit,
		pe.BookGetOne,
		pe.BookGetAll,
		pe.BookRemove,
	}

	return authors, books, posts, users, postPermissions
}
