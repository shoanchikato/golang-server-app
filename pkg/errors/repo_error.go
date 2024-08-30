package errors

import (
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
	Id    string
}

func NewErrRepoNotFound(entity, id string) error {
	err := RepoNotFoundError{entity, id}
	return &err
}

func (e *RepoNotFoundError) Error() string {
	return fmt.Sprintf("%s %s not found", e.Entity, e.Id)
}

func (e *RepoNotFoundError) Is(target error) bool {
	return ErrRepoNotFound == target
}

// RepoDuplicateError
type RepoDuplicateError struct {
	Field  string
	ErrStr string
}

func NewErrRepoDuplicate(field string) error {
	errStr := fmt.Sprintf("%s already exists", field)
	err := RepoDuplicateError{field, errStr}
	return &err
}

func (e *RepoDuplicateError) Error() string {
	return fmt.Sprintf("repo: duplicate entry: %s", e.ErrStr)
}

func (e *RepoDuplicateError) Is(target error) bool {
	return ErrRepoDuplicate == target
}
