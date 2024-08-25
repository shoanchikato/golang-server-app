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
	ErrRepoLastInsertedID = errors.New("repo: error getting last inserted id")

	ErrRepoAdd      = errors.New("repo: add")
	ErrRepoEdit     = errors.New("repo: edit")
	ErrRepoGetAll   = errors.New("repo: get all")
	ErrRepoGetOne   = errors.New("repo: get one")
	ErrRepoGetMore  = errors.New("repo: get more")
	ErrRepoRemove   = errors.New("repo: remove post")
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
