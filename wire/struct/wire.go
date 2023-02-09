//go:build wireinject
// +build wireinject

// Package main
/*
@Coding : utf-8
@Time : 2023/2/9 20:05
@Author : yizhigopher
@Software : GoLand
*/
package main

import "github.com/google/wire"

func InitStruct() *FooBar {
	wire.Build(MySet)
	return nil
}
