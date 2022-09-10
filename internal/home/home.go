package home

import (
	"fmt"
	convertionUtil "github.com/eneskzlcn/softdare/internal/util/convertion"
	"github.com/mvdan/xurls"
	"html/template"
	"net/http"
	"net/url"
	"time"
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

func sessionFromRequest(session SessionProvider, r *http.Request) SessionData {
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

func FormatPostTime(createdAt time.Time) string {
	durationBetweenNow := time.Now().Sub(createdAt)
	durationMinutes := int(durationBetweenNow.Minutes())
	if durationMinutes < 60 && durationMinutes > 0 {
		return fmt.Sprintf("%dm ago", durationMinutes)
	} else if durationMinutes == 0 {
		return "Just Now"
	} else if durationMinutes >= 60 && durationMinutes < 1440 {
		return fmt.Sprintf("%dh ago", durationMinutes/60)
	} else if durationMinutes >= 1440 && durationMinutes < 10080 {
		return fmt.Sprintf("%dd ago", durationMinutes/1440)
	} else if durationMinutes >= 10080 && durationMinutes < 40320 {
		return fmt.Sprintf("%dw ago", durationMinutes/10080)
	} else if durationMinutes >= 40320 && durationMinutes < 483840 {
		return fmt.Sprintf("%dmonths ago", durationMinutes/40320)
	} else if durationMinutes >= 483840 {
		return fmt.Sprintf("%dyears ago", durationMinutes/483840)
	}
	return ""
}
func FormatPostContent(content string) template.HTML {
	content = template.HTMLEscapeString(content)
	return template.HTML(xurls.Relaxed.
		ReplaceAllString(content, `<a href ="$0" target="_blank" rel="noopener noreferror">$0</a>`))
}
