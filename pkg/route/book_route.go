package route

import (
	h "app/pkg/handler"
	mi "app/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func NewBookRoutes(app *fiber.App, handler h.BookHandler, middleware mi.AuthMiddleware) {
	book := app.Group("/books")
	book.Use(middleware.JWTParser)
	book.Post("/", handler.Add)
	book.Post("/all", handler.AddAll)
	book.Get("/", handler.GetAll)
	book.Get("/:id", handler.GetOne)
	book.Put("/:id", handler.Edit)
	book.Delete("/:id", handler.Remove)
}
