package server

import (
	"encoding/gob"
	"github.com/eneskzlcn/softdare/internal/core/logger"
	"github.com/eneskzlcn/softdare/internal/core/router"
	"net/http"
	"net/url"
	"sync"
)

type RouteHandler interface {
	RegisterHandlers(router router.Router)
}
type Session interface {
	Enable(handler http.Handler) http.Handler
}
type Handler struct {
	logger  logger.Logger
	session Session
	handler http.Handler
	once    sync.Once
}

func NewHandler(logger logger.Logger, routeHandlers []RouteHandler, session Session) (*Handler, error) {
	if logger == nil {
		return nil, ErrLoggerNil
	}
	if session == nil {
		return nil, ErrSessionProviderNil
	}
	handler := Handler{logger: logger, session: session}
	router := router.NewMuxRouterAdapter()
	for _, routeHandler := range routeHandlers {
		if routeHandler == nil {
			logger.Error("One of the given routeHandlers to the server handler is nil")
			return nil, ErrGivenRouteHandlerNil
		}
		routeHandler.RegisterHandlers(router)
	}
	handler.handler = router
	handler.handler = session.Enable(handler.handler)
	gob.Register(url.Values{})

	return &handler, nil
}
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.handler.ServeHTTP(w, r)
}
