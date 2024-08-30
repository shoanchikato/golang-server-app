package service

import (
	e "app/pkg/errors"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService interface {
	ParseToken(*string) (*Token, error)
	GetAccessToken(userId int) (string, error)
	GetRefreshToken(userId int) (string, error)
}

type jwtService struct {
	secret               *string
	accessTokenDuration  time.Duration
	refreshTokenDuration time.Duration
}

type Token struct {
	UserId  int    `json:"user_id"`
	Expires string `json:"expires"`
	Issued  string `json:"issued"`
	jwt.RegisteredClaims
}

func (t *Token) GetExpires() (time.Time, error ){
	return time.Parse(time.RFC3339, t.Expires)
}

func (t *Token) GetIssued() (time.Time, error ){
	return time.Parse(time.RFC3339, t.Issued)
}

func NewJWTService(
	secret *string,
	accessTokenDuration time.Duration,
	refreshTokenDuration time.Duration,
) JWTService {
	return &jwtService{secret, accessTokenDuration, refreshTokenDuration}
}

func (j *jwtService) newToken(userId int, duration time.Duration) (string, error) {
	claims := Token{
		userId,
		time.Now().Add(duration).Format(time.RFC3339),
		time.Now().Format(time.RFC3339),
		jwt.RegisteredClaims{},
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

func (j *jwtService) ParseToken(tokenString *string) (*Token, error) {
	op := func(token *jwt.Token) (interface{}, error) {
		return []byte(*j.secret), nil
	}

	token, err := jwt.ParseWithClaims(*tokenString, &Token{}, op)

	switch {
	case token.Valid:
		claims, ok := token.Claims.(*Token)
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
