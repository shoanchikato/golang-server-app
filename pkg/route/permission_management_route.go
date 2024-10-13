package route

import (
	h "app/pkg/handler"
	mi "app/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func NewPermissionManagementRoutes(app *fiber.App, handler h.PermissionManagementHandler, middleware mi.AuthMiddleware) {
	permissionManagement := app.Group("/permission-management")
	permissionManagement.Use(middleware.JWTParser)
	permissionManagement.Get("/role/:roleId", handler.GetPermissionsByRoleId)
	permissionManagement.Get("/user/:userId", handler.GetPermissionsByUserId)
	permissionManagement.Post("/role/:roleId", handler.AddPermissionsToRole)
	permissionManagement.Post("/permission/:permissionId/role/:roleId", handler.AddPermissionToRole)
	permissionManagement.Delete("role/:roleId", handler.RemovePermissionsFromRole)
	permissionManagement.Delete("/permission/:permissionId/role/:roleId", handler.RemovePermissionFromRole)
}
