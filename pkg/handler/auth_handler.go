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

// Login godoc
//	@Description	login using username and password
//	@Tags			Auth
//	@Accept			json
//	@Param			credentials	body	model.Credentials	true	"User Credentials"
//	@Produce		json
//	@Success		200	{object}	model.Tokens		"Access and Refresh Tokens"
//	@Failure		400	{object}	errors.HttpErrorMap	
//	@Failure		401	{object}	errors.HttpErrorMap
//	@Failure		404	{object}	errors.HttpErrorMap
//	@Failure		500	{object}	errors.HttpErrorMap
//	@Router			/login [post]
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

// ResetPassword godoc
//	@Description	resetting password
//	@Tags			Auth
//	@Accept			json
//	@Param			credentials	body	model.Credentials	true	"User Credentials"
//  @Param Authorization header string true "Bearer token"
//	@Produce		json
//	@Success		201	{string}	created
//	@Failure		400	{object}	errors.HttpErrorMap	
//	@Failure		401	{object}	errors.HttpErrorMap
//	@Failure		404	{object}	errors.HttpErrorMap
//	@Failure		500	{object}	errors.HttpErrorMap
// @securityDefinitions.bearer  BearerAuth
//	@Router			/reset-password [post]
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
