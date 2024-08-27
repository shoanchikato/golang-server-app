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
	srv := pe.NewPermissionValidator(peRepo)
	_ = au.NewAuthorValidator(auRepo)
	bVal := b.NewBookValidator(bRepo)
	pVal := p.NewPostValidator(pRepo)
	_ = pe.NewPermissionManagementValidator(pmRepo)

	// Authorization
	_ = p.NewPostAuthorization(auth, pVal)
	_ = b.NewBookAuthorization(auth, bVal)

	// _, _, _, _, _ = Data()
	pp := []*pe.Permission{
		pe.NewPermission("fsd", "lk", "fsd"), pe.NewPermission("", "", ""),
	}
	err = srv.AddAll(&pp)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(err)
}

func Data() (
	[]*au.Author,
	[]*b.Book,
	[]*p.Post,
	[]*u.User,
	[]*pe.Permission,
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

	postPermissions := []*pe.Permission{
		pe.NewPermission(string(p.PostAdd), "post", "add"),
		pe.NewPermission(string(p.PostAddAll), "post", "add all"),
		pe.NewPermission(string(p.PostEdit), "post", "edit"),
		pe.NewPermission(string(p.PostGetOne), "post", "get one"),
		pe.NewPermission(string(p.PostGetAll), "post", "get all"),
		pe.NewPermission(string(p.PostRemove), "post", "remove"),
	}

	_ = []*pe.Permission{
		pe.NewPermission(string(b.BookAdd), "book", "add"),
		pe.NewPermission(string(b.BookAddAll), "book", "add all"),
		pe.NewPermission(string(b.BookEdit), "book", "edit"),
		pe.NewPermission(string(b.BookGetOne), "book", "get one"),
		pe.NewPermission(string(b.BookGetAll), "book", "get all"),
		pe.NewPermission(string(b.BookRemove), "book", "remove"),
	}

	return authors, books, posts, users, postPermissions
}
