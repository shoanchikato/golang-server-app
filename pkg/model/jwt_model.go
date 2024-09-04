package model

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Token struct {
	UserId  int    `json:"user_id"`
	Expires string `json:"expires"`
	Issued  string `json:"issued"`
	jwt.RegisteredClaims
}

func (t *Token) GetExpires() (time.Time, error) {
	return time.Parse(time.RFC3339, t.Expires)
}

func (t *Token) GetIssued() (time.Time, error) {
	return time.Parse(time.RFC3339, t.Issued)
}

func (t *Token) HasExpired() (bool, error) {
	exp, err := time.Parse(time.RFC3339, t.Expires)
	if err != nil {
		return true, err
	}

	return exp.Before(time.Now()), nil
}

type Tokens struct {
	Access  string `json:"access_token"`
	Refresh string `json:"refresh_token"`
}

func NewTokens(access, refresh string) *Tokens {
	return &Tokens{access, refresh}
}
