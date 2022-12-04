// Package errorx
/*
@Coding : utf-8
@time : 2022/12/2 22:17
@Author : yizhigopher
@Software : GoLand
*/
package errorx

const defaultCode = "1001"

type CodeError struct {
	code string `json:"code"`
	msg  string `json:"msg"`
}

func (c *CodeError) Code() string {
	return c.code
}

func (c *CodeError) Msg() string {
	return c.msg
}

func NewErrCodeMsg(code string, msg string) error {
	return &CodeError{
		code: code,
		msg:  msg,
	}
}

func NewErrCode(code string) error {
	return &CodeError{
		code: code,
		msg:  Code2Msg(code),
	}
}

func NewErrorMsg(msg string) error {
	return &CodeError{
		code: CommonError,
		msg:  msg,
	}
}

func (e *CodeError) Error() string {
	return e.msg
}
