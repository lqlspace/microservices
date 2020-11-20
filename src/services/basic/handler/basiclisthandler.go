package handler

import (
	"fmt"
	"net/http"

	"github.com/lqlspace/microservices/src/core/httpx"
	"github.com/lqlspace/microservices/src/services/basic/cmd/api/svc"
	"github.com/lqlspace/microservices/src/services/basic/logic"
)

func getBasicListHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req logic.BasicListRequest
		params := make(map[string]string)
		if err := httpx.Parse(r, &req, params); err != nil {
			fmt.Fprintf(w, "parse err: %s\n", err)
			return
		}
		l := logic.NewBasicListLogic()
		if val, ok := params["type"]; ok {
			req.Type = val
		}

		rsp, err := l.GetBasicListLogic(&req)
		if err != nil {
			fmt.Fprintf(w, "get basic list err := %s\n", err)
		} else {
			fmt.Fprintf(w, fmt.Sprintf("%s", rsp))
		}
	}
}
