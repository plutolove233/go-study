//go:build wireinject
// +build wireinject

// Package main
/*
@Coding : utf-8
@Time : 2023/2/8 20:54
@Author : yizhigopher
@Software : GoLand
*/
package main

import "github.com/google/wire"

func initEvent(msg string) Event {
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}
}
