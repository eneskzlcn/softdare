package login

import (
	"github.com/eneskzlcn/softdare/internal/core/logger"
	"github.com/eneskzlcn/softdare/internal/core/session"
	"github.com/eneskzlcn/softdare/internal/entity"
	sessionUtil "github.com/eneskzlcn/softdare/internal/util/session"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"net/http"
	"net/url"
)

const DomainName = "login"

type Input struct {
	Email    string
	Username *string
}

type loginPageData struct {
	Form    url.Values
	Err     error
	Session sessionData
}
type sessionData struct {
	User       entity.UserSessionData
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

func sessionDataFromRequest(session session.Session, r *http.Request, logger logger.Logger) (out sessionData) {
	isLoggedIn, user := sessionUtil.GeneralSessionDataFromRequest(logger, session, r)
	if isLoggedIn {
		out.IsLoggedIn = isLoggedIn
		out.User = user
	}
	return
}
