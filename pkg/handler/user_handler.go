package handler

import (
	ef "app/pkg/errorfmt"
	e "app/pkg/errors"
	m "app/pkg/model"
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
	service ef.UserErrorFmt
}

func NewUserHandler(service ef.UserErrorFmt) UserHandler {
	return &userHandler{service}
}

// Add
func (u *userHandler) Add(c *fiber.Ctx) error {
	user := m.User{}

	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	userId, err := getAuthUserId(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	err = u.service.Add(*userId, &user)

	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).SendString(httpErr.Message)
	}

	return c.Status(http.StatusCreated).JSON(user)
}

// AddAll
func (u *userHandler) AddAll(c *fiber.Ctx) error {
	users := []m.User{}

	err := c.BodyParser(&users)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	userId, err := getAuthUserId(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	newUsers := []*m.User{}
	for i := range users {
		newUsers[i] = &users[i]
	}

	err = u.service.AddAll(*userId, &newUsers)

	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).SendString(httpErr.Message)
	}

	return c.Status(http.StatusCreated).JSON(newUsers)
}

// Edit
func (u *userHandler) Edit(c *fiber.Ctx) error {
	user := m.EditUser{}

	err := c.BodyParser(&user)
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

	err = u.service.Edit(*userId, *id, &user)

	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).SendString(httpErr.Message)
	}

	return c.Status(http.StatusCreated).JSON(user)
}

// GetAll
func (u *userHandler) GetAll(c *fiber.Ctx) error {
	userId, err := getAuthUserId(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	users, err := u.service.GetAll(*userId, 0, 50)

	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).SendString(httpErr.Message)
	}

	return c.Status(http.StatusCreated).JSON(users)
}

// GetOne
func (u *userHandler) GetOne(c *fiber.Ctx) error {
	userId, err := getAuthUserId(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	id, err := getId(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(e.ErrProvideNumericId.Error())
	}

	user, err := u.service.GetOne(*userId, *id)

	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).SendString(httpErr.Message)
	}

	return c.Status(http.StatusCreated).JSON(user)
}

// Remove
func (u *userHandler) Remove(c *fiber.Ctx) error {
	userId, err := getAuthUserId(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	id, err := getId(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(e.ErrProvideNumericId.Error())
	}

	err = u.service.Remove(*userId, *id)

	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).SendString(httpErr.Message)
	}

	return c.SendStatus(http.StatusAccepted)
}
