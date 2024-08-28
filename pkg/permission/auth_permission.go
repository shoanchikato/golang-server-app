package permission

import m "app/pkg/model"

var (
	AuthLogin         = m.NewPermission("auth: login", "auth", "login")
	AuthResetPassword = m.NewPermission("auth: reset password", "auth", "reset password")
)
