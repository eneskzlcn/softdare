package comment

import "errors"

var (
	ErrInvalidCommentContent       = errors.New("invalid comment content")
	ErrInvalidPostID               = errors.New("invalid post id for comment")
	ErrCouldNotTakeUserFromContext = errors.New("could not take the user from ctx ")
)
