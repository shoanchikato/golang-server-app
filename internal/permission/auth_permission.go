package permission

import m "app/internal/model"

var (
	AuthLogin         = m.NewPermission("auth: login", "auth", "login")
	AuthResetPassword = m.NewPermission("auth: reset password", "auth", "reset password")
)
