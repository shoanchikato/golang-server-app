package handler

import (
	mi "app/pkg/middleware"
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func getId(c *fiber.Ctx) (*int, error) {
	idStr := c.Params("id")

	intValue, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, err
	}

	return &intValue, nil
}

func getIntParam(c *fiber.Ctx, name string) (*int, error) {
	strValue := c.Params(name)

	intValue, err := strconv.Atoi(strValue)
	if err != nil {
		return nil, err
	}

	return &intValue, nil
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
