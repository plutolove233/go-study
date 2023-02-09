//go:build wireinject
// +build wireinject

// Package main
/*
@Coding : utf-8
@Time : 2023/2/9 16:29
@Author : yizhigopher
@Software : GoLand
*/
package main

import "github.com/google/wire"

func InitFileReader(path string) (*FileReader, func(), error) {
	wire.Build(NewFileReader)
	return nil, nil, nil
}
