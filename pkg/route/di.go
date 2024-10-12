package route

import (
	h "app/pkg/handler"
	mi "app/pkg/middleware"

	"github.com/gofiber/fiber/v2"

	_ "app/docs"
	swagger "github.com/gofiber/swagger"
)

//	@title		Swagger Golang Server App API
//	@version	2.0
//	@host		localhost:3000
//	@BasePath	/
func Routes(app *fiber.App, handlers *h.Handlers, authMiddleware mi.AuthMiddleware) {
	app.Get("/swagger/*", swagger.HandlerDefault)
	NewAuthRoutes(app, handlers.Auth, authMiddleware)
	NewUserRoutes(app, handlers.User, authMiddleware)
	NewRoleRoutes(app, handlers.Role, authMiddleware)
	NewRoleManagementRoutes(app, handlers.RoleManagement, authMiddleware)
	NewPermissionRoutes(app, handlers.Permission, authMiddleware)
	NewPermissionManagementRoutes(app, handlers.PermissionManagement, authMiddleware)
	NewAuthorRoutes(app, handlers.Author, authMiddleware)
	NewBookRoutes(app, handlers.Book, authMiddleware)
	NewPostRoutes(app, handlers.Post, authMiddleware)
}
