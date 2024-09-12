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
	"io"
	"log/slog"

	"database/sql"
	"log"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

type DI struct {
	App           *fiber.App
	DB            *sql.DB
	Validators    v.Validators
	HttpErrorFmts *ef.HttpErrorFmts
}

func Di(
	dbDriver string,
	dbConn string,
	secret string,
	logWriter io.Writer,
) DI {
	db, err := sql.Open(dbDriver, dbConn)
	if err != nil {
		log.Fatal(err)
	}

	signingSecret := secret
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

	// logger
	jsonHandler := slog.NewJSONHandler(logWriter, nil)
	slogLogger := slog.New(jsonHandler)
	logger := s.NewLogger(slogLogger)

	// HttpErrorFmt
	httpErrorFmt := s.NewHttpErrorFmt(logger)
	httpErrorFmts := ef.HttpErrorFmtDi(httpErrorFmt, jwt, authorizations)

	// Handlers
	handlers := h.HandlerDi(httpErrorFmts, jwt, logger)

	// Middleware
	authMiddleware := mi.NewAuthMiddleware(jwt, httpErrorFmt, logger)

	// Fiber app
	app := fiber.New()

	// Routes
	rt.Routes(app, handlers, authMiddleware)

	return DI{
		App:           app,
		DB:            db,
		Validators: *validators,
		HttpErrorFmts: httpErrorFmts,
	}
}
