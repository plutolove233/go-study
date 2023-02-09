// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"context"
	"example/db"
	"example/services"
	"io"
)

// Injectors from wire.go:

func NewService(c *db.Config, m *services.MailConfig) (*services.Service, error) {
	mailSender := services.NewMailSender(m)
	sqlDB, err := db.NewDB(c)
	if err != nil {
		return nil, err
	}
	userService := services.NewUserService(sqlDB)
	service := services.NewService(mailSender, userService)
	return service, nil
}

func InitGreeter(ctx context.Context, s []string, w io.Writer) (*services.Greeter, error) {
	options := &services.Options{
		Message: s,
		Writer:  w,
	}
	greeter, err := services.NewGreeter(options)
	if err != nil {
		return nil, err
	}
	return greeter, nil
}
