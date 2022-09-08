package server

import "errors"

var (
	ErrLoggerNil              = errors.New("logger can not be nil")
	ErrSessionProviderNil     = errors.New("session provider can not be empty")
	ErrGivenRouteHandlerNil   = errors.New("given handler is nil")
	ErrServerAddressLength    = errors.New("server address should be the length of 5 it is like= ':4000'")
	ErrServerAddressSyntax    = errors.New("invalid server address syntax. should be like = ':4000' ")
	ErrSessionKeyStringLength = errors.New("session key string must be in 32 character length")
)
