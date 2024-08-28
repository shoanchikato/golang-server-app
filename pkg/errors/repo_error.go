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

	ErrRepoNotFound = errors.New("repo: not found")
)

type RepoNotFoundError string

func NewErrRepoNotFound(id string) error {
	err := RepoNotFoundError(id)
	return &err
}

func (e *RepoNotFoundError) Error() string {
	return fmt.Sprintf("%v not found", *e)
}

func (e *RepoNotFoundError) Is(target error) bool {
	return ErrRepoNotFound == target
}
