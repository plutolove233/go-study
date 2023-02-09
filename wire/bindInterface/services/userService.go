// Package services
/*
@Coding : utf-8
@Time : 2023/2/8 21:38
@Author : yizhigopher
@Software : GoLand
*/
package services

import (
	"github.com/google/wire"
	"wire-demo/bindInterface/models"
	"wire-demo/bindInterface/services/impl"
)

type UserRepository interface {
	GetUserById(id string) (*models.User, error)
}

type UserService struct {
	UserRepository
}

func NewUserService(userRepository UserRepository) *UserService {
	return &UserService{UserRepository: userRepository}
}

// UserServiceProvider says the binding between interface and implements
var UserServiceProvider = wire.NewSet(impl.NewUserServiceImpl, wire.Bind(new(UserRepository), new(*impl.UserServiceImpl)))
