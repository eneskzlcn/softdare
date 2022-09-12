package home

import (
	"fmt"
	"github.com/eneskzlcn/softdare/internal/core/logger"
	"github.com/eneskzlcn/softdare/internal/core/session"
	convertionUtil "github.com/eneskzlcn/softdare/internal/util/convertion"
	sessionUtil "github.com/eneskzlcn/softdare/internal/util/session"
	"net/http"
	"net/url"
)

const DomainName = "home"

type UserSessionData struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type SessionData struct {
	IsLoggedIn      bool            `json:"is_logged_in"`
	User            UserSessionData `json:"user"`
	CreatePostError error           `json:"create_post_error"`
	CreatePostForm  url.Values      `json:"create_post_form"`
}

type pageData struct {
	Session SessionData
	Posts   []Post
}
type Post struct {
	ID           string
	CreatedAt    string
	Content      string
	CommentCount int
	Username     string
}

func sessionDataFromRequest(session session.Session, r *http.Request, logger logger.Logger) SessionData {
	var out SessionData
	generalSessionData := sessionUtil.GeneralSessionDataFromRequest(logger, session, r)
	if generalSessionData.IsLoggedIn {
		out.IsLoggedIn = generalSessionData.IsLoggedIn
		out.User.ID = generalSessionData.User.ID
		out.User.Email = generalSessionData.User.Email
		out.User.Username = generalSessionData.User.Email

		out.CreatePostError = session.PopError(r, "create-post-oops")
		if session.Exists(r, "create-post-form") {
			form := session.Pop(r, "create-post-form")
			formData, err := convertionUtil.AnyToGivenType[url.Values](form)
			if err == nil {
				out.CreatePostForm = formData
			}
		}
		fmt.Printf("Session data exist for the user. Session data:%v\n", out)
	}
	return out
}
