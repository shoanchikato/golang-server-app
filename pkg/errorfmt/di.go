package errorfmt

import (
	s "app/pkg/service"
	a "app/pkg/authorization"
)

type ErrorFmts struct {
	Auth AuthErrorFmt
	User UserErrorFmt
}

func ErrorFmtDi(errorFmt s.ErrorFmt, jwt s.JWTService, authorizations *a.Authorizations) *ErrorFmts {
	user := NewUserErrorFmt(authorizations.User, errorFmt)
	auth := NewAuthErrorFmt(authorizations.Auth, jwt, errorFmt)

	return &ErrorFmts{auth, user}
}