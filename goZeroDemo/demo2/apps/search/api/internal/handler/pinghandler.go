package handler

import (
	"demo2/apps/search/api/internal/logic"
	"demo2/apps/search/api/internal/svc"
	"demo2/response" // import response which you create
	"net/http"
)

func pingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewPingLogic(r.Context(), svcCtx)
		err := l.Ping()
		response.Response(w, nil, err)

	}
}
