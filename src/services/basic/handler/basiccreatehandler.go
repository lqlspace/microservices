package handler

import (
	"fmt"
	"net/http"

	"github.com/lqlspace/microservices/src/core/httpx"
	"github.com/lqlspace/microservices/src/services/basic/cmd/api/svc"
	"github.com/lqlspace/microservices/src/services/basic/logic"
)

func creatBasicHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req logic.BasicCreateRequest
		if err := httpx.Parse(r, &req, nil); err != nil {
			fmt.Fprintf(w, "parse err: %s\n", err)
		}
		l := logic.NewBasicCreateLogic()
		l.CreateBasicElem(&req)
	}
}
