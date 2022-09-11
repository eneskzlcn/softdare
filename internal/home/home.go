package home

import (
	"fmt"
	convertionUtil "github.com/eneskzlcn/softdare/internal/util/convertion"
	"html/template"
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

type homeData struct {
	Session SessionData
	Posts   []Post
}
type Post struct {
	ID        string
	CreatedAt string
	Content   template.HTML
	Username  string
}

func sessionDataFromRequest(session SessionProvider, r *http.Request) SessionData {
	var out SessionData
	if session.Exists(r, "user") {
		user := session.Get(r, "user")
		userData, err := convertionUtil.AnyToGivenType[UserSessionData](user)
		if err == nil {
			out.User = userData
			out.IsLoggedIn = true
		}
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
