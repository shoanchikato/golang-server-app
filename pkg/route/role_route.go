package route

import (
	h "app/pkg/handler"
	mi "app/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func NewRoleRoutes(app *fiber.App, handler h.RoleHandler, middleware mi.AuthMiddleware) {
	role := app.Group("/roles")
	role.Use(middleware.JWTParser)
	role.Post("/", handler.Add)
	role.Post("/all", handler.AddAll)
	role.Get("/", handler.GetAll)
	role.Get("/:id", handler.GetOne)
	role.Put("/:id", handler.Edit)
	role.Delete("/:id", handler.Remove)
}
