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

type UserHandler interface {
	Add(c *fiber.Ctx) error
	AddAll(c *fiber.Ctx) error
	Edit(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	GetOne(c *fiber.Ctx) error
	Remove(c *fiber.Ctx) error
}

type userHandler struct {
	service ef.UserHttpErrorFmt
	logger  s.Logger
}

func NewUserHandler(service ef.UserHttpErrorFmt, logger s.Logger) UserHandler {
	return &userHandler{service, logger}
}

// Add
func (u *userHandler) Add(c *fiber.Ctx) error {
	user := m.User{}

	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	userId, err := getAuthUserId(c, u.logger)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	err = u.service.Add(*userId, &user)

	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).JSON(httpErr)
	}

	return c.Status(http.StatusCreated).JSON(user)
}

// AddAll
func (u *userHandler) AddAll(c *fiber.Ctx) error {
	users := []m.User{}

	err := c.BodyParser(&users)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	userId, err := getAuthUserId(c, u.logger)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	newUsers := []*m.User{}
	for i := range users {
		newUsers = append(newUsers, &users[i])
	}

	err = u.service.AddAll(*userId, &newUsers)

	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).JSON(httpErr)
	}

	return c.Status(http.StatusCreated).JSON(newUsers)
}

// Edit
func (u *userHandler) Edit(c *fiber.Ctx) error {
	user := m.EditUser{}

	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	id, err := getIntParam(c, u.logger, "id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	userId, err := getAuthUserId(c, u.logger)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	err = u.service.Edit(*userId, *id, &user)

	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).JSON(httpErr)
	}

	return c.Status(http.StatusCreated).JSON(user)
}

// GetAll
func (u *userHandler) GetAll(c *fiber.Ctx) error {
	userId, err := getAuthUserId(c, u.logger)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	users, err := u.service.GetAll(*userId, 0, 50)

	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).JSON(httpErr)
	}

	return c.Status(http.StatusOK).JSON(users)
}

// GetOne
func (u *userHandler) GetOne(c *fiber.Ctx) error {
	userId, err := getAuthUserId(c, u.logger)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	id, err := getIntParam(c, u.logger, "id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	user, err := u.service.GetOne(*userId, *id)

	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).JSON(httpErr)
	}

	return c.Status(http.StatusOK).JSON(user)
}

// Remove
func (u *userHandler) Remove(c *fiber.Ctx) error {
	userId, err := getAuthUserId(c, u.logger)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	id, err := getIntParam(c, u.logger, "id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	err = u.service.Remove(*userId, *id)

	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).JSON(httpErr)
	}

	return c.SendStatus(http.StatusNoContent)
}
