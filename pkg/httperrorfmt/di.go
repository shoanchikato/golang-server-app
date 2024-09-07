package httperrorfmt

import (
	a "app/pkg/authorization"
	s "app/pkg/service"
)

type HttpErrorFmts struct {
	Auth AuthHttpErrorFmt
	User UserHttpErrorFmt
}

func HttpErrorFmtDi(httpErrorFmt s.HttpErrorFmt, jwt s.JWTService, authorizations *a.Authorizations) *HttpErrorFmts {
	user := NewUserHttpErrorFmt(authorizations.User, httpErrorFmt)
	auth := NewAuthHttpErrorFmt(authorizations.Auth, jwt, httpErrorFmt)

	return &HttpErrorFmts{auth, user}
}
