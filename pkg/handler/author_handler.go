package handler

import (
	e "app/pkg/errors"
	ef "app/pkg/httperrorfmt"
	m "app/pkg/model"
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type AuthorHandler interface {
	Add(c *fiber.Ctx) error
	AddAll(c *fiber.Ctx) error
	Edit(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	GetOne(c *fiber.Ctx) error
	Remove(c *fiber.Ctx) error
}

type authorHandler struct {
	service ef.AuthorHttpErrorFmt
}

func NewAuthorHandler(service ef.AuthorHttpErrorFmt) AuthorHandler {
	return &authorHandler{service}
}

// Add
func (p *authorHandler) Add(c *fiber.Ctx) error {
	author := m.Author{}

	err := c.BodyParser(&author)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	userId, err := getAuthUserId(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	err = p.service.Add(*userId, &author)

	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).JSON(httpErr)
	}

	return c.Status(http.StatusCreated).JSON(author)
}

// AddAll
func (p *authorHandler) AddAll(c *fiber.Ctx) error {
	authors := []m.Author{}

	err := c.BodyParser(&authors)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	userId, err := getAuthUserId(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	newAuthors := []*m.Author{}
	for i := range authors {
		newAuthors[i] = &authors[i]
	}

	err = p.service.AddAll(*userId, &newAuthors)

	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).JSON(httpErr)
	}

	return c.Status(http.StatusCreated).JSON(newAuthors)
}

// Edit
func (p *authorHandler) Edit(c *fiber.Ctx) error {
	author := m.Author{}

	err := c.BodyParser(&author)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	id, err := getId(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(e.ErrProvideNumericId))
	}

	userId, err := getAuthUserId(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	err = p.service.Edit(*userId, *id, &author)

	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).JSON(httpErr)
	}

	return c.Status(http.StatusCreated).JSON(author)
}

// GetAll
func (p *authorHandler) GetAll(c *fiber.Ctx) error {
	userId, err := getAuthUserId(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	users, err := p.service.GetAll(*userId, 0, 50)

	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).JSON(httpErr)
	}

	return c.Status(http.StatusCreated).JSON(users)
}

// GetOne
func (p *authorHandler) GetOne(c *fiber.Ctx) error {
	userId, err := getAuthUserId(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	id, err := getId(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(e.ErrProvideNumericId))
	}

	user, err := p.service.GetOne(*userId, *id)

	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).JSON(httpErr)
	}

	return c.Status(http.StatusCreated).JSON(user)
}

// Remove
func (p *authorHandler) Remove(c *fiber.Ctx) error {
	userId, err := getAuthUserId(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	id, err := getId(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(e.ErrProvideNumericId))
	}

	err = p.service.Remove(*userId, *id)

	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).JSON(httpErr)
	}

	return c.SendStatus(http.StatusAccepted)
}
