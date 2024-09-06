package route

import (
	h "app/pkg/handler"
	mi "app/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func NewUserRoutes(app *fiber.App, handler h.UserHandler, middleware mi.AuthMiddleware) {
	auth := app.Group("/users")
	auth.Use(middleware.JWTParser)
	auth.Post("/", handler.Add)
	auth.Post("/all", handler.AddAll)
	auth.Get("/", handler.GetAll)
	auth.Get("/:id", handler.GetOne)
	auth.Put("/:id", handler.Edit)
	auth.Delete("/:id", handler.Remove)
}
