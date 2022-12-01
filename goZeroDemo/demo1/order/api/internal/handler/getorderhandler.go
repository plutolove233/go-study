package handler

import (
	"demo1/order/api/internal/logic"
	"demo1/order/api/internal/svc"
	"demo1/order/api/internal/types"
	"demo1/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func getOrderHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OrderReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetOrderLogic(r.Context(), ctx)
		resp, err := l.GetOrder(&req)
		response.Response(w, resp, err) //②

	}
}
