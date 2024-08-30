package di

import (
	a "app/pkg/authorization"
	e "app/pkg/encrypt"
	r "app/pkg/repo"
	s "app/pkg/service"
	v "app/pkg/validation"

	"database/sql"
	"log"
	"sync"
	"time"
)

type Valid struct {
	Auth                 v.AuthValidator
	Author               v.AuthorValidator
	Permission           v.PermissionValidator
	PermissionManagement v.PermissionManagementValidator
	Post                 v.PostValidator
	Role                 v.RoleValidator
	RoleManagement       v.RoleManagementValidator
	User                 v.UserValidator
	Book                 v.BookValidator
}

type Auth struct {
	Auth                 a.AuthAuthorization
	Author               a.AuthorAuthorization
	Permission           a.PermissionAuthorization
	PermissionManagement a.PermissionManagementAuthorization
	Post                 a.PostAuthorization
	Role                 a.RoleAuthorization
	RoleManagement       a.RoleManagementAuthorization
	User                 a.UserAuthorization
	Book                 a.BookAuthorization
}

type DI struct {
	DB    *sql.DB
	Jwt   s.JWTService
	Valid Valid
	Auth  Auth
}

func Di() DI {
	db, err := sql.Open("sqlite3", "small.db?_journal_mode=WAL")
	if err != nil {
		log.Fatal(err)
	}

	signingSecret := "my secret"
	jwt := s.NewJWTService(
		&signingSecret,
		time.Duration(20*time.Second),
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
	rmRepo := r.NewRoleManagementRepo(db, rw, dbU, uRepo, rRepo, peRepo)
	pmRepo := r.NewPermissionManagementRepo(db, rw, dbU, uRepo, rRepo, peRepo)

	// Encrypt
	en := s.NewEncryptionService()
	aEncrypt := e.NewAuthEncryption(aRepo, en)
	uEncrypt := e.NewUserEncryption(uRepo, en)

	// Validators
	uVal := v.NewUserValidator(uEncrypt)
	aVal := v.NewAuthValidator(aEncrypt)
	rVal := v.NewRoleValidator(rRepo)
	rmVal := v.NewRoleManagementValidator(rmRepo)
	peVal := v.NewPermissionValidator(peRepo)
	auVal := v.NewAuthorValidator(auRepo)
	bVal := v.NewBookValidator(bRepo)
	pVal := v.NewPostValidator(pRepo)
	pmVal := v.NewPermissionManagementValidator(pmRepo)

	// Authorization
	auth := s.NewAuthorizationService(pmRepo)
	pAuth := a.NewPostAuthorization(auth, pVal)
	bAuth := a.NewBookAuthorization(auth, bVal)
	peAuth := a.NewPermissionAuthorization(auth, peVal)
	pmAuth := a.NewPermissionManagementAuthorization(auth, pmVal)
	aAuth := a.NewAuthAuthorization(auth, aVal)
	auAuth := a.NewAuthorAuthorization(auth, auVal)
	rAuth := a.NewRoleAuthorization(auth, rVal)
	rmAuth := a.NewRoleManagementAuthorization(auth, rmRepo)
	uAuth := a.NewUserAuthorization(auth, uVal)

	return DI{
		DB:    db,
		Jwt:   jwt,
		Valid: Valid{aVal, auVal, peVal, pmVal, pVal, rVal, rmVal, uVal, bVal},
		Auth:  Auth{aAuth, auAuth, peAuth, pmAuth, pAuth, rAuth, rmAuth, uAuth, bAuth},
	}
}
