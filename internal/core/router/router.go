package router

import "net/http"

type Router interface {
	Handle(pattern, method string, handler http.HandlerFunc)
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}
