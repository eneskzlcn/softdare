package web

import (
	"github.com/eneskzlcn/softdare/internal/entity"
	"net/http"
)

type oopsSessionData struct {
	IsLoggedIn bool
	User       entity.UserIdentity
}
type oopsData struct {
	Err     error
	Session oopsSessionData
}

func (h *Handler) ShowOops(w http.ResponseWriter, r *http.Request, err error, statusCode int) {
	h.logger.Error(err)
	isLoggedIn, user := h.CommonSessionDataFromRequest(r)
	sessionData := oopsSessionData{
		IsLoggedIn: isLoggedIn,
		User:       user,
	}
	h.RenderPage("oops", w, oopsData{Session: sessionData, Err: err}, statusCode)
}
