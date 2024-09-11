package handler

import (
	e "app/pkg/errors"
	ef "app/pkg/httperrorfmt"
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type RoleManagementHandler interface {
	AddRoleToUser(c *fiber.Ctx) error
	GetRoleByUserId(c *fiber.Ctx) error
	RemoveRoleFromUser(c *fiber.Ctx) error
}

type roleManagementHandler struct {
	service ef.RoleManagementHttpErrorFmt
}

func NewRoleManagementHandler(service ef.RoleManagementHttpErrorFmt) RoleManagementHandler {
	return &roleManagementHandler{service}
}

// AddRoleToUser
func (r *roleManagementHandler) AddRoleToUser(c *fiber.Ctx) error {
	adminId, err := getAuthUserId(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	roleId, err := getIntParam(c, "roleId")
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(e.ErrProvideNumericId.Error())
	}

	userId, err := getIntParam(c, "userId")
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(e.ErrProvideNumericId.Error())
	}

	err = r.service.AddRoleToUser(*adminId, *roleId, *userId)
	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).JSON(httpErr.Errs)
	}

	return c.SendStatus(http.StatusAccepted)
}

// GetRoleByUserId
func (r *roleManagementHandler) GetRoleByUserId(c *fiber.Ctx) error {
	adminId, err := getAuthUserId(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	userId, err := getIntParam(c, "userId")
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(e.ErrProvideNumericId.Error())
	}

	role, err := r.service.GetRoleByUserId(*adminId, *userId)
	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).JSON(httpErr.Errs)
	}

	return c.Status(http.StatusOK).JSON(role)
}

// RemoveRoleFromUser
func (r *roleManagementHandler) RemoveRoleFromUser(c *fiber.Ctx) error {
	adminId, err := getAuthUserId(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	roleId, err := getIntParam(c, "roleId")
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(e.ErrProvideNumericId.Error())
	}

	userId, err := getIntParam(c, "userId")
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(e.ErrProvideNumericId.Error())
	}

	err = r.service.RemoveRoleFromUser(*adminId, *roleId, *userId)
	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).JSON(httpErr.Errs)
	}

	return c.SendStatus(http.StatusAccepted)
}
