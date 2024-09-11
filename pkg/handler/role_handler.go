package handler

import (
	e "app/pkg/errors"
	ef "app/pkg/httperrorfmt"
	m "app/pkg/model"
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type RoleHandler interface {
	Add(c *fiber.Ctx) error
	AddAll(c *fiber.Ctx) error
	Edit(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	GetOne(c *fiber.Ctx) error
	Remove(c *fiber.Ctx) error
}

type roleHandler struct {
	service ef.RoleHttpErrorFmt
}

func NewRoleHandler(service ef.RoleHttpErrorFmt) RoleHandler {
	return &roleHandler{service}
}

// Add
func (p *roleHandler) Add(c *fiber.Ctx) error {
	role := m.Role{}

	err := c.BodyParser(&role)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err.Error()))
	}

	userId, err := getAuthUserId(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err.Error()))
	}

	err = p.service.Add(*userId, &role)

	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).JSON(httpErr)
	}

	return c.Status(http.StatusCreated).JSON(role)
}

// AddAll
func (p *roleHandler) AddAll(c *fiber.Ctx) error {
	roles := []m.Role{}

	err := c.BodyParser(&roles)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err.Error()))
	}

	userId, err := getAuthUserId(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err.Error()))
	}

	newRoles := []*m.Role{}
	for i := range roles {
		newRoles[i] = &roles[i]
	}

	err = p.service.AddAll(*userId, &newRoles)

	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).JSON(httpErr)
	}

	return c.Status(http.StatusCreated).JSON(newRoles)
}

// Edit
func (p *roleHandler) Edit(c *fiber.Ctx) error {
	role := m.Role{}

	err := c.BodyParser(&role)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err.Error()))
	}

	id, err := getId(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(e.ErrProvideNumericId.Error()))
	}

	userId, err := getAuthUserId(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err.Error()))
	}

	err = p.service.Edit(*userId, *id, &role)

	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).JSON(httpErr)
	}

	return c.Status(http.StatusCreated).JSON(role)
}

// GetAll
func (p *roleHandler) GetAll(c *fiber.Ctx) error {
	userId, err := getAuthUserId(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err.Error()))
	}

	users, err := p.service.GetAll(*userId, 0, 50)

	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).JSON(httpErr)
	}

	return c.Status(http.StatusCreated).JSON(users)
}

// GetOne
func (p *roleHandler) GetOne(c *fiber.Ctx) error {
	userId, err := getAuthUserId(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err.Error()))
	}

	id, err := getId(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(e.ErrProvideNumericId.Error()))
	}

	user, err := p.service.GetOne(*userId, *id)

	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).JSON(httpErr)
	}

	return c.Status(http.StatusCreated).JSON(user)
}

// Remove
func (p *roleHandler) Remove(c *fiber.Ctx) error {
	userId, err := getAuthUserId(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err.Error()))
	}

	id, err := getId(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(e.ErrProvideNumericId.Error()))
	}

	err = p.service.Remove(*userId, *id)

	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).JSON(httpErr)
	}

	return c.SendStatus(http.StatusAccepted)
}
