package server

import (
	"errors"
	muxRouter "github.com/eneskzlcn/mux-router"
	"go.uber.org/zap"
	"net/http"
	"sync"
)

type RouteHandler interface {
	RegisterHandlers(router *muxRouter.Router)
}
type Session interface {
	Enable(handler http.Handler) http.Handler
}
type Handler struct {
	logger          *zap.SugaredLogger
	sessionProvider Session
	handler         http.Handler
	once            sync.Once
}

func NewHandler(logger *zap.SugaredLogger, routeHandlers []RouteHandler, sessionProvider Session) (*Handler, error) {
	handler := Handler{logger: logger, sessionProvider: sessionProvider}
	router := muxRouter.New()
	for _, routeHandler := range routeHandlers {
		if routeHandler == nil {
			logger.Error("One of the given routeHandlers to the server handler is nil")
			return nil, errors.New("given handler is nil")
		}
		routeHandler.RegisterHandlers(router)
	}
	handler.handler = router
	handler.handler = sessionProvider.Enable(handler.handler)
	return &handler, nil
}
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.handler.ServeHTTP(w, r)
}
