package server

import (
	"fmt"
	"github.com/eneskzlcn/softdare/internal/config"
	"github.com/eneskzlcn/softdare/internal/core/logger"
	"net/http"
)

type RootHandler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

type Server struct {
	server  *http.Server
	handler RootHandler
	logger  logger.Logger
}

func New(config config.Server, handler RootHandler, logger logger.Logger) *Server {
	server := Server{}
	if logger == nil {
		fmt.Printf("Given logger to server is nil\n")
		return nil
	}
	if handler == nil {
		logger.Error("Given root handler to server is nil")
		return nil
	}
	server.server = &http.Server{Addr: config.Address, Handler: handler}
	return &server
}
func (s *Server) ListenAndServe() error {
	return s.server.ListenAndServe()
}

func (s *Server) Close() error {
	return s.server.Close()
}
