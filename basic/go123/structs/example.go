package main

import "structs"

type HasHostLayer struct {
	// 申明指定结构体采用和主机相同的内存布局
	_ structs.HostLayout

}