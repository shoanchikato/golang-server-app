package handler

import (
	e "app/pkg/errors"
	ef "app/pkg/httperrorfmt"
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
}

func NewPermissionManagementHandler(service ef.PermissionManagementHttpErrorFmt) PermissionManagementHandler {
	return &permissionManagementHandler{service}
}

// AddPermissionToRole
func (p *permissionManagementHandler) AddPermissionToRole(c *fiber.Ctx) error {
	adminId, err := getAuthUserId(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	permissionId, err := getIntParam(c, "permissionId")
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(e.ErrProvideNumericId.Error())
	}

	roleId, err := getIntParam(c, "roleId")
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(e.ErrProvideNumericId.Error())
	}

	err = p.service.AddPermissionToRole(*adminId, *permissionId, *roleId)
	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).SendString(httpErr.Message)
	}

	return c.SendStatus(http.StatusAccepted)
}

// AddPermissionsToRole
func (p *permissionManagementHandler) AddPermissionsToRole(c *fiber.Ctx) error {
	adminId, err := getAuthUserId(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	permissionIds := []int{}
	err = c.BodyParser(&permissionIds)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	roleId, err := getIntParam(c, "roleId")
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(e.ErrProvideNumericId.Error())
	}

	err = p.service.AddPermissionsToRole(*adminId, permissionIds, *roleId)
	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).SendString(httpErr.Message)
	}

	return c.SendStatus(http.StatusAccepted)
}

// GetPermissionsByRoleId
func (p *permissionManagementHandler) GetPermissionsByRoleId(c *fiber.Ctx) error {
	adminId, err := getAuthUserId(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	roleId, err := getIntParam(c, "roleId")
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(e.ErrProvideNumericId.Error())
	}

	permissions, err := p.service.GetPermissionsByRoleId(*adminId, *roleId)
	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).SendString(httpErr.Message)
	}

	return c.Status(http.StatusOK).JSON(permissions)
}

// GetPermissonsByUserId
func (p *permissionManagementHandler) GetPermissionsByUserId(c *fiber.Ctx) error {
	adminId, err := getAuthUserId(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	userId, err := getIntParam(c, "userId")
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(e.ErrProvideNumericId.Error())
	}

	permissions, err := p.service.GetPermissionsByUserId(*adminId, *userId)
	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).SendString(httpErr.Message)
	}

	return c.Status(http.StatusOK).JSON(permissions)
}

// RemovePermissionFromRole
func (p *permissionManagementHandler) RemovePermissionFromRole(c *fiber.Ctx) error {
	adminId, err := getAuthUserId(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	roleId, err := getIntParam(c, "roleId")
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(e.ErrProvideNumericId.Error())
	}

	permissionId, err := getIntParam(c, "permissionId")
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(e.ErrProvideNumericId.Error())
	}

	err = p.service.RemovePermissionFromRole(*adminId, *roleId, *permissionId)
	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).SendString(httpErr.Message)
	}

	return c.SendStatus(http.StatusAccepted)
}

// RemovePermissionsFromRole
func (p *permissionManagementHandler) RemovePermissionsFromRole(c *fiber.Ctx) error {
	adminId, err := getAuthUserId(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	permissionIds := []int{}
	err = c.BodyParser(&permissionIds)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	roleId, err := getIntParam(c, "roleId")
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(e.ErrProvideNumericId.Error())
	}

	err = p.service.RemovePermissionsFromRole(*adminId, *roleId, permissionIds)
	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).SendString(httpErr.Message)
	}

	return c.SendStatus(http.StatusAccepted)
}
