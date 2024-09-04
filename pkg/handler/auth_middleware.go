package handler

import (
	s "app/pkg/service"
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type AuthMiddleware interface {
	JWTParser(c *fiber.Ctx) error
}

type authMiddleware struct {
	auth s.JWTService
}

type userContextKey string

func NewAuthMiddleware(jwt s.JWTService) AuthMiddleware {
	return &authMiddleware{jwt}
}

// JWTParser
func (a *authMiddleware) JWTParser(c *fiber.Ctx) error {
	value := c.Get("Authorization")
	if value == "" {
		return c.SendStatus(http.StatusUnauthorized)
	}

	tokenStr := value
	token, err := a.auth.ParseToken(&tokenStr)
	if err != nil {
		return c.Status(http.StatusUnauthorized).SendString(err.Error())
	}

	userIDKey := userContextKey("userId")

	ctx := context.WithValue(context.Background(), userIDKey, token.UserId)

	c.SetUserContext(ctx)

	return c.Next()
}
