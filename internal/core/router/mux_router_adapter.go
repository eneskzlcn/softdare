package router

import (
	"github.com/nicolasparada/go-mux"
	"net/http"
)

type MuxRouterAdapter struct {
	router *mux.Router
}

func NewMuxRouterAdapter() *MuxRouterAdapter {
	router := mux.NewRouter()
	return &MuxRouterAdapter{router: router}
}
func (m *MuxRouterAdapter) Handle(pattern string, method string, handler http.HandlerFunc) {
	m.router.Handle(pattern, mux.MethodHandler{
		method: handler,
	})
}

func (m *MuxRouterAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.router.ServeHTTP(w, r)
}
