// Package errorx
/*
@Coding : utf-8
@time : 2022/12/2 11:16
@Author : yizhigopher
@Software : GoLand
*/
package errorx

const defaultCode = "1001"

type CodeError struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

type CodeErrorResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

func NewCodeError(code string, msg string) error {
	return &CodeError{Code: code, Msg: msg}
}

func NewDefaultError(msg string) error {
	return NewCodeError(defaultCode, msg)
}

func (e *CodeError) Error() string {
	return e.Msg
}

//
//func (e *CodeError) Data() *CodeErrorResponse {
//	return &CodeErrorResponse{
//		Code: e.Code,
//		Msg:  e.Msg,
//	}
//}
