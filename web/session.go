package web

import (
	"net/http"
	"softdare/web/login"
)

type Session struct {
	IsLoggedIn bool
	User       login.User
}

func (h *Handler) sessionFromRequest(r *http.Request) Session {
	var out Session
	if h.session.Exists(r, "user") {
		user, ok := h.session.Get(r, "user").(login.User)
		out.IsLoggedIn = ok
		out.User = user
	}
	return out
}
