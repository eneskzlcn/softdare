package web

import (
	"context"
	"encoding/gob"
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
	sessionKey   []byte
	loginService LoginService
	handler      http.Handler
	once         sync.Once
}

func NewHandler(logger *zap.SugaredLogger, loginService LoginService, sessionKey []byte) *Handler {
	return &Handler{logger: logger, loginService: loginService, sessionKey: sessionKey}
}
func (h *Handler) init() {
	router := mux_router.New()
	router.HandleFunc(http.MethodGet, "/login", h.showLogin)
	router.HandleFunc(http.MethodPost, "/login", h.login)
	router.HandleFunc(http.MethodGet, "/", h.showHome)
	router.HandleFunc(http.MethodPost, "/logout", h.logout)
	h.session = sessions.New(h.sessionKey)
	gob.Register(login.User{})
	h.handler = router
	h.handler = h.session.Enable(h.handler)
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.once.Do(h.init)
	h.handler.ServeHTTP(w, r)
}
