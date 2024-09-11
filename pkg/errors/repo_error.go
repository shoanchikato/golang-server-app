package errors

import (
	"encoding/json"
	"errors"
	"fmt"
)

var (
	ErrRepoBeginTx        = errors.New("repo: error at the beginning of a transaction")
	ErrRepoPreparingStmt  = errors.New("repo: error preparing a statement")
	ErrRepoExecutingStmt  = errors.New("repo: error executing statement")
	ErrRepoCommitTx       = errors.New("repo: error committing transaction")
	ErrRepoLastInsertedId = errors.New("repo: error getting last inserted id")

	ErrRepoNotFound  = errors.New("repo: not found")
	ErrRepoDuplicate = errors.New("repo: duplicate entry")
)

// RepoNotFoundError
type RepoNotFoundError struct {
	Entity string
	Id     string
	Err    error
}

func NewErrRepoNotFound(entity, id string) error {
	notFoundErr := fmt.Errorf("%s %s not found", entity, id)
	err := RepoNotFoundError{entity, id, notFoundErr}
	return &err
}

func (e *RepoNotFoundError) Error() string {
	return e.Err.Error()
}

func (e *RepoNotFoundError) Is(target error) bool {
	return ErrRepoNotFound == target
}

func (e *RepoNotFoundError) MarshalJSON() ([]byte, error) {
	return json.Marshal(
		struct {
			Message string `json:"message"`
		}{e.Err.Error()},
	)
}

// RepoDuplicateError
type RepoDuplicateError struct {
	Field string
	Err   error
}

func NewErrRepoDuplicate(field string) error {
	duplicateErr := fmt.Errorf("%s already exists", field)
	err := RepoDuplicateError{field, duplicateErr}
	return &err
}

func (e *RepoDuplicateError) Error() string {
	return e.Err.Error()
}

func (e *RepoDuplicateError) Is(target error) bool {
	return ErrRepoDuplicate == target
}
