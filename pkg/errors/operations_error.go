package errors

import "errors"

var (
	ErrOnAdd     = errors.New("add: ")
	ErrOnAddAll     = errors.New("add app: ")
	ErrOnEdit    = errors.New("edit: ")
	ErrOnGetAll  = errors.New("get all: ")
	ErrOnGetOne  = errors.New("get one: ")
	ErrOnGetMore = errors.New("get more: ")
	ErrOnRemove  = errors.New("remove post: ")
)
