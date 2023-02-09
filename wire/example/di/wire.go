//go:build wireinject
// +build wireinject

// Package di
/*
@Coding : utf-8
@Time : 2023/2/9 22:18
@Author : yizhigopher
@Software : GoLand
*/
package di

import (
	"context"
	"example/db"
	"example/services"
	"github.com/google/wire"
	"io"
)

func NewService(c *db.Config, m *services.MailConfig) (*services.Service, error) {
	wire.Build(services.NewService, services.MailProvider, services.UserProvider, db.NewDB)
	return nil, nil
}

func InitGreeter(ctx context.Context, s []string, w io.Writer) (*services.Greeter, error) {
	wire.Build(services.GreeterProvider)
	return nil, nil
}
