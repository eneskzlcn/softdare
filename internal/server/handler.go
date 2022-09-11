package server

import (
	"encoding/gob"
	"github.com/nicolasparada/go-mux"
	"go.uber.org/zap"
	"net/http"
	"net/url"
	"sync"
)

type RouteHandler interface {
	RegisterHandlers(router *mux.Router)
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
	if logger == nil {
		return nil, ErrLoggerNil
	}
	if sessionProvider == nil {
		return nil, ErrSessionProviderNil
	}
	handler := Handler{logger: logger, sessionProvider: sessionProvider}
	router := mux.NewRouter()
	for _, routeHandler := range routeHandlers {
		if routeHandler == nil {
			logger.Error("One of the given routeHandlers to the server handler is nil")
			return nil, ErrGivenRouteHandlerNil
		}
		routeHandler.RegisterHandlers(router)
	}
	handler.handler = router
	handler.handler = sessionProvider.Enable(handler.handler)
	gob.Register(url.Values{})

	return &handler, nil
}
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.handler.ServeHTTP(w, r)
}
