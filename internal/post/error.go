package post

import "errors"

var (
	ErrDatabaseNil           = errors.New("given database is nil")
	ErrPostContentValidation = errors.New("post content is not valid")
	ErrUnauthorized          = errors.New("unauthorized user")
	ErrPostRepositoryNil     = errors.New("post repository is nil")
	ErrPostServiceNil        = errors.New("post service is nil")
	ErrRendererNil           = errors.New("renderer is nil")
	ErrSessionProviderNil    = errors.New("session provider is nil")
)
