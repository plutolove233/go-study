// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package services

import (
	"wire-demo/bindInterface/models"
	"wire-demo/bindInterface/services/impl"
)

// Injectors from wire.go:

func InitUserService(user models.User) *UserService {
	userServiceImpl := impl.NewUserServiceImpl(user)
	userService := NewUserService(userServiceImpl)
	return userService
}
