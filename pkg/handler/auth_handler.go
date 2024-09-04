package handler

import (
	aa "app/pkg/authorization"
	m "app/pkg/model"
	s "app/pkg/service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler interface {
	Login(c *fiber.Ctx) error
	ResetPassword(c *fiber.Ctx) error
}

type authHandler struct {
	auth aa.AuthAuthorization
	jwt  s.JWTService
}

func NewAuthHandler(auth aa.AuthAuthorization, jwt s.JWTService) AuthHandler {
	return &authHandler{auth, jwt}
}

// Login
func (a *authHandler) Login(c *fiber.Ctx) error {
	credentials := m.Credentials{}

	err := c.BodyParser(&credentials)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	userId, err := a.auth.Login(&credentials)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	tokens, err := a.jwt.GetTokens(*userId)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	return c.JSON(tokens)
}

// ResetPassword
func (a *authHandler) ResetPassword(c *fiber.Ctx) error {
	credentials := m.Credentials{}

	err := c.BodyParser(&credentials)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	err = a.auth.ResetPassword(credentials.Username, credentials.Password)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	return c.SendStatus(http.StatusCreated)
}
