package web

import (
	"net/http"
	"net/url"
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
