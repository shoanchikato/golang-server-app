package di

import (
	a "app/pkg/authorization"
	e "app/pkg/encrypt"
	h "app/pkg/handler"
	ef "app/pkg/httperrorfmt"
	mi "app/pkg/middleware"
	r "app/pkg/repo"
	rt "app/pkg/route"
	s "app/pkg/service"
	v "app/pkg/validation"

	"database/sql"
	"log"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type DI struct {
	App *fiber.App
	DB  *sql.DB
}

func Di() DI {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	env, err := godotenv.Read()
	if err != nil {
		log.Fatal("error loading env map")
	}

	db, err := sql.Open(env["DB_DRIVER"], env["DB_CONN"])
	if err != nil {
		log.Fatal(err)
	}

	signingSecret := env["SECRET"]
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
	validation := s.NewValidationService()
	validators := v.ValidationDi(repos, encryptions, validation)

	// Authorization
	auth := s.NewAuthorizationService(repos.PermissionManagement)
	authorizations := a.AuthorizationDi(auth, validators)

	// HttpErrorFmt
	httpErrorFmt := s.NewHttpErrorFmt()
	httpErrorFmts := ef.HttpErrorFmtDi(httpErrorFmt, jwt, authorizations)

	// Handlers
	handlers := h.HandlerDi(httpErrorFmts, jwt)

	// Middleware
	authMiddleware := mi.NewAuthMiddleware(jwt, httpErrorFmt)

	// Fiber app
	app := fiber.New()

	// Routes
	rt.Routes(app, handlers, authMiddleware)

	return DI{
		App: app,
		DB:  db,
	}
}
