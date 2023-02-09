// Package impl
/*
@Coding : utf-8
@Time : 2023/2/8 21:40
@Author : yizhigopher
@Software : GoLand
*/
package impl

import (
	"wire-demo/bindInterface/models"
)

// UserServiceImpl 具体实现了UserRepository接口
type UserServiceImpl struct {
	models.User
}

func NewUserServiceImpl(user models.User) *UserServiceImpl {
	return &UserServiceImpl{User: user}
}

func (u *UserServiceImpl) GetUserById(id string) (*models.User, error) {
	return &u.User, nil
}
