package handler

import (
	"fmt"
	"net/http"

	"github.com/lqlspace/microservices/src/basic/cmd/api/svc"
)

func getBasicListHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "basic list response\n")
	}
}
