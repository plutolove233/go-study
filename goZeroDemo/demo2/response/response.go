// Package response
/*
@Coding : utf-8
@time : 2022/12/1 21:05
@Author : yizhigopher
@Software : GoLand
*/
package response

import (
	"demo2/common/errorx"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type Body struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Response(w http.ResponseWriter, resp interface{}, err error) {
	var body Body

	if err != nil {
		e := err.(*errorx.CodeError)
		body.Code = e.Code
		body.Msg = e.Msg
	} else {
		body.Code = errorx.OK
		body.Msg = "OK"
		body.Data = resp
	}
	httpx.OkJson(w, body)
}
