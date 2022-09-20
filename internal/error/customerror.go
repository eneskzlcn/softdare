package customerror

import "errors"

var (
	InvalidPostID                           = errors.New("invalid post ID")
	InvalidUsername                         = errors.New("invalid username")
	UsernameNotGiven                        = errors.New("username not given")
	InvalidEmail                            = errors.New("invalid email")
	InvalidUserID                           = errors.New("invalid user ID")
	CouldNotTakeUserFromContext             = errors.New("could not take the user from ctx")
	Unauthorized                            = errors.New("given unauthorized user")
	UserNotFound                            = errors.New("user not found")
	UsernameAlreadyTaken                    = errors.New("username already taken")
	InvalidConstructorArguments             = errors.New("invalid constructor arguments(some of them is probably nil)")
	NilLogger                               = errors.New("given logger is nil")
	AdjustmentNotValid                      = errors.New("not valid adjustment")
	UserNotInContext                        = errors.New("username not found in context")
	UserCanNotFollowItself                  = errors.New("a person can not follow itself")
	AlreadyFollowsTheUser                   = errors.New("already following user")
	UserFollowNotFound                      = errors.New("user follow not found")
	NotLoggedInUser                         = errors.New("user not logged in")
	UserAlreadyLikedTheComment              = errors.New("given user already liked the comment")
	UserAlreadyLikedThePost                 = errors.New("given user already liked the post")
	SearchUserWithFollowingOptionByUsername = errors.New("can not searched users with following option by username search criteria")
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
