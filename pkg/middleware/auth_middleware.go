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
}

type UserContextKey string

func NewAuthMiddleware(jwt s.JWTService, httpErrorFmt s.HttpErrorFmt) AuthMiddleware {
	return &authMiddleware{jwt, httpErrorFmt}
}

// JWTParser
func (a *authMiddleware) JWTParser(c *fiber.Ctx) error {
	value := c.Get("Authorization")
	if value == "" {
		// TODO: assign to default userId with default set permissions
		return c.SendStatus(http.StatusUnauthorized)
	}

	tokenStr := value
	token, err := a.auth.ParseToken(&tokenStr)

	httpErr := &e.HttpError{}
	if err = a.httpErrorFmt.GetError(err); errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).JSON(httpErr)
	}

	userIDKey := UserContextKey("userId")

	ctx := context.WithValue(context.Background(), userIDKey, token.UserId)

	c.SetUserContext(ctx)

	return c.Next()
}
