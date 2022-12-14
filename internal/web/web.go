package web

import (
	"github.com/eneskzlcn/softdare/internal/entity"
	"github.com/eneskzlcn/softdare/internal/util/convertutil"
	"net/http"
)

type CommonSessionData struct {
	IsLoggedIn bool `json:"is_logged_in"`
	User       entity.UserIdentity
}

func (h *Handler) CommonSessionDataFromRequest(r *http.Request) (isLoggedIn bool, user entity.UserIdentity) {
	if h.session.Exists(r, "user") {
		data := h.session.Get(r, "user")
		userData, err := convertutil.AnyTo[entity.UserIdentity](data)
		if err == nil {
			user = userData
			isLoggedIn = true
		}
	}
	return
}
