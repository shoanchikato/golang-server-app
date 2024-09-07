package handler

import (
	e "app/pkg/errors"
	ef "app/pkg/httperrorfmt"
	m "app/pkg/model"
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type PermissionHandler interface {
	Add(c *fiber.Ctx) error
	AddAll(c *fiber.Ctx) error
	Edit(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	GetOne(c *fiber.Ctx) error
	Remove(c *fiber.Ctx) error
}

type permissionHandler struct {
	service ef.PermissionHttpErrorFmt
}

func NewPermissionHandler(service ef.PermissionHttpErrorFmt) PermissionHandler {
	return &permissionHandler{service}
}

// Add
func (p *permissionHandler) Add(c *fiber.Ctx) error {
	permission := m.Permission{}

	err := c.BodyParser(&permission)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	userId, err := getAuthUserId(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	err = p.service.Add(*userId, &permission)

	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).SendString(httpErr.Message)
	}

	return c.Status(http.StatusCreated).JSON(permission)
}

// AddAll
func (p *permissionHandler) AddAll(c *fiber.Ctx) error {
	permissions := []m.Permission{}

	err := c.BodyParser(&permissions)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	userId, err := getAuthUserId(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	newPermissions := []*m.Permission{}
	for i := range permissions {
		newPermissions[i] = &permissions[i]
	}

	err = p.service.AddAll(*userId, &newPermissions)

	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).SendString(httpErr.Message)
	}

	return c.Status(http.StatusCreated).JSON(newPermissions)
}

// Edit
func (p *permissionHandler) Edit(c *fiber.Ctx) error {
	permission := m.Permission{}

	err := c.BodyParser(&permission)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	id, err := getId(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(e.ErrProvideNumericId.Error())
	}

	userId, err := getAuthUserId(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	err = p.service.Edit(*userId, *id, &permission)

	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).SendString(httpErr.Message)
	}

	return c.Status(http.StatusCreated).JSON(permission)
}

// GetAll
func (p *permissionHandler) GetAll(c *fiber.Ctx) error {
	userId, err := getAuthUserId(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	users, err := p.service.GetAll(*userId, 0, 50)

	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).SendString(httpErr.Message)
	}

	return c.Status(http.StatusCreated).JSON(users)
}

// GetOne
func (p *permissionHandler) GetOne(c *fiber.Ctx) error {
	userId, err := getAuthUserId(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	id, err := getId(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(e.ErrProvideNumericId.Error())
	}

	user, err := p.service.GetOne(*userId, *id)

	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).SendString(httpErr.Message)
	}

	return c.Status(http.StatusCreated).JSON(user)
}

// Remove
func (p *permissionHandler) Remove(c *fiber.Ctx) error {
	userId, err := getAuthUserId(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	id, err := getId(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(e.ErrProvideNumericId.Error())
	}

	err = p.service.Remove(*userId, *id)

	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).SendString(httpErr.Message)
	}

	return c.SendStatus(http.StatusAccepted)
}
