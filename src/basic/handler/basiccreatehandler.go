package handler

import (
	"fmt"
	"net/http"

	"github.com/lqlspace/microservices/src/basic/cmd/api/svc"
)

func creatBasicHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "basic create response\n")
	}
}
