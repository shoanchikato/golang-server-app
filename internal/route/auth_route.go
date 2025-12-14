package route

import (
	h "app/internal/handler"
	mi "app/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func NewAuthRoutes(app *fiber.App, handler h.AuthHandler, middleware mi.AuthMiddleware) {
	auth := app.Group("/")
	auth.Post("/login", handler.Login)
	auth.Use(middleware.JWTParser)
	auth.Post("/reset-password", handler.ResetPassword)
}
