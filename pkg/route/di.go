package route

import (
	h "app/pkg/handler"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App, handlers *h.Handlers, authMiddleware h.AuthMiddleware) {
	NewAuthRoutes(app, handlers.Auth, authMiddleware)
	NewUserRoutes(app, handlers.User, authMiddleware)
}
