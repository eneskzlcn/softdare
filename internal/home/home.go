package home

import (
	"github.com/eneskzlcn/softdare/internal/core/logger"
	"github.com/eneskzlcn/softdare/internal/core/session"
	"github.com/eneskzlcn/softdare/internal/entity"
	convertionUtil "github.com/eneskzlcn/softdare/internal/util/convertion"
	sessionUtil "github.com/eneskzlcn/softdare/internal/util/session"
	"net/http"
	"net/url"
)

const DomainName = "home"

type sessionData struct {
	IsLoggedIn      bool                   `json:"is_logged_in"`
	User            entity.UserSessionData `json:"user"`
	CreatePostError error                  `json:"create_post_error"`
	CreatePostForm  url.Values             `json:"create_post_form"`
}

type pageData struct {
	Session sessionData
	Posts   []entity.FormattedPost
}

func sessionDataFromRequest(session session.Session, r *http.Request, logger logger.Logger) (out sessionData) {
	isLoggedIn, user := sessionUtil.GeneralSessionDataFromRequest(logger, session, r)
	if isLoggedIn {
		out.IsLoggedIn = isLoggedIn
		out.User = user
		out.CreatePostError = session.PopError(r, "create-post-oops")

		if session.Exists(r, "create-post-form") {
			form := session.Pop(r, "create-post-form")
			formData, err := convertionUtil.AnyToGivenType[url.Values](form)
			if err == nil {
				out.CreatePostForm = formData
			}
		}
		logger.Debugf("Session data exists for user with id %s", out.User.ID)
	}
	return out
}
