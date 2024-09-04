package route

import (
	h "app/pkg/handler"

	"github.com/gofiber/fiber/v2"
)

func NewUserRoutes(app *fiber.App, handler h.UserHandler, middleware h.AuthMiddleware) {
	auth := app.Group("/users")
	auth.Use(middleware.JWTParser)
	auth.Post("/", handler.Add)
}
