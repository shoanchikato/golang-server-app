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
	logger  s.Logger
}

func NewRoleHandler(service ef.RoleHttpErrorFmt, logger s.Logger) RoleHandler {
	return &roleHandler{service, logger}
}

// Add godoc
//
//	@Description	add a role
//	@Tags			Roles
//	@Accept			json
//	@Security		BearerAuth
//	@Param			role	body	model.Role	true	"Role Details"
//	@Produce		json
//	@Success		201	{string}	created
//	@Failure		400	{object}	errors.HttpErrorMap
//	@Failure		401	{object}	errors.HttpErrorMap
//	@Failure		404	{object}	errors.HttpErrorMap
//	@Failure		500	{object}	errors.HttpErrorMap
//	@Router			/roles [post]
func (p *roleHandler) Add(c *fiber.Ctx) error {
	role := m.Role{}

	err := c.BodyParser(&role)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	userId, err := getAuthUserId(c, p.logger)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	err = p.service.Add(*userId, &role)

	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).JSON(httpErr)
	}

	return c.Status(http.StatusCreated).JSON(role)
}

// Add All godoc
//
//	@Description	add an array of roles
//	@Tags			Roles
//	@Accept			json
//	@Security		BearerAuth
//	@Param			role	body	[]model.Role	true	"Roles' Details"
//	@Produce		json
//	@Success		201	{string}	created
//	@Failure		400	{object}	errors.HttpErrorMap
//	@Failure		401	{object}	errors.HttpErrorMap
//	@Failure		404	{object}	errors.HttpErrorMap
//	@Failure		500	{object}	errors.HttpErrorMap
//	@Router			/roles/all [post]
func (p *roleHandler) AddAll(c *fiber.Ctx) error {
	roles := []m.Role{}

	err := c.BodyParser(&roles)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	userId, err := getAuthUserId(c, p.logger)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	newRoles := []*m.Role{}
	for i := range roles {
		newRoles = append(newRoles, &roles[i])
	}

	err = p.service.AddAll(*userId, &newRoles)

	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).JSON(httpErr)
	}

	return c.Status(http.StatusCreated).JSON(newRoles)
}

// Edit godoc
//
//	@Description	edit a role
//	@Tags			Roles
//	@Accept			json
//	@Security		BearerAuth
//	@Param			role	body	model.Role	true	"Role Details"
//	@Param			id		path	int			true	"Role Id"
//	@Produce		json
//	@Success		201	{string}	created
//	@Failure		400	{object}	errors.HttpErrorMap
//	@Failure		401	{object}	errors.HttpErrorMap
//	@Failure		404	{object}	errors.HttpErrorMap
//	@Failure		500	{object}	errors.HttpErrorMap
//	@Router			/roles/{id} [put]
func (p *roleHandler) Edit(c *fiber.Ctx) error {
	role := m.Role{}

	err := c.BodyParser(&role)
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

	err = p.service.Edit(*userId, *id, &role)

	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).JSON(httpErr)
	}

	return c.Status(http.StatusCreated).JSON(role)
}

// Get All godoc
//
//	@Description	get all roles
//	@Tags			Roles
//	@Accept			json
//	@Security		BearerAuth
//	@Produce		json
//	@Success		200	{object}	[]model.Role
//	@Failure		400	{object}	errors.HttpErrorMap
//	@Failure		401	{object}	errors.HttpErrorMap
//	@Failure		404	{object}	errors.HttpErrorMap
//	@Failure		500	{object}	errors.HttpErrorMap
//	@Router			/roles [get]
func (p *roleHandler) GetAll(c *fiber.Ctx) error {
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
//	@Description	get one role
//	@Tags			Roles
//	@Accept			json
//	@Security		BearerAuth
//	@Param			id	path	int	true	"Role Id"
//	@Produce		json
//	@Success		200	{object}	model.Role
//	@Failure		400	{object}	errors.HttpErrorMap
//	@Failure		401	{object}	errors.HttpErrorMap
//	@Failure		404	{object}	errors.HttpErrorMap
//	@Failure		500	{object}	errors.HttpErrorMap
//	@Router			/roles/{id} [get]
func (p *roleHandler) GetOne(c *fiber.Ctx) error {
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

// Remove Role godoc
//
//	@Description	delete a role
//	@Tags			Roles
//	@Accept			json
//	@Security		BearerAuth
//	@Param			id	path	int	true	"Role Id"
//	@Produce		json
//	@Success		204	
//	@Failure		400	{object}	errors.HttpErrorMap
//	@Failure		401	{object}	errors.HttpErrorMap
//	@Failure		404	{object}	errors.HttpErrorMap
//	@Failure		500	{object}	errors.HttpErrorMap
//	@Router			/roles/{id} [delete]
func (p *roleHandler) Remove(c *fiber.Ctx) error {
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
