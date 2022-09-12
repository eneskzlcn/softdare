package home

import (
	"context"
	"encoding/gob"
	"github.com/eneskzlcn/softdare/internal/core/logger"
	"github.com/eneskzlcn/softdare/internal/oops"
	"github.com/eneskzlcn/softdare/internal/pkg"
	"github.com/nicolasparada/go-mux"
	"html/template"
	"net/http"
)

type Renderer interface {
	RenderTemplate(w http.ResponseWriter, template *template.Template, data any, statusCode int)
}
type SessionProvider interface {
	Exists(r *http.Request, key string) bool
	Get(r *http.Request, key string) any
	GetString(r *http.Request, key string) string
	PopError(r *http.Request, key string) error
	Pop(r *http.Request, key string) any
}

type HomeService interface {
	GetPosts(context.Context) ([]Post, error)
}

type Handler struct {
	logger          logger.Logger
	homeTemplate    *template.Template
	renderer        Renderer
	service         HomeService
	sessionProvider SessionProvider
}

func NewHandler(logger logger.Logger, renderer Renderer, provider SessionProvider, service HomeService) *Handler {
	if logger == nil {
		return nil
	}
	if renderer == nil {
		logger.Error(ErrRendererNil)
		return nil
	}
	if provider == nil {
		logger.Error(ErrSessionProviderNil)
		return nil
	}
	if service == nil {
		logger.Error("given service is nil")
		return nil
	}
	handler := Handler{logger: logger, renderer: renderer, sessionProvider: provider, service: service}
	handler.init()
	return &handler
}

func (h *Handler) init() {
	gob.Register(UserSessionData{})
	h.homeTemplate = pkg.ParseTemplate("home.gohtml")
}

func (h *Handler) Render(w http.ResponseWriter, data homeData, statusCode int) {
	h.logger.Debugf("RENDERING TEMPLATE %s", h.homeTemplate.Name())
	h.renderer.RenderTemplate(w, h.homeTemplate, data, statusCode)
}
func (h *Handler) Show(w http.ResponseWriter, r *http.Request) {
	h.logger.Debugf("HOME SHOW HANDLER ACCEPTED A REQUEST")
	session := sessionDataFromRequest(h.sessionProvider, r)
	posts, err := h.service.GetPosts(r.Context())
	if err != nil {
		h.logger.Error("oops getting posts from service")
		oops.RenderPage(h.renderer, h.logger, h.sessionProvider, r, w, err, http.StatusFound)
		return
	}
	h.Render(w, homeData{Session: session, Posts: posts}, http.StatusOK)
}

func (h *Handler) RegisterHandlers(router *mux.Router) {
	router.Handle("/", mux.MethodHandler{
		http.MethodGet: h.Show,
	})
}
