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

// AddPermissionToRole
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

	return c.SendStatus(http.StatusNoContent)
}

// AddPermissionsToRole
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

	return c.SendStatus(http.StatusNoContent)
}

// GetPermissionsByRoleId
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

// GetPermissonsByUserId
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

// RemovePermissionFromRole
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

// RemovePermissionsFromRole
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
