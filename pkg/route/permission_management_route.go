package route

import (
	h "app/pkg/handler"
	mi "app/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func NewPermissionManagementRoutes(app *fiber.App, handler h.PermissionManagementHandler, middleware mi.AuthMiddleware) {
	permissionManagement := app.Group("/permission-management")
	permissionManagement.Use(middleware.JWTParser)
	permissionManagement.Post("/:permissionId/:roleId", handler.AddPermissionToRole)
	permissionManagement.Post("/:roleId", handler.AddPermissionsToRole)
	permissionManagement.Get("/role/:roleId", handler.GetPermissionsByRoleId)
	permissionManagement.Get("/user/:userId", handler.GetPermissionsByUserId)
	permissionManagement.Delete(":permissionId/:roleId", handler.RemovePermissionFromRole)
	permissionManagement.Delete("/:roleId", handler.RemovePermissionsFromRole)
}
