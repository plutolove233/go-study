// Package services
/*
@Coding : utf-8
@Time : 2023/4/8 22:14
@Author : yizhigopher
@Software : GoLand
*/
package services

import (
	v3 "buf/api/v3"
	context "context"
)

type BookHandler struct {
	v3.UnimplementedBookServiceServer
}

func (b BookHandler) CreateBook(ctx context.Context, req *v3.CreateBookReq) (*v3.CreateBookResp, error) {
	return &v3.CreateBookResp{
		Book: &v3.Book{
			Name:   "xxx",
			Author: req.GetAuthor(),
			Price:  110,
		},
	}, nil
}
