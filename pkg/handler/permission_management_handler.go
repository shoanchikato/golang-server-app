package handler

import (
	e "app/pkg/errors"
	ef "app/pkg/httperrorfmt"
	s "app/pkg/service"
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type PermissionManagementHandler interface {
	AddPermissionToRole(c *fiber.Ctx) error
	AddPermissionsToRole(c *fiber.Ctx) error
	GetPermissionsByRoleId(c *fiber.Ctx) error
	GetPermissionsByUserId(c *fiber.Ctx) error
	RemovePermissionFromRole(c *fiber.Ctx) error
	RemovePermissionsFromRole(c *fiber.Ctx) error
}

type permissionManagementHandler struct {
	service ef.PermissionManagementHttpErrorFmt
	logger  s.Logger
}

func NewPermissionManagementHandler(service ef.PermissionManagementHttpErrorFmt, logger s.Logger) PermissionManagementHandler {
	return &permissionManagementHandler{service, logger}
}

// Add Permission To Role godoc
//
//	@Description	add a permission to a role
//	@Tags			Permissions Management
//	@Accept			json
//	@Security		BearerAuth
//	@Param			permissionId	path	int	true	"Permission Id"
//	@Param			roleId			path	int	true	"Role Id"
//	@Produce		json
//	@Success		201	{string}	created
//	@Failure		400	{object}	errors.HttpErrorMap
//	@Failure		401	{object}	errors.HttpErrorMap
//	@Failure		404	{object}	errors.HttpErrorMap
//	@Failure		500	{object}	errors.HttpErrorMap
//	@Router			/permission-management/permission/{permissionId}/role/{roleId} [post]
func (p *permissionManagementHandler) AddPermissionToRole(c *fiber.Ctx) error {
	adminId, err := getAuthUserId(c, p.logger)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	permissionId, err := getIntParam(c, p.logger, "permissionId")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	roleId, err := getIntParam(c, p.logger, "roleId")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	err = p.service.AddPermissionToRole(*adminId, *permissionId, *roleId)
	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).JSON(httpErr)
	}

	return c.SendStatus(http.StatusCreated)
}

// Add Permissions To Role godoc
//
//	@Description	add permissions to a role
//	@Tags			Permissions Management
//	@Accept			json
//	@Security		BearerAuth
//	@Param			roleId		path	int		true	"Role Id"
//	@Param			permissions	body	[]int	true	"Permission Ids"
//	@Produce		json
//	@Success		201	{string}	created
//	@Failure		400	{object}	errors.HttpErrorMap
//	@Failure		401	{object}	errors.HttpErrorMap
//	@Failure		404	{object}	errors.HttpErrorMap
//	@Failure		500	{object}	errors.HttpErrorMap
//	@Router			/permission-management/role/{roleId} [post]
func (p *permissionManagementHandler) AddPermissionsToRole(c *fiber.Ctx) error {
	adminId, err := getAuthUserId(c, p.logger)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	permissionIds := []int{}
	err = c.BodyParser(&permissionIds)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	roleId, err := getIntParam(c, p.logger, "roleId")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	err = p.service.AddPermissionsToRole(*adminId, permissionIds, *roleId)
	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).JSON(httpErr)
	}

	return c.SendStatus(http.StatusCreated)
}

// Get Permissions By Role Id godoc
//
//	@Description	get permissions by role id
//	@Tags			Permissions Management
//	@Accept			json
//	@Security		BearerAuth
//	@Param			roleId	path	int	true	"Role Id"
//	@Produce		json
//	@Success		200	{object}	[]model.Permission
//	@Failure		400	{object}	errors.HttpErrorMap
//	@Failure		401	{object}	errors.HttpErrorMap
//	@Failure		404	{object}	errors.HttpErrorMap
//	@Failure		500	{object}	errors.HttpErrorMap
//	@Router			/permission-management/role/{roleId} [get]
func (p *permissionManagementHandler) GetPermissionsByRoleId(c *fiber.Ctx) error {
	adminId, err := getAuthUserId(c, p.logger)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	roleId, err := getIntParam(c, p.logger, "roleId")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	permissions, err := p.service.GetPermissionsByRoleId(*adminId, *roleId)
	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).JSON(httpErr)
	}

	return c.Status(http.StatusOK).JSON(permissions)
}

// Get Permissons By User Id godoc
//
//	@Description	get permissions by user id
//	@Tags			Permissions Management
//	@Accept			json
//	@Security		BearerAuth
//	@Param			userId	path	int	true	"User Id"
//	@Produce		json
//	@Success		200	{object}	[]model.Permission
//	@Failure		400	{object}	errors.HttpErrorMap
//	@Failure		401	{object}	errors.HttpErrorMap
//	@Failure		404	{object}	errors.HttpErrorMap
//	@Failure		500	{object}	errors.HttpErrorMap
//	@Router			/permission-management/user/{userId} [get]
func (p *permissionManagementHandler) GetPermissionsByUserId(c *fiber.Ctx) error {
	adminId, err := getAuthUserId(c, p.logger)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	userId, err := getIntParam(c, p.logger, "userId")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	permissions, err := p.service.GetPermissionsByUserId(*adminId, *userId)
	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).JSON(httpErr)
	}

	return c.Status(http.StatusOK).JSON(permissions)
}

// Remove Permission From Role godoc
//
//	@Description	remove a permission from a role
//	@Tags			Permissions Management
//	@Accept			json
//	@Security		BearerAuth
//	@Param			permissionId	path	int	true	"Permission Id"
//	@Param			roleId			path	int	true	"Role Id"
//	@Produce		json
//	@Success		204 
//	@Failure		400	{object}	errors.HttpErrorMap
//	@Failure		401	{object}	errors.HttpErrorMap
//	@Failure		404	{object}	errors.HttpErrorMap
//	@Failure		500	{object}	errors.HttpErrorMap
//	@Router			/permission-management/permission/{permissionId}/role/{roleId} [delete]
func (p *permissionManagementHandler) RemovePermissionFromRole(c *fiber.Ctx) error {
	adminId, err := getAuthUserId(c, p.logger)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	roleId, err := getIntParam(c, p.logger, "roleId")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	permissionId, err := getIntParam(c, p.logger, "permissionId")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	err = p.service.RemovePermissionFromRole(*adminId, *roleId, *permissionId)
	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).JSON(httpErr)
	}

	return c.SendStatus(http.StatusNoContent)
}

// Remove Permissions From Role godoc
//
//	@Description	remove permissions from a role
//	@Tags			Permissions Management
//	@Accept			json
//	@Security		BearerAuth
//	@Param			roleId		path	int		true	"Role Id"
//	@Param			permissions	body	[]int	true	"Permission Ids"
//	@Produce		json
//	@Success		204
//	@Failure		400	{object}	errors.HttpErrorMap
//	@Failure		401	{object}	errors.HttpErrorMap
//	@Failure		404	{object}	errors.HttpErrorMap
//	@Failure		500	{object}	errors.HttpErrorMap
//	@Router			/permission-management/role/{roleId} [delete]
func (p *permissionManagementHandler) RemovePermissionsFromRole(c *fiber.Ctx) error {
	adminId, err := getAuthUserId(c, p.logger)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	permissionIds := []int{}
	err = c.BodyParser(&permissionIds)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	roleId, err := getIntParam(c, p.logger, "roleId")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	err = p.service.RemovePermissionsFromRole(*adminId, *roleId, permissionIds)
	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).JSON(httpErr)
	}

	return c.SendStatus(http.StatusNoContent)
}
