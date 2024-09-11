package middleware

import (
	e "app/pkg/errors"
	s "app/pkg/service"
	"context"
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type AuthMiddleware interface {
	JWTParser(c *fiber.Ctx) error
}

type authMiddleware struct {
	auth         s.JWTService
	httpErrorFmt s.HttpErrorFmt
	logger       s.Logger
}

type UserContextKey string

func NewAuthMiddleware(jwt s.JWTService, httpErrorFmt s.HttpErrorFmt, logger s.Logger) AuthMiddleware {
	return &authMiddleware{jwt, httpErrorFmt, logger}
}

// JWTParser
func (a *authMiddleware) JWTParser(c *fiber.Ctx) error {
	value := c.Get("Authorization")
	if value == "" {
		// TODO: assign to default userId with default set permissions
		a.logger.Error(e.ErrMissingAuthorizationHeader.Error())
		return c.Status(http.StatusUnauthorized).JSON(e.NewHttpErrorMap(e.ErrMissingAuthorizationHeader))
	}

	tokenStr := value
	token, err := a.auth.ParseToken(&tokenStr)

	httpErr := &e.HttpError{}
	if err = a.httpErrorFmt.GetError(err); errors.As(err, &httpErr) {
		a.logger.Error(httpErr.Error())
		return c.Status(httpErr.HTTPStatus).JSON(httpErr)
	}

	userIDKey := UserContextKey("userId")

	ctx := context.WithValue(context.Background(), userIDKey, token.UserId)

	c.SetUserContext(ctx)

	return c.Next()
}
