package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 0{
		for index,arg := range os.Args{ // os.Args是一个存储命令行参数的字符串切片，它的第一个元素是执行文件的名称。
			fmt.Printf("args[%d]=%v\n",index,arg)
		}
	}
}
