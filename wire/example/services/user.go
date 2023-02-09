// Package services
/*
@Coding : utf-8
@Time : 2023/2/9 22:07
@Author : yizhigopher
@Software : GoLand
*/
package services

import (
	"database/sql"
	"example/models"
	"github.com/google/wire"
	"log"
)

type UserRepo interface {
	AddUser(model models.UserModel)
}

type UserService struct {
	*sql.DB
}

func (u *UserService) AddUser(user models.UserModel) {
	log.Println("add user=", user.ToString())
}

func NewUserService(DB *sql.DB) *UserService {
	return &UserService{DB: DB}
}

var UserProvider = wire.NewSet(NewUserService, wire.Bind(new(UserRepo), new(*UserService)))
