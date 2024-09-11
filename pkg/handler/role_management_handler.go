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
	GetRoleByUserId(c *fiber.Ctx) error
	RemoveRoleFromUser(c *fiber.Ctx) error
}

type roleManagementHandler struct {
	service ef.RoleManagementHttpErrorFmt
	logger  s.Logger
}

func NewRoleManagementHandler(service ef.RoleManagementHttpErrorFmt, logger s.Logger) RoleManagementHandler {
	return &roleManagementHandler{service, logger}
}

// AddRoleToUser
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

	return c.SendStatus(http.StatusAccepted)
}

// GetRoleByUserId
func (r *roleManagementHandler) GetRoleByUserId(c *fiber.Ctx) error {
	adminId, err := getAuthUserId(c, r.logger)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	userId, err := getIntParam(c, r.logger, "userId")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	role, err := r.service.GetRoleByUserId(*adminId, *userId)
	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).JSON(httpErr)
	}

	return c.Status(http.StatusOK).JSON(role)
}

// RemoveRoleFromUser
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

	return c.SendStatus(http.StatusAccepted)
}
