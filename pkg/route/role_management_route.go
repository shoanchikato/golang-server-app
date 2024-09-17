package route

import (
	h "app/pkg/handler"
	mi "app/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func NewRoleManagementRoutes(app *fiber.App, handler h.RoleManagementHandler, middleware mi.AuthMiddleware) {
	roleManagement := app.Group("/role-management")
	roleManagement.Use(middleware.JWTParser)
	roleManagement.Post("/:roleId/:userId", handler.AddRoleToUser)
	roleManagement.Get("/user/:userId", handler.GetRoleByUserId)
	roleManagement.Delete("/:roleId/:userId", handler.RemoveRoleFromUser)
}
