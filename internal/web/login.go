package web

import (
	"github.com/eneskzlcn/softdare/internal/entity"
	"net/http"
	"net/url"
)

type loginPageData struct {
	Form    url.Values
	Err     error
	Session CommonSessionData
}

func (h *Handler) ShowLogin(w http.ResponseWriter, r *http.Request) {
	sessionData := h.GetLoginSessionData(r)
	h.RenderLogin(w, loginPageData{Session: sessionData}, http.StatusOK)
}

func (h *Handler) RenderLogin(w http.ResponseWriter, data loginPageData, status int) {
	h.RenderPage("login", w, data, status)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	h.logger.Infof("LOGIN HANDLER ACCEPTED A REQUEST WITH email %s", r.PostFormValue("email"))
	if err := r.ParseForm(); err != nil {
		h.logger.Debugf("oops occurred when parsing request's form with email %s", r.PostFormValue("email"))
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	ctx := r.Context()
	email := r.PostFormValue("email")
	username := r.PostFormValue("username")
	user, err := h.service.Login(ctx, email, &username)
	if err != nil {
		h.logger.Debug("could not login", h.logger.ErrorModifier(err))
		if entity.IsLoginError(err) {
			h.RenderLogin(w, loginPageData{Form: r.PostForm, Err: err}, http.StatusBadRequest)
			return
		}
		http.Error(w, "could not login, err", http.StatusInternalServerError)
		return
	}
	h.logger.Debugf("Successfully logged in user %v", user)
	h.session.Put(r, "user", entity.UserIdentity{ID: user.ID, Email: user.Email, Username: user.Username})
	http.Redirect(w, r, "/", http.StatusFound)
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	h.session.Remove(r, "user")
	http.Redirect(w, r, "/", http.StatusFound)
}

func (h *Handler) GetLoginSessionData(r *http.Request) (out CommonSessionData) {
	isLoggedIn, user := h.CommonSessionDataFromRequest(r)
	if isLoggedIn {
		out.IsLoggedIn = isLoggedIn
		out.User = user
	}
	return
}
