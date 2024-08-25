package errors

import "errors"

var (
	ErrEncryptionHashPassword    = errors.New("encryption: hash password")
	ErrEncryptionComparePassword = errors.New("encryption: compare password")
)
