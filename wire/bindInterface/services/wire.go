//go:build wireinject
// +build wireinject

// Package services
/*
@Coding : utf-8
@Time : 2023/2/8 23:27
@Author : yizhigopher
@Software : GoLand
*/
package services

import (
	"github.com/google/wire"
	"wire-demo/bindInterface/models"
)

func InitUserService(user models.User) *UserService {
	wire.Build(NewUserService, UserServiceProvider)
	return nil
}
