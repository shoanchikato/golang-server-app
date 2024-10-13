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

type UserHandler interface {
	Add(c *fiber.Ctx) error
	AddAll(c *fiber.Ctx) error
	Edit(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	GetOne(c *fiber.Ctx) error
	Remove(c *fiber.Ctx) error
}

type userHandler struct {
	service ef.UserHttpErrorFmt
	logger  s.Logger
}

func NewUserHandler(service ef.UserHttpErrorFmt, logger s.Logger) UserHandler {
	return &userHandler{service, logger}
}

// Add godoc
//
//	@Description	add a user
//	@Tags			Users
//	@Accept			json
//	@Security		BearerAuth
//	@Param			user	body	model.User	true	"User Details"
//	@Produce		json
//	@Success		201	{string}	created
//	@Failure		400	{object}	errors.HttpErrorMap
//	@Failure		401	{object}	errors.HttpErrorMap
//	@Failure		404	{object}	errors.HttpErrorMap
//	@Failure		500	{object}	errors.HttpErrorMap
//	@Router			/users [post]
func (u *userHandler) Add(c *fiber.Ctx) error {
	user := m.User{}

	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	userId, err := getAuthUserId(c, u.logger)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	err = u.service.Add(*userId, &user)

	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).JSON(httpErr)
	}

	return c.Status(http.StatusCreated).JSON(user)
}

// Add All godoc
//
//	@Description	add an array of users
//	@Tags			Users
//	@Accept			json
//	@Security		BearerAuth
//	@Param			user	body	[]model.User	true	"Users' Details"
//	@Produce		json
//	@Success		201	{string}	created
//	@Failure		400	{object}	errors.HttpErrorMap
//	@Failure		401	{object}	errors.HttpErrorMap
//	@Failure		404	{object}	errors.HttpErrorMap
//	@Failure		500	{object}	errors.HttpErrorMap
//	@Router			/users/all [post]
func (u *userHandler) AddAll(c *fiber.Ctx) error {
	users := []m.User{}

	err := c.BodyParser(&users)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	userId, err := getAuthUserId(c, u.logger)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	newUsers := []*m.User{}
	for i := range users {
		newUsers = append(newUsers, &users[i])
	}

	err = u.service.AddAll(*userId, &newUsers)

	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).JSON(httpErr)
	}

	return c.Status(http.StatusCreated).JSON(newUsers)
}

// Edit godoc
//
//	@Description	edit a user
//	@Tags			Users
//	@Accept			json
//	@Security		BearerAuth
//	@Param			user	body	model.User	true	"User Details"
//	@Param			id		path	int			true	"User Id"
//	@Produce		json
//	@Success		201	{string}	created
//	@Failure		400	{object}	errors.HttpErrorMap
//	@Failure		401	{object}	errors.HttpErrorMap
//	@Failure		404	{object}	errors.HttpErrorMap
//	@Failure		500	{object}	errors.HttpErrorMap
//	@Router			/users/{id} [put]
func (u *userHandler) Edit(c *fiber.Ctx) error {
	user := m.EditUser{}

	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	id, err := getIntParam(c, u.logger, "id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	userId, err := getAuthUserId(c, u.logger)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	err = u.service.Edit(*userId, *id, &user)

	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).JSON(httpErr)
	}

	return c.Status(http.StatusCreated).JSON(user)
}

// Get All godoc
//
//	@Description	get all users
//	@Tags			Users
//	@Accept			json
//	@Security		BearerAuth
//	@Produce		json
//	@Success		200	{object}	[]model.User
//	@Failure		400	{object}	errors.HttpErrorMap
//	@Failure		401	{object}	errors.HttpErrorMap
//	@Failure		404	{object}	errors.HttpErrorMap
//	@Failure		500	{object}	errors.HttpErrorMap
//	@Router			/users [get]
func (u *userHandler) GetAll(c *fiber.Ctx) error {
	userId, err := getAuthUserId(c, u.logger)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	users, err := u.service.GetAll(*userId, 0, 50)

	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).JSON(httpErr)
	}

	return c.Status(http.StatusOK).JSON(users)
}

// Get One godoc
//
//	@Description	get one user
//	@Tags			Users
//	@Accept			json
//	@Security		BearerAuth
//	@Param			id	path	int	true	"User Id"
//	@Produce		json
//	@Success		200	{object}	model.User
//	@Failure		400	{object}	errors.HttpErrorMap
//	@Failure		401	{object}	errors.HttpErrorMap
//	@Failure		404	{object}	errors.HttpErrorMap
//	@Failure		500	{object}	errors.HttpErrorMap
//	@Router			/users/{id} [get]
func (u *userHandler) GetOne(c *fiber.Ctx) error {
	userId, err := getAuthUserId(c, u.logger)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	id, err := getIntParam(c, u.logger, "id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	user, err := u.service.GetOne(*userId, *id)

	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).JSON(httpErr)
	}

	return c.Status(http.StatusOK).JSON(user)
}

// Remove User godoc
//
//	@Description	delete a user
//	@Tags			Users
//	@Accept			json
//	@Security		BearerAuth
//	@Param			id	path	int	true	"User Id"
//	@Produce		json
//	@Success		204	
//	@Failure		400	{object}	errors.HttpErrorMap
//	@Failure		401	{object}	errors.HttpErrorMap
//	@Failure		404	{object}	errors.HttpErrorMap
//	@Failure		500	{object}	errors.HttpErrorMap
//	@Router			/users/{id} [delete]
func (u *userHandler) Remove(c *fiber.Ctx) error {
	userId, err := getAuthUserId(c, u.logger)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	id, err := getIntParam(c, u.logger, "id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(e.NewHttpErrorMap(err))
	}

	err = u.service.Remove(*userId, *id)

	httpErr := &e.HttpError{}
	if errors.As(err, &httpErr) {
		return c.Status(httpErr.HTTPStatus).JSON(httpErr)
	}

	return c.SendStatus(http.StatusNoContent)
}
