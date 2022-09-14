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
	InvalidCommentContent       Error = "invalid comment content"
	NilRepository               Error = "given repository is nil"
	CouldNotTakeUserFromContext Error = "could not take the user from ctx"
	NilDatabase                 Error = "given database is nil"
	Unauthorized                Error = "given unauthorized user"
	NilSession                  Error = "given session is nil"
	NilService                  Error = "given service is nil"
	UserNotFound                Error = "user not found"
	UsernameAlreadyTaken        Error = "username already taken"
	InvalidConstructorArguments Error = "invalid constructor arguments(some of them is probably nil)"
	NilLogger                   Error = "given logger is nil"
	NilRouteHandler             Error = "given route handler is nil"
	IncreaseAmountNotValid      Error = "increase amount not valid"
)
