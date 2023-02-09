// Package models
/*
@Coding : utf-8
@Time : 2023/2/9 22:11
@Author : yizhigopher
@Software : GoLand
*/
package models

import "fmt"

type UserModel struct {
	Name string
	Age  int
}

func (u *UserModel) ToString() string {
	return fmt.Sprintf("name=%s\nage=%d\n", u.Name, u.Age)
}
