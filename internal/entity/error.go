package entity

import (
	"errors"
)

var (
	InvalidPostID               = errors.New("invalid post ID")
	InvalidUsername             = errors.New("invalid username")
	UsernameNotGiven            = errors.New("username not given")
	InvalidEmail                = errors.New("invalid email")
	InvalidUserID               = errors.New("invalid user ID")
	InvalidContent              = errors.New("invalid content")
	InvalidCommentContent       = errors.New("invalid comment content")
	NilRepository               = errors.New("given repository is nil")
	CouldNotTakeUserFromContext = errors.New("could not take the user from ctx")
	NilDatabase                 = errors.New("given database is nil")
	Unauthorized                = errors.New("given unauthorized user")
	NilSession                  = errors.New("given session is nil")
	NilService                  = errors.New("given service is nil")
	UserNotFound                = errors.New("user not found")
	UsernameAlreadyTaken        = errors.New("username already taken")
	InvalidConstructorArguments = errors.New("invalid constructor arguments(some of them is probably nil)")
	NilLogger                   = errors.New("given logger is nil")
	NilRouteHandler             = errors.New("given route handler is nil")
	AdjustmentNotValid          = errors.New("not valid adjustment")
	UserNotInContext            = errors.New("username not found in context")
	UserCanNotFollowItself      = errors.New("a person can not follow itself")
	AlreadyFollowsTheUser       = errors.New("already following user")
	UserFollowNotFound          = errors.New("user follow not found")
	NotValidLoggerEnvironment   = errors.New("not valid logger environment")
)

func IsLoginError(err error) bool {
	return errors.Is(err, UserNotFound) ||
		errors.Is(err, UsernameAlreadyTaken) || errors.Is(err, InvalidUsername) ||
		errors.Is(err, InvalidEmail) || errors.Is(err, UsernameNotGiven)
}
func IsLoginErrorStr(err string) bool {
	return err == UserNotFound.Error() ||
		err == UsernameAlreadyTaken.Error() || err == InvalidUsername.Error() ||
		err == InvalidEmail.Error() || err == UsernameNotGiven.Error()
}
