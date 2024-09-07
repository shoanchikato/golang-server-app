package handler

import (
	ef "app/pkg/httperrorfmt"
	s "app/pkg/service"
)

type Handlers struct {
	Auth AuthHandler
	User UserHandler
}

func HandlerDi(httpErrorFmts *ef.HttpErrorFmts, jwt s.JWTService) *Handlers {
	auth := NewAuthHandler(httpErrorFmts.Auth, jwt)
	user := NewUserHandler(httpErrorFmts.User)

	return &Handlers{auth, user}
}
