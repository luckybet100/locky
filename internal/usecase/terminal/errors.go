package terminal

import "errors"

var (
	ErrSpaceNotFound         = errors.New("space not found")
	ErrTerminalAlreadyExists = errors.New("terminal already exists")
	ErrInvalidInput          = errors.New("input validation failed")
)
