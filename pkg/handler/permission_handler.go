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
	logger  s.Logger
}

func NewPermissionHandler(service ef.PermissionHttpErrorFmt, logger s.Logger) PermissionHandler {
	return &permissionHandler{service, logger}
}

// Add godoc
//
//	@Description	add a permission
//	@Tags			Permissions
//	@Accept			json
//	@Security		BearerAuth
//	@Param			permission	body	model.Permission	true	"Permission Details"
//	@Produce		json
//	@Success		201	{string}	created
//	@Failure		400	{object}	errors.HttpErrorMap
//	@Failure		401	{object}	errors.HttpErrorMap
//	@Failure		404	{object}	errors.HttpErrorMap
//	@Failure		500	{object}	errors.HttpErrorMap
//	@Router			/permissions [post]
func (p *permissionHandler) Add(c *fiber.Ctx) error {
	permission := m.Permission{}

	err := c.BodyParser(&permission)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	userId, err := getAuthUserId(c, p.logger)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	err = p.service.Add(*userId, &permission)

	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).JSON(httpErr)
	}

	return c.Status(http.StatusCreated).JSON(permission)
}

// Add All godoc
//
//	@Description	add an array of permissions
//	@Tags			Permissions
//	@Accept			json
//	@Security		BearerAuth
//	@Param			permission	body	[]model.Permission	true	"Permissions' Details"
//	@Produce		json
//	@Success		201	{string}	created
//	@Failure		400	{object}	errors.HttpErrorMap
//	@Failure		401	{object}	errors.HttpErrorMap
//	@Failure		404	{object}	errors.HttpErrorMap
//	@Failure		500	{object}	errors.HttpErrorMap
//	@Router			/permissions/all [post]
func (p *permissionHandler) AddAll(c *fiber.Ctx) error {
	permissions := []m.Permission{}

	err := c.BodyParser(&permissions)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	userId, err := getAuthUserId(c, p.logger)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	newPermissions := []*m.Permission{}
	for i := range permissions {
		newPermissions = append(newPermissions, &permissions[i])
	}

	err = p.service.AddAll(*userId, &newPermissions)

	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).JSON(httpErr)
	}

	return c.Status(http.StatusCreated).JSON(newPermissions)
}

// Edit godoc
//
//	@Description	edit a permission
//	@Tags			Permissions
//	@Accept			json
//	@Security		BearerAuth
//	@Param			permission	body	model.Permission	true	"Permission Details"
//	@Param			id			path	int					true	"Permission Id"
//	@Produce		json
//	@Success		201	{string}	created
//	@Failure		400	{object}	errors.HttpErrorMap
//	@Failure		401	{object}	errors.HttpErrorMap
//	@Failure		404	{object}	errors.HttpErrorMap
//	@Failure		500	{object}	errors.HttpErrorMap
//	@Router			/permissions/{id} [put]
func (p *permissionHandler) Edit(c *fiber.Ctx) error {
	permission := m.Permission{}

	err := c.BodyParser(&permission)
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

	err = p.service.Edit(*userId, *id, &permission)

	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).JSON(httpErr)
	}

	return c.Status(http.StatusCreated).JSON(permission)
}

// Get All godoc
//
//	@Description	get all permissions
//	@Tags			Permissions
//	@Accept			json
//	@Security		BearerAuth
//	@Produce		json
//	@Success		200	{object}	[]model.Permission
//	@Failure		400	{object}	errors.HttpErrorMap
//	@Failure		401	{object}	errors.HttpErrorMap
//	@Failure		404	{object}	errors.HttpErrorMap
//	@Failure		500	{object}	errors.HttpErrorMap
//	@Router			/permissions [get]
func (p *permissionHandler) GetAll(c *fiber.Ctx) error {
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
//	@Description	get one permission
//	@Tags			Permissions
//	@Accept			json
//	@Security		BearerAuth
//	@Param			id	path	int	true	"Permission Id"
//	@Produce		json
//	@Success		200	{object}	model.Permission
//	@Failure		400	{object}	errors.HttpErrorMap
//	@Failure		401	{object}	errors.HttpErrorMap
//	@Failure		404	{object}	errors.HttpErrorMap
//	@Failure		500	{object}	errors.HttpErrorMap
//	@Router			/permissions/{id} [get]
func (p *permissionHandler) GetOne(c *fiber.Ctx) error {
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

// Remove Permission godoc
//
//	@Description	delete a permission
//	@Tags			Permissions
//	@Accept			json
//	@Security		BearerAuth
//	@Param			id	path	int	true	"Permission Id"
//	@Produce		json
//	@Success		204	
//	@Failure		400	{object}	errors.HttpErrorMap
//	@Failure		401	{object}	errors.HttpErrorMap
//	@Failure		404	{object}	errors.HttpErrorMap
//	@Failure		500	{object}	errors.HttpErrorMap
//	@Router			/permissions/{id} [delete]
func (p *permissionHandler) Remove(c *fiber.Ctx) error {
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
