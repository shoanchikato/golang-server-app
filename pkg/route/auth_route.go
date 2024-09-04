package route

import (
	h "app/pkg/handler"

	"github.com/gofiber/fiber/v2"
)

func NewAuthRoutes(app *fiber.App, handler h.AuthHandler, middleware h.AuthMiddleware) {
	auth := app.Group("/")
	auth.Post("/login", handler.Login)
	auth.Use(middleware.JWTParser)
	auth.Post("/reset-password", handler.ResetPassword)
}
