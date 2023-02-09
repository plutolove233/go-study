// Package _struct
/*
@Coding : utf-8
@Time : 2023/2/9 19:27
@Author : yizhigopher
@Software : GoLand
*/
package main

import (
	"fmt"
	"github.com/google/wire"
)

type (
	Foo int
	Bar int
)

func NewFoo() Foo {
	return 1
}

func NewBar() Bar {
	return 2
}

type FooBar struct {
	MyFoo Foo
	MyBar Bar
}

// MySet 实现了类似Setter方法
var MySet = wire.NewSet(NewFoo, NewBar, wire.Struct(new(FooBar), "*"))

func main() {
	fmt.Println(InitStruct())
}
