package route

import (
	h "app/pkg/handler"
	mi "app/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func NewPostRoutes(app *fiber.App, handler h.PostHandler, middleware mi.AuthMiddleware) {
	post := app.Group("/posts")
	post.Use(middleware.JWTParser)
	post.Post("/", handler.Add)
	post.Post("/all", handler.AddAll)
	post.Get("/", handler.GetAll)
	post.Get("/:id", handler.GetOne)
	post.Put("/:id", handler.Edit)
	post.Delete("/:id", handler.Remove)
}
