package web

import (
	"go.uber.org/zap"
	"net/http"
	"net/url"
	"softdare/web/login"
)

var loginTemplate = parseTemplate("login.gohtml")

type loginData struct {
	Form url.Values
	Err  error
}

func (h *Handler) renderLogin(w http.ResponseWriter, data loginData, statusCode int) {
	h.renderTemplate(w, loginTemplate, data, statusCode)
}

func (h *Handler) showLogin(w http.ResponseWriter, r *http.Request) {
	h.renderLogin(w, loginData{}, http.StatusOK)
}

func (h *Handler) login(w http.ResponseWriter, r *http.Request) {
	h.logger.Infof("LOGIN HANDLER ACCEPTED A REQUEST WITH email %s", r.PostFormValue("email"))
	if err := r.ParseForm(); err != nil {
		h.logger.Debugf("error occurred when parsing request's form with email %s", r.PostFormValue("email"))
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	ctx := r.Context()
	inp := login.LoginInput{
		Email:    r.PostFormValue("email"),
		Username: extractFormValue(r.Form, "username"),
	}
	user, err := h.loginService.Login(ctx, inp)
	if err != nil {
		h.logger.Debug("could not login", zap.Error(err))
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	h.logger.Debugf("Successfully logged in user %v", user)
	h.session.Put(r, "user", *user)
	h.logger.Debugf("Session defined for the user %v", user)
	http.Redirect(w, r, "/", http.StatusFound)
}

func extractFormValue(form url.Values, key string) *string {
	if !form.Has(key) {
		return nil
	}
	s := form.Get(key)
	return &s
}
