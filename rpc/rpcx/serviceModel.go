// Package rpcx
/*
@Coding : utf-8
@Time : 2023/1/10 20:49
@Author : yizhigopher
@Software : GoLand
*/
package rpcx

import (
	"context"
	"fmt"
)

type Request struct {
	A int
	B int
}
type Response struct {
	C int
}

type Arith struct {
}

func (s *Arith) Add(c context.Context, request *Request, response *Response) error {
	response.C = request.A + request.B
	fmt.Printf("call %d + %d = %d\n", request.A, request.B, response.C)
	return nil
}

func (s *Arith) Mul(c context.Context, request *Request, response *Response) error {
	response.C = request.A * request.B
	fmt.Printf("call %d * %d = %d\n", request.A, request.B, response.C)
	return nil
}

type Greeter struct {
}

func (s *Greeter) Say(c context.Context, request *string, response *string) error {
	*response = fmt.Sprintf("hello %s\n", *request)
	return nil
}
