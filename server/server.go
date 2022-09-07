package server

import (
	"github.com/eneskzlcn/softdare/internal/config"
	"net/http"
)

type RootHandler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}
type Server struct {
	server  *http.Server
	handler Handler
}

func New(config config.Server, handler RootHandler) *Server {
	server := Server{}
	server.server = &http.Server{Addr: config.Address, Handler: handler}
	return &server
}
func (s *Server) ListenAndServe() error {
	return s.server.ListenAndServe()
}

func (s *Server) Close() error {
	return s.server.Close()
}
