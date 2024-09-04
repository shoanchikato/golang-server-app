package di

import (
	a "app/pkg/authorization"
	e "app/pkg/encrypt"
	h "app/pkg/handler"
	r "app/pkg/repo"
	rt "app/pkg/route"
	s "app/pkg/service"
	v "app/pkg/validation"

	"database/sql"
	"log"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

type DI struct {
	App *fiber.App
	DB  *sql.DB
}

func Di() DI {
	db, err := sql.Open("sqlite3", "small.db?_journal_mode=WAL")
	if err != nil {
		log.Fatal(err)
	}

	signingSecret := "my secret"
	jwt := s.NewJWTService(
		&signingSecret,
		time.Duration(20*time.Minute),
		time.Duration(1*time.Hour),
	)

	rw := &sync.RWMutex{}
	dbU := r.NewDBUtil(db, rw)

	// Repos
	repos := r.RepoDi(db, rw, dbU)

	// Encrypt
	en := s.NewEncryptionService()
	encryptions := e.EncryptDi(en, repos)

	// Validators
	validators := v.ValidationDi(repos, encryptions)

	// Authorization
	auth := s.NewAuthorizationService(repos.PermissionManagement)
	authorizations := a.AuthorizationDi(auth, validators)

	// Handlers
	handlers := h.HandlerDi(authorizations, jwt)

	// Middleware
	authMiddleware := h.NewAuthMiddleware(jwt)

	// Fiber app
	app := fiber.New()

	// Routes
	rt.Routes(app, handlers, authMiddleware)

	return DI{
		App: app,
		DB:  db,
	}
}
