package home

import (
	"encoding/gob"
	"fmt"
	muxRouter "github.com/eneskzlcn/mux-router"
	"github.com/eneskzlcn/softdare/internal/pkg"
	"go.uber.org/zap"
	"html/template"
	"net/http"
)

type Renderer interface {
	RenderTemplate(w http.ResponseWriter, template *template.Template, data any, statusCode int)
}
type SessionProvider interface {
	Exists(r *http.Request, key string) bool
	Get(r *http.Request, key string) any
}

func (h *Handler) init() error {
	gob.Register(UserSessionData{})
	template, err := pkg.ParseTemplate(DomainName)
	if err != nil {
		return err
	}
	h.homeTemplate = template
	return nil
}

type Handler struct {
	logger          *zap.SugaredLogger
	homeTemplate    *template.Template
	renderer        Renderer
	sessionProvider SessionProvider
}

func NewHandler(logger *zap.SugaredLogger, renderer Renderer, provider SessionProvider) *Handler {
	handler := Handler{logger: logger, renderer: renderer, sessionProvider: provider}
	if err := handler.init(); err != nil {
		logger.Error("Error occurred when initializing home handler ", zap.Error(err))
		return nil
	}
	return &handler
}

func (h *Handler) RegisterHandlers(router *muxRouter.Router) {
	router.HandleFunc(http.MethodGet, "/", h.Show)

}
func (h *Handler) Render(w http.ResponseWriter, data homeData, statusCode int) {
	h.logger.Debugf("RENDERING TEMPLATE %s", h.homeTemplate.Name())
	h.renderer.RenderTemplate(w, h.homeTemplate, data, statusCode)
}
func (h *Handler) Show(w http.ResponseWriter, r *http.Request) {
	h.logger.Debugf("HOME SHOW HANDLER ACCEPTED A REQUEST")
	session := sessionFromRequest(h.sessionProvider, r)
	h.Render(w, homeData{Session: session}, http.StatusOK)
}
func sessionFromRequest(session SessionProvider, r *http.Request) SessionData {
	var out SessionData
	if session.Exists(r, "user") {
		user := session.Get(r, "user")
		userData, _ := SessionDataFromAny(user)
		out.User = userData
		out.IsLoggedIn = true
		fmt.Printf("Session data exist for the user. Session data:%v\n", out)
	}
	return out
}
