package home

import (
	"context"
	"encoding/gob"
	coreTemplate "github.com/eneskzlcn/softdare/internal/core/html/template"
	"github.com/eneskzlcn/softdare/internal/core/logger"
	"github.com/eneskzlcn/softdare/internal/core/router"
	"github.com/eneskzlcn/softdare/internal/core/session"
	"github.com/eneskzlcn/softdare/internal/entity"
	"github.com/eneskzlcn/softdare/internal/oops"
	"github.com/eneskzlcn/softdare/internal/pkg"
	"html/template"
	"net/http"
)

type HomeService interface {
	GetPosts(context.Context) ([]entity.FormattedPost, error)
}

type Handler struct {
	logger       logger.Logger
	homeTemplate *template.Template
	service      HomeService
	session      session.Session
}

func NewHandler(logger logger.Logger, session session.Session, service HomeService) *Handler {
	if logger == nil {
		return nil
	}
	if session == nil {
		logger.Error(ErrSessionProviderNil)
		return nil
	}
	if service == nil {
		logger.Error("given service is nil")
		return nil
	}
	handler := Handler{logger: logger, session: session, service: service}
	handler.init()
	return &handler
}

func (h *Handler) init() {
	gob.Register(entity.UserSessionData{})
	h.homeTemplate = pkg.ParseTemplate("home.gohtml")
}

func (h *Handler) Render(w http.ResponseWriter, data pageData, statusCode int, renderFn coreTemplate.RenderFn) {
	h.logger.Debugf("RENDERING TEMPLATE %s", h.homeTemplate.Name())
	renderFn(h.logger, w, h.homeTemplate, data, statusCode)
}
func (h *Handler) Show(w http.ResponseWriter, r *http.Request) {
	h.logger.Debugf("HOME SHOW HANDLER ACCEPTED A REQUEST")
	session := sessionDataFromRequest(h.session, r, h.logger)
	posts, err := h.service.GetPosts(r.Context())
	if err != nil {
		h.logger.Error("oops getting posts from service")
		oops.RenderPage(h.logger, h.session, r, w, err, http.StatusFound, coreTemplate.Render)
		return
	}
	h.Render(w, pageData{Session: session, Posts: posts}, http.StatusOK, coreTemplate.Render)
}

func (h *Handler) RegisterHandlers(_router router.Router) {
	_router.Handle("/", router.MethodHandlers{
		http.MethodGet: h.Show,
	})
}
