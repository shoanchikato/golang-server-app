package handler

import (
	e "app/pkg/errors"
	mi "app/pkg/middleware"
	s "app/pkg/service"
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func getIntParam(c *fiber.Ctx, logger s.Logger, name string) (*int, error) {
	strValue := c.Params(name)

	intValue, err := strconv.Atoi(strValue)
	if err != nil {

		logger.Error(e.NewIntParamError(name, strValue).Error())
		return nil, e.NewIntParamError(name, strValue)
	}

	return &intValue, nil
}

func getAuthUserId(c *fiber.Ctx, logger s.Logger) (*int, error) {
	ctx := c.UserContext()
	userIdKey := mi.UserContextKey("userId")
	userId := ctx.Value(userIdKey).(int)
	if userId == 0 {

		err := errors.New("user not authorized: user unknown")
		logger.Error(err.Error())
		return nil, err
	}

	return &userId, nil
}
