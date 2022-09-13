package router

import "net/http"

type MethodHandlers map[string]http.HandlerFunc

type Router interface {
	Handle(pattern string, methodHandlers MethodHandlers)
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}
