package handler

import (
	"demo2/apps/user/api/internal/logic"
	"demo2/apps/user/api/internal/svc"
	"demo2/apps/user/api/internal/types"
	"demo2/response" // import response which you create
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func loginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		response.Response(w, resp, err)

	}
}
