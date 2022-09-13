package post

import (
	"github.com/eneskzlcn/softdare/internal/core/logger"
	"github.com/eneskzlcn/softdare/internal/core/session"
	"github.com/eneskzlcn/softdare/internal/entity"
	"github.com/eneskzlcn/softdare/internal/util/convertion"
	sessionUtil "github.com/eneskzlcn/softdare/internal/util/session"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type CreatePostRequest struct {
	ID      string
	UserID  string
	Content string
}
type CreatePostInput struct {
	Content string
}

func (c *CreatePostInput) Prepare() {
	c.Content = strings.TrimSpace(c.Content)
	c.Content = strings.ReplaceAll(c.Content, "\n\n", "\n")
	c.Content = strings.ReplaceAll(c.Content, "  ", " ")
}

func (c *CreatePostInput) Validate() error {
	return validation.Validate(c.Content, validation.Length(2, 1000))
}

type CreatePostResponse struct {
	ID        string
	CreatedAt time.Time
}

const userContextKey = "user"

type postData struct {
	Session  sessionData               `json:"session"`
	Post     entity.FormattedPost      `json:"post"`
	Comments []entity.FormattedComment `json:"comments"`
}
type sessionData struct {
	IsLoggedIn         bool
	User               entity.UserSessionData
	CreateCommentForm  url.Values
	CreateCommentError error
}

func sessionDataFromRequest(session session.Session, r *http.Request, logger logger.Logger) (out sessionData) {
	isLoggedIn, user := sessionUtil.GeneralSessionDataFromRequest(logger, session, r)
	if isLoggedIn {
		out.IsLoggedIn = isLoggedIn
		out.User = user
	}
	if session.Exists(r, "create-comment-error") {
		out.CreateCommentError = session.PopError(r, "create-comment-error")
	}
	if session.Exists(r, "create-comment-form") {
		form := session.Get(r, "create-comment-form")
		urlForm, err := convertion.AnyToGivenType[url.Values](form)
		if err == nil {
			out.CreateCommentForm = urlForm
		}
	}
	return
}
