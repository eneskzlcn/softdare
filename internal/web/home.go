package web

import (
	"github.com/eneskzlcn/softdare/internal/entity"
	convertionUtil "github.com/eneskzlcn/softdare/internal/util/convertion"
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

func (h *Handler) ShowHome(w http.ResponseWriter, req *http.Request) {
	h.logger.Debugf("HOME SHOW HANDLER ACCEPTED A REQUEST")
	session := h.GetHomeSessionData(req)
	posts, err := h.service.GetFormattedPosts(req.Context(), "")
	if err != nil {
		h.logger.Error("oops getting posts from service")
		h.ShowOops(w, req, err, http.StatusFound)
		return
	}
	h.RenderHome(w, homeData{Session: session, Posts: posts}, http.StatusFound)
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
			formData, err := convertionUtil.AnyToGivenType[url.Values](form)
			if err == nil {
				out.CreatePostForm = formData
			}
		}
		h.logger.Debugf("Session data exists for user with id %s", out.User.ID)
	}
	return out
}
