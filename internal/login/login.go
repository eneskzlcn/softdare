package login

import (
	"net/url"
	"time"
)

const DomainName = "login"

type LoginInput struct {
	Email    string
	Username *string
}

type loginPageData struct {
	Form url.Values
	Err  error
}

type CreateUserRequest struct {
	ID       string
	Email    string
	Username string
}

type User struct {
	ID        string
	Email     string
	Username  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserSessionData struct {
	Email    string
	Username string
}
