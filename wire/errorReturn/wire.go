//go:build wireinject
// +build wireinject

// Package main
/*
@Coding : utf-8
@Time : 2023/2/9 15:56
@Author : yizhigopher
@Software : GoLand
*/
package main

import "github.com/google/wire"

func InitClient(config Config) (*Service, error) {
	wire.Build(NewAPIClient, NewService)
	return nil, nil
}
