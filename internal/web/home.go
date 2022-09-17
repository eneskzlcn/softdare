package web

import (
	"context"
	"github.com/eneskzlcn/softdare/internal/entity"
	"github.com/eneskzlcn/softdare/internal/util/convertutil"
	"github.com/eneskzlcn/softdare/internal/util/postutil"
	"github.com/eneskzlcn/softdare/internal/util/timeutil"
	"net/http"
	"net/url"
)

type homeSessionData struct {
	IsLoggedIn      bool                `json:"is_logged_in"`
	User            entity.UserIdentity `json:"user"`
	CreatePostError error               `json:"create_post_error"`
	CreatePostForm  url.Values          `json:"create_post_form"`
}

type homeData struct {
	Session homeSessionData
	Posts   []entity.FormattedPost
}

func (h *Handler) ShowHome(w http.ResponseWriter, r *http.Request) {
	h.logger.Debugf("HOME SHOW HANDLER ACCEPTED A REQUEST")
	session := h.GetHomeSessionData(r)
	ctx := context.WithValue(r.Context(), "user", session.User)

	posts, err := h.service.GetFollowingUsersPosts(ctx, 5)
	if err != nil {
		h.logger.Error("oops getting following users posts from service")
		h.ShowOops(w, r, err, http.StatusFound)
		return
	}
	formattedPosts := postutil.FormatPosts(posts, timeutil.ToAgoFormatter)

	h.RenderHome(w, homeData{Session: session, Posts: formattedPosts}, http.StatusFound)
}

func (h *Handler) RenderHome(w http.ResponseWriter, data homeData, status int) {
	h.RenderPage("home", w, data, status)
}

func (h *Handler) GetHomeSessionData(r *http.Request) (out homeSessionData) {
	isLoggedIn, user := h.CommonSessionDataFromRequest(r)
	if isLoggedIn {
		out.IsLoggedIn = isLoggedIn
		out.User = user
		out.CreatePostError = h.session.PopError(r, "create-post-oops")

		if h.session.Exists(r, "create-post-form") {
			form := h.session.Pop(r, "create-post-form")
			formData, err := convertutil.AnyTo[url.Values](form)
			if err == nil {
				out.CreatePostForm = formData
			}
		}
		h.logger.Debugf("Session data exists for user with id %s", out.User.ID)
	}
	return out
}
