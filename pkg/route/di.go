package route

import (
	h "app/pkg/handler"
	mi "app/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App, handlers *h.Handlers, authMiddleware mi.AuthMiddleware) {
	NewAuthRoutes(app, handlers.Auth, authMiddleware)
	NewUserRoutes(app, handlers.User, authMiddleware)
}
