// Package bindInterface
/*
@Coding : utf-8
@Time : 2023/2/8 21:35
@Author : yizhigopher
@Software : GoLand
*/
package main

import (
	"fmt"
	"wire-demo/bindInterface/models"
	"wire-demo/bindInterface/services"
)

func main() {
	userService := services.InitUserService(models.User{
		Id:   "12020180412",
		Name: "shy hao",
	})

	u, _ := userService.GetUserById("12")

	fmt.Println(u)
}
