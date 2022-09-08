package login

import "errors"

var (
	ErrInvalidHandlerArgs = errors.New("service, renderer and session provider can not be nil")
	ErrRepositoryNil      = errors.New("given repository is nil")
	ErrDBNil              = errors.New("given database is nil")
)
