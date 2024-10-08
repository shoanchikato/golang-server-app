package handler

import (
	e "app/pkg/errors"
	ef "app/pkg/httperrorfmt"
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
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	tokens, err := a.service.Login(&credentials)

	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).JSON(httpErr)
	}

	return c.JSON(tokens)
}

// ResetPassword
func (a *authHandler) ResetPassword(c *fiber.Ctx) error {
	credentials := m.Credentials{}

	err := c.BodyParser(&credentials)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	err = a.service.ResetPassword(credentials.Username, credentials.Password)

	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).JSON(httpErr)
	}

	return c.SendStatus(http.StatusCreated)
}
