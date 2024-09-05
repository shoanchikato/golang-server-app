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

	userId, err := getAuthUserId(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	err = u.auth.Add(*userId, &user)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	return c.Status(http.StatusCreated).JSON(user)
}

// AddAll implements UserHandler.
func (u *userHandler) AddAll(c *fiber.Ctx) error {
	users := []m.User{}

	err := c.BodyParser(&users)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	userId, err := getAuthUserId(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	newUsers := []*m.User{}
	for i := range users {
		newUsers[i] = &users[i]
	}

	err = u.auth.AddAll(*userId, &newUsers)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	return c.Status(http.StatusCreated).JSON(newUsers)
}

// Edit implements UserHandler.
func (u *userHandler) Edit(c *fiber.Ctx) error {
	user := m.User{}

	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	id, err := getId(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("please provide an numeric id")
	}

	userId, err := getAuthUserId(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	err = u.auth.Edit(*userId, *id, &user)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	return c.Status(http.StatusCreated).JSON(user)
}

// GetAll implements UserHandler.
func (u *userHandler) GetAll(c *fiber.Ctx) error {
	userId, err := getAuthUserId(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	users, err := u.auth.GetAll(*userId, 0, 50)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	return c.Status(http.StatusCreated).JSON(users)
}

// GetOne implements UserHandler.
func (u *userHandler) GetOne(c *fiber.Ctx) error {
	userId, err := getAuthUserId(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	id, err := getId(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("please provide an numeric id")
	}

	user, err := u.auth.GetOne(*userId, *id)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	return c.Status(http.StatusCreated).JSON(user)
}

// Remove implements UserHandler.
func (u *userHandler) Remove(c *fiber.Ctx) error {
	userId, err := getAuthUserId(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	id, err := getId(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("please provide an numeric id")
	}

	err = u.auth.Remove(*userId, *id)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	return c.SendStatus(http.StatusAccepted)
}
