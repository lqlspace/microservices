package engine

import (
	"fmt"
	"net/http"

	"github.com/lqlspace/microservices/src/basic/cmd/api/config"
)

type (
	Engine struct {
		router *router
		server *http.Server
	}
)


func NewEngine(c config.Config) (*Engine, error) {
	e :=  &Engine{
		router: newRouter(),
		server: &http.Server{
			Addr: fmt.Sprintf("%s:%d", c.Host, c.Port),
		},
	}

	return e, nil
}


func (e *Engine) Start() {
	e.server.Handler = e.router
	e.server.ListenAndServe()
}

func (e *Engine) AddRoutes(routes []Route) {
	for _, item := range routes {
		e.router.registerRoute(item)
	}
}
