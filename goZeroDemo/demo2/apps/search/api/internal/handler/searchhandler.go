package handler

import (
	"demo2/apps/search/api/internal/logic"
	"demo2/apps/search/api/internal/svc"
	"demo2/apps/search/api/internal/types"
	"demo2/response" // import response which you create
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func searchHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SearchReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewSearchLogic(r.Context(), svcCtx)
		resp, err := l.Search(&req)
		response.Response(w, resp, err)

	}
}
