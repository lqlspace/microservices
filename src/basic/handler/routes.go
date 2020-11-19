package handler

import (
	"net/http"

	"github.com/lqlspace/microservices/src/basic/cmd/api/svc"
	"github.com/lqlspace/microservices/src/engine"
)

func RegisterHandlers(e *engine.Engine, ctx *svc.ServiceContext) {
	e.AddRoutes([]engine.Route{
		{
			Method: http.MethodPost,
			Path: "/basic/elem/create",
			Handler:creatBasicHandler(ctx),
		},
		{
			Method: http.MethodGet,
			Path: "/basic/elem",
			Handler: getBasicListHandler(ctx),
		},
	})
}
