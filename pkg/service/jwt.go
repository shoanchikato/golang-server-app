package service

import (
	e "app/pkg/errors"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService interface {
	ParseToken(*string) (*jwt.RegisteredClaims, error)
	GetAccessToken() (string, error)
	GetRefreshToken() (string, error)
}

type jwtService struct {
	secret               *string
	accessTokenDuration  time.Duration
	refreshTokenDuration time.Duration
}

func NewJWTService(
	secret *string,
	accessTokenDuration time.Duration,
	refreshTokenDuration time.Duration,
) JWTService {
	return &jwtService{secret, accessTokenDuration, refreshTokenDuration}
}

func (j *jwtService) newToken(duration time.Duration) (string, error) {
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Issuer:    "app",
	}

	unsignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := unsignedToken.SignedString([]byte(*j.secret))

	if err != nil {
		return "", err
	}

	return token, nil
}

func (j *jwtService) GetAccessToken() (string, error) {
	token, err := j.newToken(j.accessTokenDuration)
	if err != nil {
		return "", errors.Join(e.ErrGetAccessToken, err)
	}

	return token, nil
}

func (j *jwtService) GetRefreshToken() (string, error) {
	token, err := j.newToken(j.refreshTokenDuration)
	if err != nil {
		return "", errors.Join(e.ErrGetRefreshToken, err)
	}

	return token, nil
}

func (j *jwtService) ParseToken(tokenString *string) (*jwt.RegisteredClaims, error) {
	op := func(token *jwt.Token) (interface{}, error) {
		return []byte(*j.secret), nil
	}

	token, err := jwt.ParseWithClaims(*tokenString, &jwt.RegisteredClaims{}, op)

	switch {
	case token.Valid:
		claims, ok := token.Claims.(*jwt.RegisteredClaims)
		if !ok {
			return nil, errors.Join(e.ErrParseToken, err)
		}

		return claims, nil

	case errors.Is(err, jwt.ErrTokenMalformed):
		return nil, errors.Join(e.ErrInvalidToken, err)

	case errors.Is(err, jwt.ErrTokenSignatureInvalid):
		// Invalid signature
		return nil, errors.Join(e.ErrInvalidToken, err)

	case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
		// Token is either expired or not active yet
		return nil, errors.Join(e.ErrTokenExpired, err)

	default:
		return nil, errors.Join(e.ErrInvalidToken, err)
	}
}
