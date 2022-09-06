package web

import (
	mux_router "github.com/eneskzlcn/mux-router"
	"go.uber.org/zap"
	"net/http"
	"sync"
)

type Handler struct {
	logger  *zap.SugaredLogger
	handler http.Handler
	once    sync.Once
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
