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
	logger  s.Logger
}

func NewAuthorHandler(service ef.AuthorHttpErrorFmt, logger s.Logger) AuthorHandler {
	return &authorHandler{service, logger}
}

// Add godoc
//
//	@Description	add an author
//	@Tags			Authors
//	@Accept			json
//	@Security		BearerAuth
//	@Param			author	body	model.Author	true	"Author Details"
//	@Produce		json
//	@Success		201	{string}	created
//	@Failure		400	{object}	errors.HttpErrorMap
//	@Failure		401	{object}	errors.HttpErrorMap
//	@Failure		404	{object}	errors.HttpErrorMap
//	@Failure		500	{object}	errors.HttpErrorMap
//	@Router			/authors [post]
func (p *authorHandler) Add(c *fiber.Ctx) error {
	author := m.Author{}

	err := c.BodyParser(&author)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	userId, err := getAuthUserId(c, p.logger)
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

// Add All godoc
//
//	@Description	add an array of authors
//	@Tags			Authors
//	@Accept			json
//	@Security		BearerAuth
//	@Param			authors	body	[]model.Author	true	"Authors' Details"
//	@Produce		json
//	@Success		201	{string}	created
//	@Failure		400	{object}	errors.HttpErrorMap
//	@Failure		401	{object}	errors.HttpErrorMap
//	@Failure		404	{object}	errors.HttpErrorMap
//	@Failure		500	{object}	errors.HttpErrorMap
//	@Router			/authors/all [post]
func (p *authorHandler) AddAll(c *fiber.Ctx) error {
	authors := []m.Author{}

	err := c.BodyParser(&authors)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	userId, err := getAuthUserId(c, p.logger)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	newAuthors := []*m.Author{}
	for i := range authors {
		newAuthors = append(newAuthors, &authors[i])
	}

	err = p.service.AddAll(*userId, &newAuthors)

	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).JSON(httpErr)
	}

	return c.Status(http.StatusCreated).JSON(newAuthors)
}

// Edit godoc
//
//	@Description	edit an author
//	@Tags			Authors
//	@Accept			json
//	@Security		BearerAuth
//	@Param			author	body	model.Author	true	"Author Details"
//	@Param			id		path	int				true	"Author Id"
//	@Produce		json
//	@Success		201	{string}	created
//	@Failure		400	{object}	errors.HttpErrorMap
//	@Failure		401	{object}	errors.HttpErrorMap
//	@Failure		404	{object}	errors.HttpErrorMap
//	@Failure		500	{object}	errors.HttpErrorMap
//	@Router			/authors/{id} [put]
func (p *authorHandler) Edit(c *fiber.Ctx) error {
	author := m.Author{}

	err := c.BodyParser(&author)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	id, err := getIntParam(c, p.logger, "id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	userId, err := getAuthUserId(c, p.logger)
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

// Get All godoc
//
//	@Description	get all authors
//	@Tags			Authors
//	@Accept			json
//	@Security		BearerAuth
//	@Produce		json
//	@Success		200	{object}	[]model.Author
//	@Failure		400	{object}	errors.HttpErrorMap
//	@Failure		401	{object}	errors.HttpErrorMap
//	@Failure		404	{object}	errors.HttpErrorMap
//	@Failure		500	{object}	errors.HttpErrorMap
//	@Router			/authors [get]
func (p *authorHandler) GetAll(c *fiber.Ctx) error {
	userId, err := getAuthUserId(c, p.logger)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	users, err := p.service.GetAll(*userId, 0, 50)

	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).JSON(httpErr)
	}

	return c.Status(http.StatusOK).JSON(users)
}

// Get One godoc
//
//	@Description	get one author
//	@Tags			Authors
//	@Accept			json
//	@Security		BearerAuth
//	@Param			id	path	int	true	"Author Id"
//	@Produce		json
//	@Success		200	{object}	model.Author
//	@Failure		400	{object}	errors.HttpErrorMap
//	@Failure		401	{object}	errors.HttpErrorMap
//	@Failure		404	{object}	errors.HttpErrorMap
//	@Failure		500	{object}	errors.HttpErrorMap
//	@Router			/authors/{id} [get]
func (p *authorHandler) GetOne(c *fiber.Ctx) error {
	userId, err := getAuthUserId(c, p.logger)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	id, err := getIntParam(c, p.logger, "id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	user, err := p.service.GetOne(*userId, *id)

	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).JSON(httpErr)
	}

	return c.Status(http.StatusOK).JSON(user)
}

// Remove Author godoc
//
//	@Description	delete an author
//	@Tags			Authors
//	@Accept			json
//	@Security		BearerAuth
//	@Param			id	path	int	true	"Author Id"
//	@Produce		json
//	@Success		204	
//	@Failure		400	{object}	errors.HttpErrorMap
//	@Failure		401	{object}	errors.HttpErrorMap
//	@Failure		404	{object}	errors.HttpErrorMap
//	@Failure		500	{object}	errors.HttpErrorMap
//	@Router			/authors/{id} [delete]
func (p *authorHandler) Remove(c *fiber.Ctx) error {
	userId, err := getAuthUserId(c, p.logger)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	id, err := getIntParam(c, p.logger, "id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	err = p.service.Remove(*userId, *id)

	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).JSON(httpErr)
	}

	return c.SendStatus(http.StatusNoContent)
}
