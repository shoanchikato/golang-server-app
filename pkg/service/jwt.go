package service

import (
	e "app/pkg/errors"
	m "app/pkg/model"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService interface {
	ParseToken(*string) (*m.Token, error)
	GetAccessToken(userId int) (string, error)
	GetRefreshToken(userId int) (string, error)
	GetTokens(userId int) (*m.Tokens, error)
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

func (j *jwtService) newToken(userId int, duration time.Duration) (string, error) {
	claims := m.Token{
		UserId:           userId,
		Expires:          time.Now().Add(duration).Format(time.RFC3339),
		Issued:           time.Now().Format(time.RFC3339),
		RegisteredClaims: jwt.RegisteredClaims{},
	}

	unsignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := unsignedToken.SignedString([]byte(*j.secret))

	if err != nil {
		return "", err
	}

	return token, nil
}

func (j *jwtService) GetAccessToken(userId int) (string, error) {
	token, err := j.newToken(userId, j.accessTokenDuration)
	if err != nil {
		return "", errors.Join(e.ErrGetAccessToken, err)
	}

	return token, nil
}

func (j *jwtService) GetRefreshToken(userId int) (string, error) {
	token, err := j.newToken(userId, j.refreshTokenDuration)
	if err != nil {
		return "", errors.Join(e.ErrGetRefreshToken, err)
	}

	return token, nil
}

func (j *jwtService) GetTokens(userId int) (*m.Tokens, error) {
	accessToken, err := j.GetAccessToken(userId)
	if err != nil {
		return nil, err
	}

	refreshToken, err := j.GetRefreshToken(userId)
	if err != nil {
		return nil, err
	}

	return m.NewTokens(accessToken, refreshToken), nil
}

func (j *jwtService) ParseToken(tokenString *string) (*m.Token, error) {
	op := func(token *jwt.Token) (interface{}, error) {
		return []byte(*j.secret), nil
	}

	token, err := jwt.ParseWithClaims(*tokenString, &m.Token{}, op)

	switch {
	case token.Valid:
		claims, ok := token.Claims.(*m.Token)
		if !ok {
			return nil, errors.Join(e.ErrParseToken, err)
		}

		hasExpired, err := claims.HasExpired()
		if err != nil {
			return nil, err
		}

		if hasExpired {
			return nil, errors.Join(e.ErrTokenExpired)
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
