package engine

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrInvalidMethod = errors.New("invalid http method")
	ErrInvalidPath = errors.New("path must begin with '/'")
)


type (
	routeHandler map[string]http.HandlerFunc

	router struct {
		mux map[string]routeHandler
	}
)


func newRouter() *router {
	return &router{mux: make(map[string]routeHandler)}
}


func (r *router) registerRoute(route Route) error {
	if !r.validMethod(route.Method) {
		return ErrInvalidMethod
	}

	if len(route.Path) == 0 || route.Path[0] != '/' {
		return ErrInvalidPath
	}

	if _, ok := r.mux[route.Method]; !ok {
		r.mux[route.Method] = make(routeHandler)
	}

	r.mux[route.Method][route.Path] = route.Handler

	return nil
}

func (r *router) validMethod(method string) bool {
	switch method {
	case http.MethodGet, http.MethodDelete, http.MethodPost, http.MethodPatch,
		http.MethodPut:
		return true
	default:
		return false
	}
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if handler, ok := r.mux[req.Method]; ok {
		if h, ok := handler[req.URL.Path]; ok {
			h.ServeHTTP(w, req)
		} else {
			fmt.Fprintf(w, "no path: %s\n", req.URL.Path)
		}
	}
}
