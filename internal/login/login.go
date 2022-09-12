package login

import (
	"github.com/eneskzlcn/softdare/internal/core/logger"
	"github.com/eneskzlcn/softdare/internal/core/session"
	sessionUtil "github.com/eneskzlcn/softdare/internal/util/session"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"net/http"
	"net/url"
	"time"
)

const DomainName = "login"

type LoginInput struct {
	Email    string
	Username *string
}

type loginPageData struct {
	Form    url.Values
	Err     error
	Session SessionData
}
type SessionData struct {
	User       UserSessionData
	IsLoggedIn bool
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

func sessionDataFromRequest(session session.Session, r *http.Request, logger logger.Logger) SessionData {
	var out SessionData
	generalSession := sessionUtil.GeneralSessionDataFromRequest(logger, session, r)
	if generalSession.IsLoggedIn {
		out.IsLoggedIn = generalSession.IsLoggedIn
		out.User.ID = generalSession.User.ID
		out.User.Email = generalSession.User.Email
		out.User.Username = generalSession.User.Username
	}
	return out
}
