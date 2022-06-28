package handler

import (
	"account_gateway/tools"
	"account_gateway/tools/xhttp"
	"github.com/tal-tech/go-zero/rest/httpx"
	"net/http"

	"greet/internal/logic"
	"greet/internal/svc"
	"greet/internal/types"
)

func GreetHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			xhttp.ParamErrorResult(r, w, err)
			return
		}

		l := logic.NewGreetLogic(tools.SetMateDataContext(r), ctx)
		resp, err := l.Greet(req)
		xhttp.HttpResult(r, w, resp, err)
	}
}
