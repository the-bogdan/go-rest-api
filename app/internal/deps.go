package internal

import "net/http"

type Router interface {
	HandlerFunc(method, path string, handler http.HandlerFunc)
	Handler(method, path string, handler http.Handler)
}

type Handler interface {
	Register(router Router)
}
