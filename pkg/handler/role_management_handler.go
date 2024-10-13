package handler

import (
	e "app/pkg/errors"
	ef "app/pkg/httperrorfmt"
	s "app/pkg/service"
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type RoleManagementHandler interface {
	AddRoleToUser(c *fiber.Ctx) error
	GetRolesByUserId(c *fiber.Ctx) error
	RemoveRoleFromUser(c *fiber.Ctx) error
}

type roleManagementHandler struct {
	service ef.RoleManagementHttpErrorFmt
	logger  s.Logger
}

func NewRoleManagementHandler(service ef.RoleManagementHttpErrorFmt, logger s.Logger) RoleManagementHandler {
	return &roleManagementHandler{service, logger}
}

// Add Role To User godoc
//
//	@Description	add a role to a user
//	@Tags			Roles Management
//	@Accept			json
//	@Security		BearerAuth
//	@Param			roleId	path	int	true	"Role Id"
//	@Param			userId	path	int	true	"User Id"
//	@Produce		json
//	@Success		201	{string}	created
//	@Failure		400	{object}	errors.HttpErrorMap
//	@Failure		401	{object}	errors.HttpErrorMap
//	@Failure		404	{object}	errors.HttpErrorMap
//	@Failure		500	{object}	errors.HttpErrorMap
//	@Router			/role-management/role/{roleId}/user/{userId} [post]
func (r *roleManagementHandler) AddRoleToUser(c *fiber.Ctx) error {
	adminId, err := getAuthUserId(c, r.logger)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	roleId, err := getIntParam(c, r.logger, "roleId")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	userId, err := getIntParam(c, r.logger, "userId")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	err = r.service.AddRoleToUser(*adminId, *roleId, *userId)
	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).JSON(httpErr)
	}

	return c.SendStatus(http.StatusNoContent)
}

// Get Roles By User Id godoc
//
//	@Description	get roles by user id
//	@Tags			Roles Management
//	@Accept			json
//	@Security		BearerAuth
//	@Param			userId	path	int	true	"User Id"
//	@Produce		json
//	@Success		200	{object}	model.Role
//	@Failure		400	{object}	errors.HttpErrorMap
//	@Failure		401	{object}	errors.HttpErrorMap
//	@Failure		404	{object}	errors.HttpErrorMap
//	@Failure		500	{object}	errors.HttpErrorMap
//	@Router			/role-management/user/{userId} [get]
func (r *roleManagementHandler) GetRolesByUserId(c *fiber.Ctx) error {
	adminId, err := getAuthUserId(c, r.logger)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	userId, err := getIntParam(c, r.logger, "userId")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	roles, err := r.service.GetRolesByUserId(*adminId, *userId)
	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).JSON(httpErr)
	}

	return c.Status(http.StatusOK).JSON(roles)
}

// Remove Role From User godoc
//
//	@Description	remove a role from a user
//	@Tags			Roles Management
//	@Accept			json
//	@Security		BearerAuth
//	@Param			roleId	path	int	true	"Role Id"
//	@Param			userId	path	int	true	"User Id"
//	@Produce		json
//	@Success		204	
//	@Failure		400	{object}	errors.HttpErrorMap
//	@Failure		401	{object}	errors.HttpErrorMap
//	@Failure		404	{object}	errors.HttpErrorMap
//	@Failure		500	{object}	errors.HttpErrorMap
//	@Router			/role-management/role/{roleId}/user/{userId} [delete]
func (r *roleManagementHandler) RemoveRoleFromUser(c *fiber.Ctx) error {
	adminId, err := getAuthUserId(c, r.logger)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	roleId, err := getIntParam(c, r.logger, "roleId")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	userId, err := getIntParam(c, r.logger, "userId")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	err = r.service.RemoveRoleFromUser(*adminId, *roleId, *userId)
	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).JSON(httpErr)
	}

	return c.SendStatus(http.StatusNoContent)
}
