package web

import "net/http"

var homeTemplate = parseTemplate("home.gohtml")

type homeData struct {
	Session Session
}

func (h *Handler) renderHome(w http.ResponseWriter, data homeData, statusCode int) {
	h.renderTemplate(w, homeTemplate, data, statusCode)
}
func (h *Handler) showHome(w http.ResponseWriter, r *http.Request) {
	session := h.sessionFromRequest(r)
	h.renderHome(w, homeData{Session: session}, http.StatusOK)
}
