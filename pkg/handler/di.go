package handler

import (
	a "app/pkg/authorization"
	s "app/pkg/service"
)

type Handlers struct {
	Auth AuthHandler
	User UserHandler
}

func HandlerDi(authorizations *a.Authorizations, jwt s.JWTService) *Handlers {
	auth := NewAuthHandler(authorizations.Auth, jwt)
	user := NewUserHandler(authorizations.User)

	return &Handlers{auth, user}
}
