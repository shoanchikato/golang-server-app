package handler

import (
	ef "app/pkg/errorfmt"
	e "app/pkg/errors"
	m "app/pkg/model"
	s "app/pkg/service"
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler interface {
	Login(c *fiber.Ctx) error
	ResetPassword(c *fiber.Ctx) error
}

type authHandler struct {
	service ef.AuthHttpErrorFmt
	jwt     s.JWTService
}

func NewAuthHandler(service ef.AuthHttpErrorFmt, jwt s.JWTService) AuthHandler {
	return &authHandler{service, jwt}
}

// Login
func (a *authHandler) Login(c *fiber.Ctx) error {
	credentials := m.Credentials{}

	err := c.BodyParser(&credentials)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	tokens, err := a.service.Login(&credentials)

	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).SendString(httpErr.Message)
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

	err = a.service.ResetPassword(credentials.Username, credentials.Password)

	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).SendString(httpErr.Message)
	}

	return c.SendStatus(http.StatusCreated)
}
