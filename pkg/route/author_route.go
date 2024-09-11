package route

import (
	h "app/pkg/handler"
	mi "app/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func NewAuthorRoutes(app *fiber.App, handler h.AuthorHandler, middleware mi.AuthMiddleware) {
	author := app.Group("/authors")
	author.Use(middleware.JWTParser)
	author.Post("/", handler.Add)
	author.Post("/all", handler.AddAll)
	author.Get("/", handler.GetAll)
	author.Get("/:id", handler.GetOne)
	author.Put("/:id", handler.Edit)
	author.Delete("/:id", handler.Remove)
}
