package errorfmt

import (
	a "app/pkg/authorization"
	m "app/pkg/model"
	s "app/pkg/service"
)

type AuthErrorFmt interface {
	Login(credentials *m.Credentials) (tokens *m.Tokens, err error)
	ResetPassword(username, newPassword string) error
}

type authErrorFmt struct {
	auth    a.AuthAuthorization
	jwt     s.JWTService
	service s.ErrorFmt
}

func NewAuthErrorFmt(auth a.AuthAuthorization, jwt s.JWTService, service s.ErrorFmt) AuthErrorFmt {
	return &authErrorFmt{auth, jwt, service}
}

// Login
func (a *authErrorFmt) Login(credentials *m.Credentials) (*m.Tokens, error) {
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
func (a *authErrorFmt) ResetPassword(username string, newPassword string) error {
	err := a.auth.ResetPassword(username, newPassword)
	if err != nil {
		return a.service.GetError(err)
	}

	return nil
}
