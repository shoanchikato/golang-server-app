package handler

import (
	aa "app/pkg/authorization"
	m "app/pkg/model"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	Add(c *fiber.Ctx) error
	AddAll(c *fiber.Ctx) error
	Edit(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	GetOne(c *fiber.Ctx) error
	Remove(c *fiber.Ctx) error
}

type userHandler struct {
	auth aa.UserAuthorization
}

func NewUserHandler(auth aa.UserAuthorization) UserHandler {
	return &userHandler{auth}
}

// Add implements UserHandler.
func (u *userHandler) Add(c *fiber.Ctx) error {
	user := m.User{}

	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	ctx := c.UserContext()
	userIdKey := userContextKey("userId")
	userId := ctx.Value(userIdKey).(int)
	if userId == 0 {
		return c.SendStatus(http.StatusUnauthorized)
	}

	err = u.auth.Add(userId, &user)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	return c.Status(http.StatusCreated).JSON(user)
}

// AddAll implements UserHandler.
func (u *userHandler) AddAll(c *fiber.Ctx) error {
	panic("unimplemented")
}

// Edit implements UserHandler.
func (u *userHandler) Edit(c *fiber.Ctx) error {
	panic("unimplemented")
}

// GetAll implements UserHandler.
func (u *userHandler) GetAll(c *fiber.Ctx) error {
	panic("unimplemented")
}

// GetOne implements UserHandler.
func (u *userHandler) GetOne(c *fiber.Ctx) error {
	panic("unimplemented")
}

// Remove implements UserHandler.
func (u *userHandler) Remove(c *fiber.Ctx) error {
	panic("unimplemented")
}
