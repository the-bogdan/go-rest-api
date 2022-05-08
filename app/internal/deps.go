package internal

import "net/http"

type Router interface {
	HandlerFunc(method, path string, handler http.HandlerFunc)
}
