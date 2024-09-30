package adapters

import "errors"

var (
	ErrUnsupportedCredentials = errors.New("unsupported credentials")
	ErrInvalidCredentials     = errors.New("invalid credentials")
)
