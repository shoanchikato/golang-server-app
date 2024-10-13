package route

import (
	h "app/pkg/handler"
	mi "app/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func NewRoleManagementRoutes(app *fiber.App, handler h.RoleManagementHandler, middleware mi.AuthMiddleware) {
	roleManagement := app.Group("/role-management")
	roleManagement.Use(middleware.JWTParser)
	roleManagement.Post("/role/:roleId/user/:userId", handler.AddRoleToUser)
	roleManagement.Get("/user/:userId", handler.GetRolesByUserId)
	roleManagement.Delete("/role/:roleId/user/:userId", handler.RemoveRoleFromUser)
}
