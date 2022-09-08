package login

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
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

func (c *CreateUserRequest) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.Email, is.Email),
		validation.Field(&c.Username, is.Alphanumeric, validation.Length(5, 12)),
	)
}

type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserSessionData struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}
