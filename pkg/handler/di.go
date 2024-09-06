package handler

import (
	ef "app/pkg/errorfmt"
	s "app/pkg/service"
)

type Handlers struct {
	Auth AuthHandler
	User UserHandler
}

func HandlerDi(errorFmts *ef.ErrorFmts, jwt s.JWTService) *Handlers {
	auth := NewAuthHandler(errorFmts.Auth, jwt)
	user := NewUserHandler(errorFmts.User)

	return &Handlers{auth, user}
}
