package web

import (
	"context"
	mux_router "github.com/eneskzlcn/mux-router"
	"github.com/golangcollege/sessions"
	"go.uber.org/zap"
	"net/http"
	"softdare/web/login"
	"sync"
)

type LoginService interface {
	Login(ctx context.Context, input login.LoginInput) (*login.User, error)
}
type Handler struct {
	logger       *zap.SugaredLogger
	session      *sessions.Session
	SessionKey   []byte
	loginService LoginService
	handler      http.Handler
	once         sync.Once
}

func NewHandler(logger *zap.SugaredLogger) *Handler {
	return &Handler{logger: logger}
}
func (h *Handler) init() {
	router := mux_router.New()
	router.HandleFunc(http.MethodGet, "/login", h.showLogin)
	h.handler = router
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.once.Do(h.init)
	h.handler.ServeHTTP(w, r)
}
