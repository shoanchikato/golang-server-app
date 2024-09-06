package handler

import (
	"errors"
	"strconv"
	mi "app/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func getId(c *fiber.Ctx) (*int, error) {
	idStr := c.Params("id")

	int, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, err
	}

	return &int, nil
}

func getAuthUserId(c *fiber.Ctx) (*int, error) {
	ctx := c.UserContext()
	userIdKey := mi.UserContextKey("userId")
	userId := ctx.Value(userIdKey).(int)
	if userId == 0 {
		return nil, errors.New("user not authorized: user unknown")
	}

	return &userId, nil

}
