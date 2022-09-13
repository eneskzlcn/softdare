package entity

import (
	"errors"
)

type Error string

func (e *Error) Err() error {
	return errors.New(e.String())
}
func (e *Error) String() string {
	return *(*string)(e)
}

var (
	InvalidPostID               Error = "invalid post ID"
	InvalidContent              Error = "invalid content"
	ErrInvalidCommentContent    Error = "invalid comment content"
	CouldNotTakeUserFromContext Error = "could not take the user from ctx "
)
