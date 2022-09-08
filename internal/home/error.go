package home

import "errors"

var (
	ErrRendererNil        = errors.New("given renderer is nil")
	ErrSessionProviderNil = errors.New("given session provider is nil")
)
