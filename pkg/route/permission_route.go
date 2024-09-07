package route

import (
	h "app/pkg/handler"
	mi "app/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func NewPermissionRoutes(app *fiber.App, handler h.PermissionHandler, middleware mi.AuthMiddleware) {
	permission := app.Group("/permissions")
	permission.Use(middleware.JWTParser)
	permission.Post("/", handler.Add)
	permission.Post("/all", handler.AddAll)
	permission.Get("/", handler.GetAll)
	permission.Get("/:id", handler.GetOne)
	permission.Put("/:id", handler.Edit)
	permission.Delete("/:id", handler.Remove)
}
