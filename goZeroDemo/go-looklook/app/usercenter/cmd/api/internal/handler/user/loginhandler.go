package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"go-looklook/app/usercenter/cmd/api/internal/logic/user"
	"go-looklook/app/usercenter/cmd/api/internal/svc"
	"go-looklook/app/usercenter/cmd/api/internal/types"
	"go-looklook/common/response" // import response which you create
	"net/http"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		response.Response(w, resp, err)

	}
}
