package httperrorfmt

import (
	a "app/pkg/authorization"
	m "app/pkg/model"
	s "app/pkg/service"
)

type AuthHttpErrorFmt interface {
	Login(credentials *m.Credentials) (tokens *m.Tokens, err error)
	ResetPassword(username, newPassword string) error
}

type authHttpErrorFmt struct {
	auth    a.AuthAuthorization
	jwt     s.JWTService
	service s.HttpErrorFmt
}

func NewAuthHttpErrorFmt(auth a.AuthAuthorization, jwt s.JWTService, service s.HttpErrorFmt) AuthHttpErrorFmt {
	return &authHttpErrorFmt{auth, jwt, service}
}

// Login
func (a *authHttpErrorFmt) Login(credentials *m.Credentials) (*m.Tokens, error) {
	userId, err := a.auth.Login(credentials)
	if err != nil {
		return nil, a.service.GetError(err)
	}

	tokens, err := a.jwt.GetTokens(*userId)
	if err != nil {
		return nil, a.service.GetError(err)
	}

	return tokens, nil
}

// ResetPassword
func (a *authHttpErrorFmt) ResetPassword(username string, newPassword string) error {
	err := a.auth.ResetPassword(username, newPassword)
	if err != nil {
		return a.service.GetError(err)
	}

	return nil
}
